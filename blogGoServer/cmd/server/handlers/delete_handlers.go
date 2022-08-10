package handlers

import (
	"blogGoServer/cmd/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DeleteAllPostsHandler()gin.HandlerFunc{
	return func(c *gin.Context) {
		// clear memory
		inMemory := models.GetMemory()	
		inMemory.DeleteAllPosts()
		c.JSON(http.StatusOK, gin.H{
			"message": "Delete all succesful",
		})
	}
}
func DeletePostByIdHandler()gin.HandlerFunc{
	return func(c *gin.Context){
		inMemory := models.GetMemory()

		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad ID",
			})
			return 
		}
		error := inMemory.DeleteById(id)
		if error != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "No ID found",
			})
			return 
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Deleted Post",
		})
	}
}
