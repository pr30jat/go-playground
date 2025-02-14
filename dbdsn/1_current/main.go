package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/samber/lo"
	"github.com/volatiletech/sqlboiler/v4/queries"
)

func main() {
	lo.Must0(godotenv.Load())

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	v := url.Values{}
	v.Set("sslmode", "disable")
	v.Set("dbname", dbName)
	dbURL := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(dbUsername, dbPassword),
		Host:     dbHost + ":" + dbPort,
		RawQuery: v.Encode(),
	}
	fmt.Println(dbURL.String())
	q := dbURL.Query()
	dbURL.RawQuery = q.Encode()

	db := lo.Must(sql.Open("pgx", dbURL.String()))
	defer db.Close()

	obj := struct {
		X int `boil:"x" db:"x"`
	}{}
	lo.Must0(queries.Raw(`select 1 x`).Bind(context.Background(), db, &obj))
	fmt.Println(obj)
}
