package server

import (
	"blogGoServer/cmd/server/handlers"
	"time"
	"github.com/itsjamie/gin-cors"
	"github.com/gin-gonic/gin"
)

func RunServer(){

	r := gin.Default()


	// Set up CORS middleware options
	// Apply the middleware to the router (works with groups too)
	r.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: false,
		ValidateHeaders: false,
	}))


	r.GET("/", handlers.RootGetHandler())
	r.GET("/ping", handlers.PingGetHandler())
	r.GET("/posts", handlers.GetAllPostsHandler())
	r.POST("/posts", handlers.AddPostHandler())
	r.DELETE("/posts", handlers.DeleteAllPostsHandler())
	r.DELETE("/posts/:id", handlers.DeletePostByIdHandler())

	r.Run()
}
