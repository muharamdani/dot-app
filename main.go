package main

import (
	"strconv"

	"dot-app/db"
	"dot-app/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Load environment variables
	port := utils.Env("PORT", "8080")
	mode := utils.Env("MODE", "debug")
	dbHost := utils.Env("DB_HOST", "localhost")
	dbPort := utils.Env("DB_PORT", "5432")
	dbUser := utils.Env("DB_USER", "postgres")
	dbPass := utils.Env("DB_PASS", "postgres")
	dbName := utils.Env("DB_NAME", "postgres")

	// Connect to database
	db.Connect(dbUser, dbPass, dbHost, dbPort, dbName)

	// Start gin instance
	gin.SetMode(mode)
	r := gin.Default()
	r.SetTrustedProxies([]string{
		"127.0.0.1",
	})

	// Run server
	run := r.Run(":" + port)

	// Hande if port is already in use
	basePort := 8081
	for run != nil {
		run = r.Run(":" + strconv.Itoa(basePort))
		basePort++
	}
}
