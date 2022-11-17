package requests

import (
	"dot-app/pkg/comments/models"

	"github.com/gofrs/uuid"
)

type ArticleRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Author  string `json:"author" binding:"required"`
}

type ArticleShow struct {
	Id       uuid.UUID        `json:"id"`
	Title    string           `json:"title"`
	Content  string           `json:"content"`
	Author   string           `json:"author"`
	Comments []models.Comment `json:"comments" gorm:"foreignKey:ArticleId;references:Id"`
}

type ArticlePatch struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Author  string `json:"author,omitempty"`
}

type Paginate struct {
	PerPage  int    `form:"per_page,default=10" binding:"omitempty,min=1"`
	Page     int    `form:"page,default=1" binding:"omitempty,min=1"`
	Paginate bool   `form:"paginate,default=true"`
	OrderBy  string `form:"order_by,default=created_at" binding:"oneof=full_name email status status_info last_active_at created_at"`
	Sort     string `form:"sort,default=desc"`
}

func (ArticleRequest) TableName() string {
	return "articles"
}

func (ArticleShow) TableName() string {
	return "articles"
}

func (ArticlePatch) TableName() string {
	return "articles"
}
