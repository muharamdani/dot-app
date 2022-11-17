package routers

import (
	"dot-app/pkg/comments/controllers"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.POST("/comments", controllers.Create)
	r.PATCH("/comments/:id", controllers.Patch)
	r.DELETE("/comments/:id", controllers.Delete)
}
