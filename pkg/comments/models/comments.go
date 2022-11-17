package models

import (
	"dot-app/utils"

	"github.com/gofrs/uuid"
)

type Comment struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	ArticleId uuid.UUID `json:"article_id" gorm:"type:uuid;default:null;not null"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	utils.TimeStamps
}
