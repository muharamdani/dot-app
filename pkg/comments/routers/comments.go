package routers

import (
	"dot-app/pkg/comments/controllers"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.POST("/", controllers.Create)
	r.PATCH("/:id", controllers.Patch)
	r.DELETE("/:id", controllers.Delete)
}
