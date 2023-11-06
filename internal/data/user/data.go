package user

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kratos-demo/internal/conf"
	"kratos-demo/internal/proto/protouser"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewUserRepo,
	wire.Bind(new(protouser.IUserRepo), new(*UserRepo)),
)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data) (*Data, func(), error) {
	cleanup := func() {
		log.Info("closing the data resources")
	}

	db, err := NewMySQL(c)
	if err != nil {
		return nil, nil, err
	}

	return &Data{db: db}, cleanup, nil
}

func NewMySQL(c *conf.Data) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	return db, err
}
