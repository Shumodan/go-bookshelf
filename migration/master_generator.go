package migration

import (
	"github.com/Shumodan/go-bookshelf/model"
	"github.com/Shumodan/go-bookshelf/mycontext"
)

// InitMasterData creates the master data used in this application.
func InitMasterData(context mycontext.Context) {
	if context.GetConfig().Extension.MasterGenerator {
		rep := context.GetRepository()

		r := model.NewAuthority("Admin")
		_, _ = r.Create(rep)
		a := model.NewAccountWithPlainPassword("test", "test", r.ID)
		_, _ = a.Create(rep)

		c := model.NewCategory("Technical literature")
		_, _ = c.Create(rep)
		c = model.NewCategory("Journal")
		_, _ = c.Create(rep)
		c = model.NewCategory("Novel")
		_, _ = c.Create(rep)

		f := model.NewFormat("Phisical copy")
		_, _ = f.Create(rep)
		f = model.NewFormat("Digital copy")
		_, _ = f.Create(rep)
	}
}
