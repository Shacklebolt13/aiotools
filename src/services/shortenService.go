package services

import (
	pb "aiotools/proto"
	"context"
	"log"

	database "aiotools/src/database"
	model "aiotools/src/database/model"
)

type shortenerService struct {
	pb.UnimplementedShortenerServiceServer
	urlRepo model.URLBaseRepository
}

func NewShortenerService(db *database.Database) *shortenerService {
	return &shortenerService{urlRepo: model.NewURLBaseRepository(db)}
}

func (service *shortenerService) Shorten(ctx context.Context, in *pb.ShortenRequest) (*pb.ShortenResponse, error) {
	log.Printf("Received: %v", in.GetUrl())
	urlObj, err := service.urlRepo.Insert(in.GetUrl())
	return &pb.ShortenResponse{Id: urlObj.ID}, err
}
