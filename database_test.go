package belajar_go_database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"testing"
	"time"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer(email , balance , rating, created_at , birth_date, is_married) VALUES ('Ryan@gmail.com' , 10000 , 2.5 , '2000-01-07', '2000-01-07' , false)"
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Sukses Kirim Data")

}

func TestOpenConnectionQuery(t *testing.T) {
	db := GetConnection()

	defer db.Close()
	ctx := context.Background()
	script := "SELECT id, email , balance , rating, created_at , birth_date, is_married FROM customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {

		var id, email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool
		err := rows.Scan(&id, &email, &balance, &rating, &createdAt, &birthDate, &married)

		if err != nil {
			panic(err)
		}

		fmt.Println("========")
		fmt.Println("id", id)

		if email.Valid {
			fmt.Println("email", email.String)
		}

		fmt.Println("balance", balance)
		fmt.Println("rating", rating)
		fmt.Println("createdAt", createdAt)

		if birthDate.Valid {
			fmt.Println("Birth Date", birthDate.Time)
		}

		fmt.Println("Is Married", married)
	}
}

func TestCreateAuth(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()
	username := "RyanHartadi"
	password := "123"
	script := "INSERT INTO auth(username , password) VALUES (? , ?)"

	result, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	queryId := "SELECT username FROM auth WHERE id = ? LIMIT 1"

	rows, err := db.QueryContext(ctx, queryId, insertId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next() {

		var username string

		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println("Sukses Daftar Akun", username)
	} else {
		fmt.Println("Gagal")
	}

}

func TestSelectAuth(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()
	username := "Ryan"
	password := "123"
	script := "SELECT username FROM auth WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next() {
		var username string

		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}

		fmt.Println("Success Login")
	} else {
		fmt.Println("Gagal")
	}
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO auth(username , password) VALUES (? , ?)"

	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		username := "Ryan" + strconv.Itoa(i) + "@gmail.com"
		password := "!23"

		res, err := statement.ExecContext(ctx, username, password)

		if err != nil {
			panic(err)
		}

		id, err := res.LastInsertId()

		if err != nil {
			panic(err)
		}
		fmt.Println(id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	script := "INSERT INTO auth(username , password) VALUES (? , ?)"
	for i := 0; i < 10; i++ {
		username := "Ryan" + strconv.Itoa(i) + "@gmail.com"
		password := "!23"

		res, err := tx.ExecContext(ctx, script, username, password)

		if err != nil {
			panic(err)
		}

		id, err := res.LastInsertId()

		if err != nil {
			panic(err)
		}
		fmt.Println(id)
	}

	//tx.Commit()
	err = tx.Commit()

	if err != nil {
		panic(err)
	}
}
