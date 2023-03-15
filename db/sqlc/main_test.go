package db

import (
	"database/sql"
	"github.com/avosaga/go-simplebank/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("couldn't read config file:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("couldn't connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
