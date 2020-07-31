package main

import (
	"log"
	"net/http"
	_ "os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/", "/home/mgtniip/admin")
	router.Use(cors.Default())
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "/home/mgtniip/admin/index.html", gin.H{"title": "Home Page"})
	})
	log.Fatal(router.Run(":5885"))
}
