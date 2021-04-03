package service

import (
	"errors"
	"github.com/ShaDream/kitsune-bot/models"
	"github.com/ShaDream/kitsune-bot/repository"
)

type AccessService struct {
	repo repository.Repository
}

func (a AccessService) SetRoleAccess(roleId string, access models.RoleAccess) error {
	if a.HasRoleAccess(roleId) {
		return a.repo.AccessRepository.UpdateRoleAccess(roleId, access)
	} else {
		_, err := a.repo.AccessRepository.CreateRoleAccess(roleId, access)
		return err
	}
}

func (a AccessService) RemoveRoleAccess(roleId string) error {
	if !a.HasRoleAccess(roleId) {
		return errors.New("У роли нет никаких прав.")
	}
	return a.repo.AccessRepository.DeleteRoleAccess(roleId)
}

func (a AccessService) HasRoleAccess(roleId string) bool {
	return a.repo.AccessRepository.HasRoleAccess(roleId)
}

func (a AccessService) GetAllRoleAccesses() ([]*models.Role, error) {
	roles, err := a.repo.AccessRepository.GetAllRolesAccesses()
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func NewAccessService(repo repository.Repository) *AccessService {
	return &AccessService{repo: repo}
}
