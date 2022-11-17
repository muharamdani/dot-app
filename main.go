package main

import (
	"os"
	"reflect"
	"strconv"
	"strings"

	"dot-app/cmd"
	"dot-app/db"
	"dot-app/pkg"
	"dot-app/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Set root path
	cwd, _ := os.Getwd()
	utils.SetRootPath(cwd + "/")

	// Initialize cmd tool
	cmd.Execute()

	// Connect to database
	db.Connect()

	// Load environment variables
	port := utils.Env("PORT", "8080")
	mode := utils.Env("MODE", "debug")

	// Start gin instance
	gin.SetMode(mode)
	r := gin.Default()
	api := r.Group("/api")

	r.SetTrustedProxies([]string{
		"127.0.0.1",
	})

	// Instantiate validator json tag
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	// Register routes
	pkg.Router(api)

	// Run server
	run := r.Run(":" + port)

	// Hande if port is already in use
	basePort := 8081
	for run != nil {
		run = r.Run(":" + strconv.Itoa(basePort))
		basePort++
	}
}
