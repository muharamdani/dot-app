package routers

import (
	"dot-app/pkg/articles/controllers"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.POST("/articles", controllers.Create)
	r.GET("/articles", controllers.Index)
	r.GET("/articles/:id", controllers.Show)
	r.PUT("/articles/:id", controllers.Update)
	r.PATCH("/articles/:id", controllers.Patch)
	r.DELETE("/articles/:id", controllers.Delete)
}
