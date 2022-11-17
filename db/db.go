package db

import (
	"fmt"
	"log"

	"dot-app/utils"
	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect creates a connection to postgresql database and
func Connect(user string, password string, host string, port string, dbname string) {
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
	hhhh := utils.Env("DB_HOST", "inihost")
	fmt.Println("MANGGIL ENV", hhhh)
	DB = db
}
