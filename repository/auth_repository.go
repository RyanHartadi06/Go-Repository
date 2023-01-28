package repository

import (
	"belajar-go-database/entity"
	"context"
)

type AuthRepository interface {
	Insert(ctx context.Context, auth entity.Auth) (entity.Auth, error)
	FindById(ctx context.Context, id int32) (entity.Auth, error)
	FindAll(ctx context.Context) ([]entity.Auth, error)
}
