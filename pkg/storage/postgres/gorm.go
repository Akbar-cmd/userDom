package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Db struct {
	*gorm.DB
}

func NewDB(dsn string) (*Db, error) {
	const op = "storage.gorm.New" // package and func name
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// wrapper error
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	log.Println("База данных создана")
	return &Db{db}, nil
}
