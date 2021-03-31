package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/ShaDream/kitsune-bot/models"
	"github.com/sirupsen/logrus"
	"strings"
)

type WorkRepositoryPostgres struct {
	db DbWorker
}

var (
	setWorkStatusQuery      = fmt.Sprintf("UPDATE %s SET status=$1 WHERE id=$2", workTable)
	getWorkQuery            = fmt.Sprintf("SELECT id, user_id, chapter_id, page_start, page_end, status, work_type FROM %s WHERE id=$1", workTable)
	getWorksByWorkTypeQuery = fmt.Sprintf("SELECT id, user_id, chapter_id, page_start, page_end, status, work_type "+
		"FROM %s WHERE chapter_id=$1 AND work_type=$2", workTable)
	createWorkQuery = fmt.Sprintf("INSERT INTO %s (user_id, chapter_id, page_start, page_end, status, work_type) "+
		"VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", workTable)
	deleteWorkQuery = fmt.Sprintf("DELETE FROM %s WHERE id=$1", workTable)
	hasWorkQuery    = fmt.Sprintf("SELECT exists(SELECT * FROM %s WHERE id=$1)", workTable)
)

func (w WorkRepositoryPostgres) SetWorkStatus(workId int, status models.OwnerPageStatus) error {
	_, err := w.db.Exec(setWorkStatusQuery, status, workId)
	if err != nil {
		logrus.Error(err)
		return errors.New("Произошла ошибка при получении работы.")
	}
	return nil
}

func (w WorkRepositoryPostgres) GetWork(workId int) (*models.Owner, error) {
	work := new(models.Owner)
	err := w.db.QueryRow(getWorkQuery, workId).
		Scan(&work.Id, &work.UserId, &work.ChapterId, &work.PageStart, &work.PageEnd, &work.Status, &work.Type)
	if err != nil {
		logrus.Error(err)
		return nil, errors.New("Произошла ошибка при получении работы.")
	}
	return work, nil
}

func (w WorkRepositoryPostgres) GetWorksByWorkType(chapterId int, workType models.WorkType) ([]*models.Owner, error) {
	works := make([]*models.Owner, 0)
	rows, err := w.db.Query(getWorksByWorkTypeQuery, chapterId, workType)
	if err != nil {
		logrus.Error(err)
		return nil, errors.New("Произошла ошибка при получении работ.")
	}
	for rows.Next() {
		work := new(models.Owner)
		err := rows.Scan(&work.Id, &work.UserId, &work.ChapterId, &work.PageStart, &work.PageEnd, &work.Status, &work.Type)
		if err != nil {
			logrus.Error(err)
			return nil, errors.New("Произошла ошибка при получении работ.")
		}
		works = append(works, work)
	}
	return works, nil
}

func (w WorkRepositoryPostgres) CreateWork(userId string, chapterId int, pageStart int, pageEnd int, workType models.WorkType) (int, error) {
	var workId int
	err := w.db.QueryRow(createWorkQuery, userId, chapterId, pageStart, pageEnd, models.InProgress, workType).Scan(&workId)
	if err != nil {
		logrus.Error(err)
		return -1, errors.New("Произошла ошибка при создании работы.")
	}
	return workId, nil
}

func (w WorkRepositoryPostgres) DeleteWork(workId int) error {
	_, err := w.db.Exec(deleteWorkQuery, workId)
	if err != nil {
		logrus.Error(err)
		return errors.New("Произошла ошибка при получении работы.")
	}
	return nil
}

func (w WorkRepositoryPostgres) HasWork(workId int) bool {
	var hasWork bool
	err := w.db.QueryRow(hasWorkQuery, workId).Scan(&hasWork)
	if err != nil {
		logrus.Error(err)
		return false
	}
	return hasWork
}

func (w WorkRepositoryPostgres) MergeWorks(mergingWorks [][]*models.Owner) error {
	returnedError := errors.New("Произошла ошибка при объединении работ.")
	for _, works := range mergingWorks {
		ctx := context.Background()
		tx, err := w.db.Begin()
		if err != nil {
			logrus.Error(err)
			return returnedError
		}

		deletedArgs := createDeleteArgs(len(works))
		args := make([]interface{}, 0, len(works))
		for _, work := range works {
			args = append(args, work.Id)
		}

		query := fmt.Sprintf("DELETE FROM %s WHERE id IN (%s)", workTable, deletedArgs)
		_, err = tx.ExecContext(ctx, query, args...)
		if err != nil {
			tx.Rollback()
			logrus.Error(err)
			return returnedError
		}

		startPage := works[0].PageStart
		endPage := works[len(works)-1].PageEnd
		work := works[0]

		newWork := new(models.Owner)
		err = tx.QueryRowContext(ctx, createWorkQuery, work.UserId, work.ChapterId, startPage, endPage, models.Done, work.Type).
			Scan(&newWork.Id, &newWork.UserId, &newWork.ChapterId, &newWork.PageStart, &newWork.PageEnd, &newWork.Status, &newWork.Type)
		if err != nil {
			tx.Rollback()
			logrus.Error(err)
			return returnedError
		}
		err = tx.Commit()
		if err != nil {
			logrus.Error(err)
			return returnedError
		}
	}
	return nil
}

func createDeleteArgs(length int) string {
	result := make([]string, 0, length)
	for i := 1; i <= length; i++ {
		result = append(result, fmt.Sprintf("$%d", i))
	}
	return strings.Join(result, ", ")
}

func NewWorkRepositoryPostgres(db DbWorker) *WorkRepositoryPostgres {
	return &WorkRepositoryPostgres{db: db}
}
