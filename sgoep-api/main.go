package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

func main() {

	godotenv.Load(".env")

	var err error
	db, err = sqlx.Connect("sqlite3", os.Getenv("DB_FILE"))
	checkErr(err)

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://127.0.0.1:8080",  // Debug VueJS Host
		os.Getenv("CORS_ORIGIN"), // Production VueJS host
	}
	router.Use(cors.New(config))

	router.GET("skills", getSkills)
	router.GET("author", getAuthor)

	router.Run("127.0.0.1:5050")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
