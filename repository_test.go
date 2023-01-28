package belajar_go_database

import (
	"belajar-go-database/entity"
	"belajar-go-database/repository"
	"context"
	"fmt"
	"testing"
)

func TestRepository(t *testing.T) {
	authRepository := repository.NewAuthRepository(GetConnection())
	ctx := context.Background()

	auth := entity.Auth{
		Username: "Filial@admin.com",
		Password: "123",
	}

	res, err := authRepository.Insert(ctx, auth)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func TestFindById(t *testing.T) {
	authRepository := repository.NewAuthRepository(GetConnection())

	auth, err := authRepository.FindById(context.Background(), 21)
	if err != nil {
		panic(err)
	}

	fmt.Println(auth)
}

func TestFindAll(t *testing.T) {
	authRepository := repository.NewAuthRepository(GetConnection())

	auth, err := authRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, authenthication := range auth {
		fmt.Println(authenthication)
	}

}
