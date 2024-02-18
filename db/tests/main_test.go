package db_test

import (
	"database/sql"
	"fmt"
	db "github/adefemi/fingreat_backend/db/sqlc"
	"github/adefemi/fingreat_backend/utils"
	"log"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var testQuery *db.Store

const testDbName = "testdb"
const sslmode = "?sslmode=disable"

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("Could not load env config", err)
	}

	db_source := utils.GetDBSource(config, config.DB_name)
	conn, err := sql.Open(config.DBdriver, db_source)
	if err != nil {
		log.Fatalf("Could not connect to %s server %v", config.DBdriver, err)
	}

	// create database for testing purposes
	_, err = conn.Exec(fmt.Sprintf("CREATE DATABASE %s;", testDbName))
	if err != nil {
		log.Fatalf("Encountered an error creating database %v", err)
	}

	db_source = utils.GetDBSource(config, testDbName)
	tconn, err := sql.Open(config.DBdriver, db_source)
	if err != nil {
		teardown(conn)
		log.Fatalf("Encountered an error creating database %v", err)
	}

	driver, err := postgres.WithInstance(tconn, &postgres.Config{})
	if err != nil {
		teardown(conn)
		log.Fatalf("Cannot create migrate driver %v", err)
	}

	mig, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", "../migrations"),
		config.DBdriver, driver)

	if err != nil {
		teardown(conn)
		log.Fatalf("migration setup failed %v", err)
	}

	if err = mig.Up(); err != nil && err != migrate.ErrNoChange {
		teardown(conn)
		log.Fatalf("migration up failed %v", err)
	}

	testQuery = db.NewStore(tconn)

	code := m.Run()

	tconn.Close()

	teardown(conn)

	os.Exit(code)
}

func teardown(conn *sql.DB) {
	_, err := conn.Exec(fmt.Sprintf("DROP DATABASE %s WITH (FORCE);", testDbName))
	if err != nil {
		log.Fatalf("failed to drop test database %v", err)
	}
	conn.Close()
}
