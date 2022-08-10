package handlers

import (
	"blogGoServer/cmd/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
)

func AddPostHandler() gin.HandlerFunc{
	return func(c *gin.Context){
		inMemoryList := models.GetMemory()

		// get post body from request
		requestBody := models.Post{}
		c.Bind(&requestBody)

		if requestBody.Body == "" || requestBody.Title == "" {
			c.JSON(http.StatusBadRequest, gin.H {
				"message": "Could not be created, sent an empty value",
			})

			return 
		}
		newId := uuid.New()
		inMemoryList.AddPost(requestBody, newId)
		
		c.JSON(http.StatusCreated, gin.H{
			"message": "The Post has been added to the list",
		})
	}
}
