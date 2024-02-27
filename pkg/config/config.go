package config

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func CreateDBConnection(dbConfig string) *gorm.DB {
	DB, err := gorm.Open(mysql.Open(dbConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	dbO, err := DB.DB()
	if err != nil {
		panic(err)
	}

	dbO.SetConnMaxIdleTime(time.Duration(1) * time.Minute)
	dbO.SetMaxIdleConns(2)
	dbO.SetConnMaxLifetime(time.Duration(1) * time.Minute)

	DB.Statement.RaiseErrorOnNotFound = true

	if err != nil {
		panic(err)
	}
	return DB
}
