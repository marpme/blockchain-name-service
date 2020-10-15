package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.StaticFile("/favicon.ico", "./assets/favicon.ico")
	app.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/v1/")
	})

	v1 := app.Group("/v1")
	{
		v1.GET("/healthz", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		v1.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "Welcome home, this is the blockchain name service")
		})
	}

	app.Run()
}
