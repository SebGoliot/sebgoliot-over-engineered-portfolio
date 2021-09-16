package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAuthor(c *gin.Context) {

	author := Author{}

	err := db.Get(&author, "SELECT * FROM author LIMIT 1")
	checkErr(err)

	c.IndentedJSON(http.StatusOK, author)
}

func getSkills(c *gin.Context) {

	skills := []Skill{}

	err := db.Select(&skills, "SELECT * FROM skills")
	checkErr(err)

	c.IndentedJSON(http.StatusOK, skills)
}
