package migration

import (
	"github.com/Shumodan/go-bookshelf/model"
	"github.com/Shumodan/go-bookshelf/mycontext"
)

// CreateDatabase creates the tables used in this application.
func CreateDatabase(context mycontext.Context) {
	if context.GetConfig().Database.Migration {
		db := context.GetRepository()

		db.DropTableIfExists(&model.Book{})
		db.DropTableIfExists(&model.Category{})
		db.DropTableIfExists(&model.Format{})
		db.DropTableIfExists(&model.Account{})
		db.DropTableIfExists(&model.Authority{})

		db.AutoMigrate(&model.Book{})
		db.AutoMigrate(&model.Category{})
		db.AutoMigrate(&model.Format{})
		db.AutoMigrate(&model.Account{})
		db.AutoMigrate(&model.Authority{})
	}
}
