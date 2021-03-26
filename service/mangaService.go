package service

import "github.com/ShaDream/kitsune-bot/repository"

type MangaService struct {
	repo repository.MangaRepository
}

func NewMangaService(repo repository.MangaRepository) *MangaService {
	return &MangaService{repo: repo}
}
