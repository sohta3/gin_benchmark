package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/welcome", func(c *gin.Context) {
		//c.String(200, "pong")
		c.JSON(http.StatusOK, `[{"name": "Chris McCord"}, {"name": "Matt Sears"}, {"name": "David Stump"}, {"name": "Ricardo Thompson"}]`)
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
