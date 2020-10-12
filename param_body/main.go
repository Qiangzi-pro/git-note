package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		firstName := c.PostForm("first_name")
		lastName := c.PostForm("last_name")
		c.String(200, "%s, %s, %s", firstName, lastName, string(bodyBytes))

	})
	r.Run()
}
