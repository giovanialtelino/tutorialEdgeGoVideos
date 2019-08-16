package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func PostPage(c *gin.Context) {
	body := c.Request.Body

	value, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": string(value),
	})

}

func QueryStrings(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func PathParams(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {
	fmt.Println("hi")

	r := gin.Default()
	r.GET("/", HomePage)
	r.POST("/", PostPage)
	r.GET("/query", QueryStrings)
	r.GET("/path/:name/:age", PathParams)
	r.Run()
}
