package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
)

func main() {

	router := gin.Default()

	router.GET("/helloworld", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Byebye world!!!!"})
	})

	http.Handle("/", router)

	appengine.Main()
}
