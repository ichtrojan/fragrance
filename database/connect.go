package database

import (
	"errors"
	"fmt"
	"github.com/ichtrojan/thoth"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Init() *gorm.DB {
	logger, _ := thoth.Init("log")

	user, exist := os.LookupEnv("DB_USER")

	if !exist {
		logger.Log(errors.New("DB_USER not set in .env"))
		log.Fatal("DB_USER not set in .env")
	}

	pass, exist := os.LookupEnv("DB_PASS")

	if !exist {
		logger.Log(errors.New("DB_PASS not set in .env"))
		log.Fatal("DB_PASS not set in .env")
	}

	host, exist := os.LookupEnv("DB_HOST")

	if !exist {
		logger.Log(errors.New("DB_HOST not set in .env"))
		log.Fatal("DB_HOST not set in .env")
	}

	name, exist := os.LookupEnv("DB_NAME")

	if !exist {
		logger.Log(errors.New("DB_NAME not set in .env"))
		log.Fatal("DB_NAME not set in .env")
	}

	credentials := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, name)

	db, err := gorm.Open("mysql", credentials)

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	}

	return db
}
