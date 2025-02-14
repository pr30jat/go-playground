package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/samber/lo"
	"github.com/volatiletech/sqlboiler/v4/queries"
)

func main() {
	lo.Must0(godotenv.Load())

	dbDSN := os.Getenv("DB_DSN")

	db := lo.Must(sql.Open("pgx", dbDSN))
	defer db.Close()

	obj := struct {
		X int `boil:"x" db:"x"`
	}{}
	lo.Must0(queries.Raw(`select 1 x`).Bind(context.Background(), db, &obj))
	fmt.Println(obj)
}
