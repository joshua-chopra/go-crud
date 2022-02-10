package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/joshua-chopra/go-crud/controllers"
)

// BookRouter /*
// Router grouping together path operation functions
// for /book/ prefix.
func BookRouter(rg *gin.RouterGroup) {
	rg.GET("/", controllers.GetBooks)
	rg.GET("/:id", controllers.GetBook)
	rg.POST("/", controllers.CreateBook)
	rg.PUT("/:id", controllers.UpdateBook)
	rg.DELETE("/:id", controllers.DeleteBook)
}
