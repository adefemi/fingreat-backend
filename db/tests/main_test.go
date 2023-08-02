package db_test

import (
	"database/sql"
	db "github/adefemi/fingreat_backend/db/sqlc"
	"github/adefemi/fingreat_backend/utils"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQuery *db.Store

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("Could not load env config", err)
	}

	conn, err := sql.Open(config.DBdriver, config.DB_source)
	if err != nil {
		log.Fatal("Could not connect to database", err)
	}

	testQuery = db.NewStore(conn)

	os.Exit(m.Run())
}
