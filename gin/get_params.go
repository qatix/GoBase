package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"fmt"
	"time"
)

type Person struct {
	Name     string    `form:"name""`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"YYYY-MM-DD" time_utc:"1"`
	//Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func startPage(c *gin.Context) {
	fmt.Println("startPage")
	var person Person
	fmt.Println("person1:")
	fmt.Println(person)
	e := c.ShouldBindQuery(&person)
	if e == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
		fmt.Println(person)
	} else {
		fmt.Println("not null")
		fmt.Println(e)
	}
	fmt.Println("after")
	fmt.Println(person)
	c.JSON(200, person)
	//c.String(http.StatusOK,"Success")
}

func main() {
	route := gin.Default()
	route.GET("/person", startPage)
	route.Static("/assets", "./assets")
	//route.GET("/JSONP?callback=x", func(c *gin.Context) {
	//	data := map[string]interface{}{
	//		"foo": "bar",
	//	}
	//
	//	//callback is x
	//	// Will output  :   x({\"foo\":\"bar\"})
	//	c.JSONP(http.StatusOK, data)
	//})

	route.Run(":8080")
}
