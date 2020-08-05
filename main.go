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
	router.Use(gin.Recovery())
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "./app/index.html", nil)
	})
	log.Fatal(router.Run(":5885"))
}
