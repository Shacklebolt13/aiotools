package services

import (
	"aiotools/src/database/model"
	"aiotools/src/tools"
)

type ShortenService interface {
	Shorten(url string) (string, error)
	Expand(id string) (string, error)
}

type ShortenServiceImpl struct {
	repository model.URLBaseRepository
}

func NewShortenService(repository model.URLBaseRepository) ShortenService {
	return &ShortenServiceImpl{repository: repository}
}

func (service *ShortenServiceImpl) Shorten(url string) (string, error) {
	urlObj, err := service.repository.Insert(url)
	shortUrl := tools.Shorten(urlObj.ID)
	return shortUrl, err
}

func (service *ShortenServiceImpl) Expand(shortUrl string) (string, error) {
	id, err := tools.Expand(shortUrl)
	if err != nil {
		return "", err
	}
	urlObj, err := service.repository.GetById(id)
	return urlObj.Url, err
}
