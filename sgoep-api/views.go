package main

import (
	"net/http"
	"strings"

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

func getAchievements(c *gin.Context) {

	var achievements []Achievement

	rows, err := db.Query("SELECT * FROM achievements")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var ach Achievement
		var techs string
		err = rows.Scan(
			&ach.Id, &ach.Name, &ach.Subtitle, &ach.Desc, &ach.Icon,
			&ach.Link, &ach.Link_icon, &ach.Github, &techs,
		)
		ach.Tech = strings.Split(techs, " ")

		achievements = append(achievements, ach)
	}

	c.IndentedJSON(http.StatusOK, achievements)
}

func getInterests(c *gin.Context) {

	interests := []Interest{}

	err := db.Select(&interests, "SELECT * FROM interests")
	checkErr(err)

	c.IndentedJSON(http.StatusOK, interests)
}
