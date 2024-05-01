package database

import (
	"gorm.io/gorm"
)

type Database struct {
	Conn *gorm.DB
}

func NewDatabase(conn *gorm.DB) *Database {
	return &Database{Conn: conn}
}
