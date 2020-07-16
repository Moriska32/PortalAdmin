package main

import (
	routes "PortalAdmin/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	_ "os"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("/home/mgtniip/admin/*")
	router.Use(cors.Default())
	routes.Routes(router)
	log.Fatal(router.Run(":5885"))
}
