package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IPAuthMiddleware() gin.HandlerFunc {
	ipList := []string{
		"127.0.0.2",
	}
	return func(c *gin.Context) {
		flag := false
		clientIp := c.ClientIP()
		for _, host := range ipList {
			if clientIp == host {
				flag = true
				break
			}
		}
		if !flag {
			c.String(http.StatusUnauthorized, "%s, not in iplist", clientIp)
			c.Abort()
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(IPAuthMiddleware())
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "HELLO TEST")
	})

	// 这个方法为什么是阻塞的？ ： server.ListenAndServe()
	r.Run()
}
