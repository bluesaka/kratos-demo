package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"kratos-demo/app/order/service/internal/conf"
)

var ProviderSet = wire.NewSet(NewData, NewDB, NewOrderRepo)

type Data struct {
	db  *gorm.DB
	log *log.Helper
}

func NewDB(c *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "order-service/data/gorm"))
	log.Info("order-service connect to mysql")

	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("db error")
	}

	return db
}

func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "order-service/data"))
	return &Data{
			db:  db,
			log: log,
		}, func() {
			log.Info("order-service/data exit")
		}, nil
}
