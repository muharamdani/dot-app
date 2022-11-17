package pkg

import (
	articles "dot-app/pkg/articles/routers"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	articlesGroup := r.Group("/articles")
	articles.Router(articlesGroup)
}
