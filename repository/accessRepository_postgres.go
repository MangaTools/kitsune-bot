package repository

import (
	"errors"
	"fmt"
	"github.com/ShaDream/kitsune-bot/models"
	"github.com/sirupsen/logrus"
)

type AccessRepositoryPostgres struct {
	db DbWorker
}

func (a AccessRepositoryPostgres) GetAllRolesAccesses() ([]*models.Role, error) {
	roles := make([]*models.Role, 0)
	query := fmt.Sprintf("SELECT id, role_id, access_level FROM %s", rolesAccessTable)
	rows, err := a.db.Query(query)
	if err != nil {
		logrus.Error(err)
		return nil, errors.New("Произошла ошибка при получении ролей.")
	}

	for rows.Next() {
		role := new(models.Role)
		err := rows.Scan(&role.Id, &role.RoleId, &role.AccessLevel)
		if err != nil {
			logrus.Error(err)
			return nil, errors.New("Произошла ошибка при получении ролей.")
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (a AccessRepositoryPostgres) CreateRoleAccess(roleId string, access models.RoleAccess) (*models.Role, error) {
	role := new(models.Role)
	query := fmt.Sprintf("INSERT INTO %s (role_id, access_level) VALUES ($1, $2) RETURNING id, role_id, access_level", rolesAccessTable)
	err := a.db.QueryRow(query, roleId, access).Scan(&role.Id, &role.RoleId, &role.AccessLevel)
	if err != nil {
		logrus.Error(err)
		return nil, errors.New("Произошла ошибка при создании доступа роли.")
	}

	return role, nil
}

func (a AccessRepositoryPostgres) UpdateRoleAccess(roleId string, newAccess models.RoleAccess) error {
	query := fmt.Sprintf("UPDATE %s SET access_level=$1 WHERE role_id = $2", rolesAccessTable)
	if _, err := a.db.Exec(query, newAccess, roleId); err != nil {
		logrus.Error(err)
		return errors.New("Ошибка при обновлении доступа роли.")
	}
	return nil
}

func (a AccessRepositoryPostgres) DeleteRoleAccess(roleId string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE role_id = $1", rolesAccessTable)
	if _, err := a.db.Exec(query, roleId); err != nil {
		logrus.Error(err)
		return errors.New("Ошибка при удалении доступа роли.")
	}
	return nil
}

func (a AccessRepositoryPostgres) HasRoleAccess(roleId string) bool {
	var exists bool
	query := fmt.Sprintf("SELECT exists(SELECT * FROM %s WHERE role_id = $1)", rolesAccessTable)
	err := a.db.QueryRow(query, roleId).Scan(&exists)
	if err != nil {
		logrus.Error(err)
		return false
	}
	return exists
}

func NewAccessRepositoryPostgres(db DbWorker) *AccessRepositoryPostgres {
	return &AccessRepositoryPostgres{db: db}
}
