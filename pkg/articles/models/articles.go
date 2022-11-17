package models

import (
	"dot-app/pkg/comments/models"
	"dot-app/utils"

	"github.com/gofrs/uuid"
)

type Article struct {
	Id       uuid.UUID        `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title    string           `json:"title" gorm:"type:varchar(30);not null"`
	Content  string           `json:"content" gorm:"type:text;not null"`
	Author   string           `json:"author" gorm:"type:varchar(30);not null;unique"`
	Comments []models.Comment `json:"comments" gorm:"foreignKey:ArticleId;references:Id"`
	utils.TimeStamps
}
