package model

import (
	"log"

	database "aiotools/server/internal/database"

	"gorm.io/gorm"
)

type URLBase struct {
	gorm.Model
	Url string
}

type URLBaseRepository interface {
	GetById(id uint) (*URLBase, error)
	Insert(url string) (*URLBase, error)
}

type urlBaseRepositoryImpl struct {
	*database.Database
}

func NewURLBaseRepository(database *database.Database) URLBaseRepository {
	return &urlBaseRepositoryImpl{database}
}

func (repo *urlBaseRepositoryImpl) GetById(id uint) (*URLBase, error) {
	urlMap := &URLBase{}
	transaction := repo.Conn.First(urlMap, id)
	if transaction.Error != nil {
		log.Default().Printf("Error querying database: %v", transaction.Error)
		return nil, transaction.Error
	}
	return urlMap, nil
}

func (repo *urlBaseRepositoryImpl) Insert(url string) (*URLBase, error) {
	urlMap := &URLBase{Url: url}
	transaction := repo.Conn.Create(urlMap)
	if transaction.Error != nil {
		log.Default().Printf("Error inserting into database: %v", transaction.Error)
		return nil, transaction.Error
	}
	return urlMap, nil
}
