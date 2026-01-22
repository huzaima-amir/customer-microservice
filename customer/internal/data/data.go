package data

import (
	"customer/internal/conf"
	"gorm.io/driver/postgres"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewCustomerRepo)

// Data 
type Data struct {
	db *gorm.DB
}

// NewData 
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
    log := log.NewHelper(logger)

    // connect to PostgreSQL
    db, err := gorm.Open(postgres.Open(c.Database.Source), &gorm.Config{})
    if err != nil {
        return nil, nil, err
    }

    // auto-migrate models
    if err := db.AutoMigrate(
        &Customer{},
        &Email{},
        &PhoneNumber{},
        &Address{},
    ); err != nil {
        return nil, nil, err
    }

    cleanup := func() {
        log.Info("closing the data resources")
        sqlDB, _ := db.DB()
        sqlDB.Close()
    }

    return &Data{db: db}, cleanup, nil
}

