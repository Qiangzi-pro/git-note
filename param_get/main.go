package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		firstName := c.Query("first_name")
		lastName := c.DefaultQuery("last_name", "default_last_name")
		c.String(200, "%s %s\n", firstName, lastName)
	})

	r.Run()
}
