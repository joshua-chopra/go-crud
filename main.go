package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joshua-chopra/go-crud/internal"
	"github.com/joshua-chopra/go-crud/routers"
	"net/http"
	"os"
)

func main() {
	internal.Setup()
	internal.ConnectDatabase()
	port := os.Getenv("SERVER_PORT")
	// instantiate router and add in groups under it.
	router := gin.Default()
	mainRouter := router.Group("/api")
	{
		routers.BookRouter(mainRouter.Group("/book"))
	}
	// simple testing route.
	mainRouter.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	serverPath := fmt.Sprintf("localhost:%s", port)
	fmt.Println(serverPath)
	if err := router.Run(serverPath); err != nil {
		fmt.Printf("Couldn't start server at port %s, exiting program due to err: \n\t%s\n", port, err)
	}
}
