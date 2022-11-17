package migrations

import (
	"fmt"
	"log"

	"dot-app/utils"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	ArticleId uuid.UUID `gorm:"type:uuid;default:null;not null"`
	Article   Article   `gorm:"foreignKey:ArticleId"`
	Content   string    `gorm:"type:text;not null"`
	utils.TimeStamps
}

func MigrateCommentsTable(db *gorm.DB) {
	m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20221117131812",
			Migrate: func(tx *gorm.DB) error {

				return tx.AutoMigrate(&Comment{})
			},
		},
	})
	if err := m.Migrate(); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Migrated create_articles_comments\n")
}
