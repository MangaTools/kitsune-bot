package service

import (
	"errors"
	"fmt"
	"github.com/ShaDream/kitsune-bot/models"
	"github.com/ShaDream/kitsune-bot/repository"
	"sort"
)

type WorkService struct {
	repo repository.Repository
}

func (w WorkService) CheckWork(workId int) error {
	if !w.repo.HasWork(workId) {
		return errors.New("Такой работы не существует!")
	}
	work, err := w.repo.GetWork(workId)
	if err != nil {
		return err
	}
	if work.Status != models.InProgress && work.Status != models.OnCompletion {
		return errors.New("Работа уже выполнена или на проверке!")
	}
	if err = w.repo.SetWorkStatus(workId, models.OnCheck); err != nil {
		return err
	}
	return nil
}

func (w WorkService) BookWork(userId string, chapterId int, workType models.WorkType, startPage int, endPage int) (int, error) {
	if !w.repo.ChapterRepository.HasChapter(chapterId) {
		return -1, errors.New("Главы с таким Id не существует.")
	}
	chapter, err := w.repo.ChapterRepository.GetChapter(chapterId)
	if err != nil {
		return -1, err
	}
	if chapter.Pages < endPage {
		return -1, errors.New(fmt.Sprintf("В манге только %d страниц", chapter.Pages))
	}
	alreadyBookedWork, err := w.repo.GetWorksByWorkType(chapterId, workType)
	if err != nil {
		return -1, err
	}
	canBook := w.canBook(alreadyBookedWork, startPage, endPage)
	if !canBook {
		return -1, errors.New("Невозможно зарезервировать. Возможно эти страницы уже зарезервированы.")
	}
	work, err := w.repo.CreateWork(userId, chapterId, startPage, endPage, workType)
	return work, nil
}

func (w WorkService) RemoveBookedWork(workId int) error {
	if !w.repo.HasWork(workId) {
		return errors.New("Такой работы не существует!")
	}
	work, err := w.repo.GetWork(workId)
	if err != nil {
		return err
	}
	if work.Status != models.InProgress {
		return errors.New("Работу уже нельзя удалить!")
	}
	err = w.repo.DeleteWork(workId)
	if err != nil {
		return err
	}
	return nil
}

var workTypeToUserCharacteristic = map[models.WorkType]models.UserCharacteristic{
	models.Clean:     models.UserCharacteristicCleanedPages,
	models.Edit:      models.UserCharacteristicEditedPages,
	models.Type:      models.UserCharacteristicTypedPages,
	models.Translate: models.UserCharacteristicTranslatedPages,
}

func (w WorkService) DoneWork(workId int) (err error) {
	tx, err := w.repo.BeginTransaction()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	workRepo := repository.NewWorkRepositoryPostgres(tx)
	userRepo := repository.NewUserRepositoryPostgres(tx)
	chapterRepo := repository.NewChapterRepositoryPostgres(tx)

	work, err := workRepo.GetWork(workId)
	if err != nil {
		return err
	}
	if work.Status != models.OnCheck {
		return errors.New("Работу нельзя пометить готовой, пока она не на проверке!")
	}

	if err = workRepo.SetWorkStatus(workId, models.Done); err != nil {
		return err
	}

	works, err := workRepo.GetWorksByWorkType(work.ChapterId, work.Type)
	if err != nil {
		return err
	}
	if err = w.mergeIfCan(works); err != nil {
		return err
	}
	chapter, err := chapterRepo.GetChapter(work.ChapterId)
	if err != nil {
		return err
	}
	if workRepo.IsChapterDone(*chapter) {
		err = chapterRepo.SetChapterStatus(chapter.Id, models.DoneChapter)
		if err != nil {
			return err
		}
	}

	pages := work.PageEnd - work.PageStart + 1

	if _, err := userRepo.AddToField(work.UserId, workTypeToUserCharacteristic[work.Type], pages); err != nil {
		return err
	}
	if _, err := userRepo.AddToField(work.UserId, models.UserCharacteristicScore, pages); err != nil {
		return err
	}

	return tx.Commit()
}

func (w WorkService) canBook(books []*models.Owner, startPage int, endPage int) bool {
	for _, book := range books {
		left := max(book.PageStart, startPage)
		right := min(book.PageEnd, endPage)
		if left <= right {
			return false
		}
	}
	return true
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Merges works if they are done and made by one user
func (w WorkService) mergeIfCan(books []*models.Owner) error {
	result := make([][]*models.Owner, 0)
	books = filterWorkByDone(books)
	if len(books) < 2 {
		return nil
	}
	sort.SliceStable(books, func(i, j int) bool {
		return books[i].PageStart < books[j].PageStart
	})
	groups := groupBooksByOwner(books)

	newSlice := true
	for _, groupBooks := range groups {
		if len(groupBooks) < 2 {
			continue
		}
		for i, j := 0, 1; j < len(groupBooks); i, j = i+1, j+1 {
			bookStart := groupBooks[i]
			bookEnd := groupBooks[j]
			if bookStart.PageEnd+1 == bookEnd.PageStart {
				if newSlice {
					result = append(result, []*models.Owner{bookStart, bookEnd})
				} else {
					arr := result[len(result)-1]
					arr = append(arr, bookEnd)
				}
			} else {
				newSlice = true
			}
		}
		newSlice = true
	}
	if len(result) == 0 {
		return nil
	}
	return w.repo.MergeWorks(result)
}

func filterWorkByDone(books []*models.Owner) []*models.Owner {
	result := make([]*models.Owner, 0)
	for _, book := range books {
		if book.Status == models.Done {
			result = append(result, book)
		}
	}
	return result
}

func groupBooksByOwner(books []*models.Owner) map[string][]*models.Owner {
	result := make(map[string][]*models.Owner, 0)
	for _, book := range books {
		if val, ok := result[book.UserId]; ok {
			val = append(val, book)
		}
		result[book.UserId] = []*models.Owner{book}
	}
	return result
}

func NewWorkService(repo repository.Repository) *WorkService {
	return &WorkService{
		repo: repo,
	}
}
