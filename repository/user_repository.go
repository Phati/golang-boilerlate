package repository

import (
	"context"

	"github.com/Phati/golang-boilerplate/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(context context.Context, user *models.User) (*models.User, error)
	// GetUserByID(context context.Context, id int) (*models.User, error)
	// GetAllUsers(context context.Context) ([]*models.User, error)
}

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &UserRepo{
		db: db,
	}
}

func (userRepo *UserRepo) CreateUser(context context.Context, user *models.User) (*models.User, error) {

	return nil, nil
}
