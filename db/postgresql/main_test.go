package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"minibank/config"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	testQueries *Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	config.Init()

	db, err := sql.Open("postgres", viper.GetString("DATABASE_URL"))

	if err != nil {
		log.Fatal(fmt.Errorf("sql.Open %w", err))
	}

	if err := db.Ping(); err != nil {
		log.Fatal(fmt.Errorf("db.Ping %w", err))
	}

	testQueries = New(db)
	os.Exit(m.Run())
}
