package repository

import (
	"belajar-go-database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type authRepositoryImpl struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepositoryImpl{
		DB: db,
	}
}

func (repository *authRepositoryImpl) Insert(ctx context.Context, auth entity.Auth) (entity.Auth, error) {
	script := "INSERT INTO auth(username , password) VALUES (? ,?)"
	res, err := repository.DB.ExecContext(ctx, script, auth.Username, auth.Password)

	if err != nil {
		return auth, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return auth, err
	}

	auth.Id = id
	return auth, nil
}

func (repository *authRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Auth, error) {

	script := "SELECT id , username , password FROM auth WHERE id = ? LIMIT 1"

	rows, err := repository.DB.QueryContext(ctx, script, id)

	auth := entity.Auth{}

	if err != nil {
		return auth, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&auth.Id, &auth.Username, &auth.Password)
		return auth, nil
	} else {
		return auth, errors.New("id " + strconv.Itoa(int(id)))
	}
}

func (repository *authRepositoryImpl) FindAll(ctx context.Context) ([]entity.Auth, error) {
	script := "SELECT id , username , password FROM auth"

	rows, err := repository.DB.QueryContext(ctx, script)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var authentication []entity.Auth
	for rows.Next() {
		auth := entity.Auth{}
		rows.Scan(&auth.Id, &auth.Username, &auth.Password)
		authentication = append(authentication, auth)
	}
	return authentication, nil
}
