package gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

type ConnectionValues struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
}

type Database struct {
	DB               *gorm.DB
	connectionValues ConnectionValues
	entities         []interface{}
}

var instance *Database
var once sync.Once

func NewDatabase(connectionValues ConnectionValues, entities []interface{}) *Database {
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", connectionValues.User, connectionValues.Password, connectionValues.Host, connectionValues.Port, connectionValues.DatabaseName)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic("failed to connect to database" + err.Error())
		}

		err = db.AutoMigrate(entities...)

		if err != nil {
			panic("failed to migrate entities to database" + err.Error())
		}

		log.Println("Database connected")
		instance = &Database{DB: db}
	})
	return instance
}
