package main

import (
	// "encoding/json"
	// "fmt"
	"net/http"
	// "strconv"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

/*
	All Go programs start running from a function called `main` in a package called `main`
*/
func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/home/", getHome)
	router.Run("localhost:8080")
}

func getHome(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Called successfully"})
}