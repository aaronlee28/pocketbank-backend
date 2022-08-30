package db

import (
	"fmt"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var (
	c  = config.Config.DBConfig
	db *gorm.DB
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
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", c.Host, c.User, c.Password, c.DBName, c.Port)
	//dsn := fmt.Sprintf("host=#{c.Host} user=#{c.User} password=#{c.Password} dbname=#{c.DBName} port=#{c.Port} sslmode=disable TimeZone=Asia/Jakarta", c.Host, c.User, c.Password, c.DBName, c.Port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: getLogger(),
	})
	return err
}

func Get() *gorm.DB {
	return db
}
