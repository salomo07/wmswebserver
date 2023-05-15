package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(static.Serve("/", static.LocalFile("./static", false)))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "404", "message": "Endpoint isn't found (Gateway)"})
	})
	r.Run(":7771")
}
