package migrations

import (
	"fmt"
	"log"

	"dot-app/utils"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Article struct {
	Id      uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title   string    `gorm:"type:varchar(30);not null"`
	Content string    `gorm:"type:text;not null"`
	Author  string    `gorm:"type:varchar(30);not null;unique"`
	utils.TimeStamps
}

func MigrateArticlesTable(db *gorm.DB) {
	m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20221117131727",
			Migrate: func(tx *gorm.DB) error {

				return tx.AutoMigrate(&Article{})
			},
		},
	})
	if err := m.Migrate(); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Migrated create_articles_table\n")
}
