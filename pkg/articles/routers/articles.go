package routers

import (
	"dot-app/pkg/articles/controllers"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.POST("/", controllers.Create)
	r.GET("/", controllers.Index)
	r.GET("/:id", controllers.Show)
	r.PUT("/:id", controllers.Update)
	r.PATCH("/:id", controllers.Patch)
	r.DELETE("/:id", controllers.Delete)
}
