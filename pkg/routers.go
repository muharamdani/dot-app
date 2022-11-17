package pkg

import (
	articles "dot-app/pkg/articles/routers"
	comments "dot-app/pkg/comments/routers"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	articles.Router(r)
	comments.Router(r)
}
