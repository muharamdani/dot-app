package migrations

import (
	"fmt"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func MigrateAddParentIdInComments(db *gorm.DB) {
	m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20221117204502",
			Migrate: func(tx *gorm.DB) error {
				tx.Exec(`ALTER TABLE comments 
    			ADD COLUMN parent_id uuid 
    			    REFERENCES comments(id) 
    			        ON DELETE CASCADE ON UPDATE CASCADE`)

				tx.Exec("ALTER TABLE comments ALTER COLUMN article_id DROP NOT NULL")
				return nil
			},
		},
	})
	if err := m.Migrate(); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Migrated alter_table_comments_add_parent_id\n")
}
