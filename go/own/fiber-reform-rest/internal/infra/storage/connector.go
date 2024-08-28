package storage

import (
	"database/sql"
	"fmt"
	"log"

	"fiber-reform-rest/internal/infra/env"

	"github.com/go-sql-driver/mysql"
	"gopkg.in/reform.v1"
	dialectsMysql "gopkg.in/reform.v1/dialects/mysql"
)

func Setup() *reform.DB {
	username := env.GetEnv("MYSQL_USER", "admin")
	password := env.GetEnv("MYSQL_PASSWORD", "123")
	dbHost := env.GetEnv("DB_HOST", "localhost")
	dbPort := env.GetEnv("DB_PORT", "3307")
	dbName := env.GetEnv("MYSQL_DATABASE", "frr")
	dbAddr := fmt.Sprintf("%s:%s", dbHost, dbPort)
	cfg := mysql.Config{
		User:   username,
		Passwd: password,
		Net:    "tcp",
		Addr:   dbAddr,
		DBName: dbName,
	}

	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// defer db.Close()
	reformDB := reform.NewDB(db, dialectsMysql.Dialect, nil)
	return reformDB
}
