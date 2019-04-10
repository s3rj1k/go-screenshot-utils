package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(logger())

	createScreenshotFromWeb(r)

	r.NoRoute(func(c *gin.Context) {
		c.File("form.html")
	})

	err := r.Run(fmt.Sprintf(":%d", cmdPort))
	if err != nil {
		log.Fatal(err)
	}
}
