package db

import (
	"database/sql"
	"log"

	"github.com/BlackestDawn/loganalysis-demo/go-api/internal/config"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/jackc/pgx/v5"
)

func Init(conf *config.Config) (conn *pgx.Conn, err error) {
	RunMigrations(conf.DbUrl)

	conn, err = pgx.Connect(conf.DbCtx, conf.DbUrl)
	if err != nil {
		return
	}

	return
}

func RunMigrations(databaseURL string) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://sql/schema",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	log.Println("Migrations applied successfully")
}
