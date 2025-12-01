package db

import (
	// "context"
	"log"
	"os"
	"testing"
	"database/sql"

	// "github.com/jackc/pgx/v5/pgxpool"
	// "github.com/dev-ed-bau/simplebank/util"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

// var testStore Store

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testQueries = New(conn)
	// config, err := util.LoadConfig("../..")


	// connPool, err := pgxpool.New(context.Background(), config.DBSource)
	// if err != nil {
	// 	log.Fatal("cannot connect to db:", err)
	// }

	// testStore = NewStore(connPool)
	os.Exit(m.Run())
}