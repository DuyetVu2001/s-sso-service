package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/api/endpoint", func(c *gin.Context) {
		// Define a struct to represent the JSON body
		type Data struct {
			Key1 string `json:"key1"`
			Key2 string `json:"key2"`
		}

		// Initialize a variable to store the JSON data
		var requestData Data

		// Bind the JSON data from the request body to the struct
		if err := c.BindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("=========")
		fmt.Println(requestData.Key1)
		fmt.Println(requestData.Key2)
		fmt.Println("=========")

		// You can now access requestData.Key1 and requestData.Key2

		// Optionally, you can do something with the data here

		// Respond with a JSON message indicating success
		c.JSON(http.StatusOK, gin.H{"message": "POST request successful"})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
