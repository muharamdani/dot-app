package db

import (
	"fmt"
	"log"

	"dot-app/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func credentials() (string, string, string, string, string) {
	dbHost := utils.Env("DB_HOST", "localhost")
	dbPort := utils.Env("DB_PORT", "5432")
	dbUser := utils.Env("DB_USER", "postgres")
	dbPass := utils.Env("DB_PASS", "postgres")
	dbName := utils.Env("DB_NAME", "postgres")

	return dbUser, dbPass, dbHost, dbPort, dbName
}

// Connect creates a connection to postgresql database and
func Connect() {

	user, password, host, port, dbname := credentials()

	// postgres://user:password@host:port/dbname
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	log.Println("Database connected")

	DB = db
}
