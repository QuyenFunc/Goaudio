package db

import (
	util "GoAudio/db/util"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

//var testStore Store

// tao ra bien testDB de luu ket qua sql
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	//connPool, err := pgxpool.New(context.Background(), config.DBSource)
	//if err != nil {
	//	log.Fatal("cannot connect to db:", err)
	//}
	//
	//testStore = NewStore(connPool)
	//os.Exit(m.Run())

	//var err error
	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("connot connect to db")
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
