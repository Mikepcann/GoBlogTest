package handlers

import (
	"blogGoServer/cmd/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func RootGetHandler() gin.HandlerFunc {
	return func(c *gin.Context){
		c.JSON(200, gin.H{
				"message": "All Good",
			})
	}
}

func PingGetHandler() gin.HandlerFunc {
	return func(c *gin.Context){
		c.JSON(200, gin.H{
				"message": "pong",
			})
	}
}

func GetAllPostsHandler()gin.HandlerFunc{

	inMemory := models.GetMemory()
	return func(c *gin.Context){
		listOfPosts := inMemory.GetAllPosts()


		fmt.Println("Check the number of posts", len(listOfPosts))
		if len(listOfPosts) == 0 {
			c.JSON(http.StatusOK,gin.H{
				"message": "No List data Found",
			})
			
			return
		}

		c.JSON(200,listOfPosts)
	}
}
