package exec

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testcrud/database"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestExecSql(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id,name, email)VALUES('C005','Supargo','supargo@gmail.com')"

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Create Data Customer")
}

func TestQuerySql(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id,name,email FROM customer"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id string
		var name string
		var email string

		err := rows.Scan(&id, &name, &email)
		if err != nil {
			panic(err)
		}

		fmt.Println("ID :", id)
		fmt.Println("NAME :", name)
		fmt.Println("EMAIL :", email)
	}
}

func TestQueryWithDataType(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id,name,email,balance,birthdate,created_at,married from customer"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var birthdate sql.NullTime
		var created_at time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &birthdate, &created_at, &married)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
		if email.Valid {
			fmt.Println("Email :", email.String)
		}
		fmt.Println("Balance :", balance)
		if birthdate.Valid {
			fmt.Println("BirthDate :", birthdate.Time)
		}
		fmt.Println("Created_At :", created_at)
		fmt.Println("Married :", married)
	}
}

func TestSqlInjection(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "arif"
	password := "salah"

	query := "SELECT username from user where username = ? and password = ? limit 1"
	rows, err := db.QueryContext(ctx, query, username, password)
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
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestExecSafe(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "ulla"
	password := "ulla"

	query := "INSERT INTO user(username,password)VALUES(?,?)"

	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("Succes Create Username Password")
}

func TestQuerySqlLast(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT username,password from user"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var username, password string
		err := rows.Scan(&username, &password)
		if err != nil {
			panic(err)
		}
		fmt.Println("USER :", username)
		fmt.Println("PW :", password)
	}
}

func TestExecSqlWithId(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()
	ctx := context.Background()

	name := "Holla"
	email := "Holla@gmail.com"
	comment := "Test Comment Aja"

	query := "INSERT INTO comment(name,email,comment)Values(?,?,?)"
	for i := 1; i < 5; i++ {
		result, err := db.ExecContext(ctx, query, name, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Success Create Comment ID :", id)
	}
}

func TestPrepareStatement(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comment(name,email,comment)VALUES(?,?,?)"

	statement, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		name := "Arif" + strconv.Itoa(i)
		email := "arifbudiman" + strconv.Itoa(i) + "@gmail.com"
		comment := "TEST AJA" + strconv.Itoa(i)
		result, err := statement.ExecContext(ctx, name, email, comment)
		if err != nil {
			panic(err)
		}

		id, _ := result.LastInsertId()
		fmt.Println("Success Create Comment Id :", id)
	}
}
