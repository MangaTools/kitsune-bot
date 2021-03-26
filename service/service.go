package service

import "github.com/ShaDream/kitsune-bot/repository"

type MangaMethods interface {
}

type Service struct {
	MangaMethods
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		MangaMethods: NewMangaService(r),
	}
}
