package services

import (
	pb "aiotools/proto"
	"context"
	"log"

	database "aiotools/src/database"
	model "aiotools/src/database/model"
	tools "aiotools/src/tools"
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
	shortUrl := tools.Shorten(urlObj.ID)
	return &pb.ShortenResponse{Id: shortUrl}, err
}

func (service *shortenerService) Expand(ctx context.Context, in *pb.ExpandRequest) (*pb.ExpandResponse, error) {
	log.Printf("Received: %v", in.GetId())
	id, err := tools.Expand(in.GetId())
	if err != nil {
		return nil, err
	}
	urlObj, err := service.urlRepo.GetById(id)
	return &pb.ExpandResponse{Url: urlObj.Url}, err
}
