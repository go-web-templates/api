package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/go-web-templates/api/internal/infra/data"
	"github.com/go-web-templates/api/pkg/conf"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func printUsage(scriptName string) {
	fmt.Printf(`Usage: %[1]s [setup|up|down]

%[1]s will use the "migrations" folder in
the current directory to perform a full migration
operation using the golang-migrate/migrate lib.

setup   Create the application database if not exists
up      Perform a full migration up
down    Perform a full migration down
`, scriptName)
}

func getMigrate() *migrate.Migrate {
	appConf, err := conf.NewAppConf()
	if err != nil {
		log.Fatal(err)
	}

	db, err := data.NewDatabase(appConf)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db.Ctx, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	return m
}

func checkIfDatabaseExists(db *data.Database, dbName string) (bool, error) {
	var result int
	err := db.Ctx.QueryRow(
		`SELECT 1 FROM pg_database WHERE datname=$1`,
		dbName,
	).Scan(&result)

	if err == sql.ErrNoRows {
		return false, nil
	}

	return result == 1, err
}

func createDatabase() {
	appConf, err := conf.NewAppConf()
	if err != nil {
		log.Fatal(err)
	}

	db, err := data.NewDefaultDatabase(appConf)
	if err != nil {
		log.Fatal(err)
	}

	dbName := appConf.Data.Database.Name
	exists, err := checkIfDatabaseExists(db, dbName)
	if err != nil {
		log.Fatal(err)
	} else if exists {
		fmt.Println("The database already exists, Nothing to do.")
		return
	}

	_, err = db.Ctx.Exec(`CREATE DATABASE ` + dbName)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	args := os.Args
	scriptName := filepath.Base(args[0])

	if len(args) < 2 {
		printUsage(scriptName)
		os.Exit(1)
	}

	switch args[1] {
	case "setup":
		createDatabase()
	case "up":
		m := getMigrate()
		m.Up()
	case "down":
		m := getMigrate()
		m.Down()
	default:
		printUsage(scriptName)
		os.Exit(1)
	}
}
