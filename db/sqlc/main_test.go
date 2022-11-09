package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/wizlif/simplebank/util"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config,err:=util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	
	testDB, err = sql.Open(config.DbDriver, config.DbSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
