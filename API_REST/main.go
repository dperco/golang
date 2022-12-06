package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type music struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

var albums = []music{
	{Id: "1", Title: "Family", Artist: "Camila Cabello", Year: 2022},
	{Id: "2", Title: "21", Artist: "Adele", Year: 2011},
	{Id: "3", Title: "The Enimen Show", Artist: "Enimen", Year: 2022},
}

func getmusic(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getmusic)
	fmt.Println("Welcome to GO ")

	router.Run("localhost:8080")
}
