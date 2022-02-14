package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joshua-chopra/go-crud/api/routers"
	"log"
	"net/http"
	"os"
)

func StartRouter() {
	router, serverURL := SetupRouter()
	if err := router.Run(serverURL); err != nil {
		log.Fatalf("Couldn't start server at URL %s, exiting program due to err: \n\t%s\n", serverURL, err)
	}
}

func SetupRouter() (*gin.Engine, string) {
	// instantiate initialRouter and add in groups under it.
	initialRouter := gin.Default()
	// initially we have our API router, which has
	// children routers under it
	apiRouter := initialRouter.Group("/api")
	{
		routers.BookRouter(apiRouter.Group("/book"))
	}
	// simple testing route for base initialRouter. localhost:3000/ping.
	initialRouter.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	serverURL := getServerURL()
	return initialRouter, serverURL
}

func getServerURL() string {
	host, port := os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT")
	return fmt.Sprintf("%s:%s", host, port)
}
