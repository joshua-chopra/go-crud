package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/joshua-chopra/go-crud/api/controllers"
)

// BookRouter /*
// Router grouping together path operation functions
// for /book/ prefix.
func BookRouter(rg *gin.RouterGroup) {
	rg.GET("/", controllers.GetAllBooks)
	rg.GET("/:id", controllers.GetOneBook)
	rg.POST("/", controllers.CreateBook)
	rg.PUT("/:id", controllers.UpdateBook)
	rg.DELETE("/:id", controllers.DeleteBook)
}
