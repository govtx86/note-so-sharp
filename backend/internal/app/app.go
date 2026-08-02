package app

import (
	"log"
	"nss/internal/store"
	"os"

	"gorm.io/gorm"
)

type Application struct {
	Logger *log.Logger
	DB *gorm.DB
}


func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	pgDB, err := store.Open(logger)
	if err != nil {
		return nil, err
	}

	app := &Application{
		Logger: logger,
		DB: pgDB,
	}
	return app, nil
}