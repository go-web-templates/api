package data

import (
	"database/sql"
	"fmt"

	"github.com/go-web-templates/api/pkg/conf"
	_ "github.com/lib/pq"
)

type Database struct {
	Ctx *sql.DB
}

func NewDatabase(appConf *conf.AppConf) (*Database, error) {
	ctx, err := newPostgresConnection(appConf, nil)
	database := Database{Ctx: ctx}

	return &database, err
}

func NewDefaultDatabase(appConf *conf.AppConf) (*Database, error) {
	defaultDb := "postgres"

	ctx, err := newPostgresConnection(appConf, &defaultDb)
	database := Database{Ctx: ctx}

	return &database, err
}

func newPostgresConnection(appConf *conf.AppConf, database *string) (*sql.DB, error) {
	dbConf := appConf.Data.Database
	dbName := dbConf.Name
	if database != nil {
		dbName = *database
	}

	connectionStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConf.Host,
		dbConf.Port,
		dbConf.User,
		dbConf.Pass,
		dbName,
	)

	return sql.Open("postgres", connectionStr)
}
