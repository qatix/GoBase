package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"io"
)

func main()  {
	router := gin.Default()

	f,_ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"hello simple");
	})

	router.GET("/1", func(c *gin.Context) {
		c.String(http.StatusOK,"test 1");
	})

	router.POST("/post", func(c *gin.Context) {
		c.String(http.StatusUnauthorized,"not authorized");
	})

	router.PUT("/put", func(c *gin.Context) {
		c.String(http.StatusOK,"put ok")
	})
	router.Run(":8081")
}
