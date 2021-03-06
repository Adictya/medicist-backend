package db

import(
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQuerries *Queries

const(
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/root?sslmode=disable"
)

func TestMain(m *testing.M){
	conn, err := sql.Open(dbDriver,dbSource)
	if err != nil {
		log.Fatal("Connection to database failed:",err)
	}

	testQuerries = New(conn)

	os.Exit(m.Run())
}
