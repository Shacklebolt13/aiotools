package services

import (
	"aiotools/proto"

	"fyne.io/fyne/v2/data/binding"
)

type ShortenClientHandler interface {
	ShortenUrl(longUrlBinding binding.String, shortUrlBinding binding.String) error
	ExpandUrl(shortUrlBinding binding.String, longUrlBinding binding.String) error
}

type shortenClientHandlerImpl struct {
	ShortenClient proto.ShortenerServiceClient
}

func NewShortenClientHandler(shortenClient proto.ShortenerServiceClient) ShortenClientHandler {
	return &shortenClientHandlerImpl{
		ShortenClient: shortenClient,
	}
}

func (s *shortenClientHandlerImpl) ShortenUrl(longUrlBinding binding.String, shortUrlBinding binding.String) error {
	longUrl := longUrlBinding.Get()
	s.ShortenClient.Shorten()
}

func (s *shortenClientHandlerImpl) ExpandUrl(shortUrlBinding binding.String, longUrlBinding binding.String) error {
	return nil
}
