package repositories

import (
	"github.com/SofyanHadiA/goscrape/models"

	"github.com/SofyanHadiA/linqcore/database"
)

type userRepo struct {
	db database.IDB
}

func NewUserRepo(db database.IDB) userRepo {
	return userRepo{
		db: db,
	}
}

func (repo userRepo) Login(userName string, password string) (*models.User, error) {
	selectQuery := "SELECT users.* FROM users WHERE users.name = ? AND password = ? "

	user := &models.User{}
	rows, err := repo.db.ResolveSingle(selectQuery, userName, password)
	if err != nil {
		return nil, err
	}
	rows.StructScan(user)

	userRole, err := repo.GetRole(user.RoleID)
	if err != nil {
		return nil, err
	}

	user.UserRole = *userRole

	return user, err
}

func (repo userRepo) GetRole(roleID int) (*models.UserRole, error) {
	selectQuery := "SELECT * FROM user_roles WHERE id = ?"

	role := &models.UserRole{}
	rows, err := repo.db.ResolveSingle(selectQuery, roleID)
	if err != nil {
		return nil, err
	}
	rows.StructScan(role)

	return role, err
}
