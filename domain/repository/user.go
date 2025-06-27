package repository

import (
	"context"

	"github.com/railgun-0402/ddd-go-jwt/domain/model"
)

type Repository interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}