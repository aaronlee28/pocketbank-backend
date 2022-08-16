package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var db *gorm.DB
var (
	host   = "localhost"
	port   = "5432"
	user   = "postgres"
	pwd    = "postgres"
	dbName = "wallet_db_aaronlee"
)

func getLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)
}

func Connect() (err error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", host, user, pwd, dbName, port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: getLogger(),
	})

	return err
}

func Get() *gorm.DB {
	return db
}
