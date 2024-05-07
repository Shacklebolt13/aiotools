package handlers

import (
	pb "aiotools/proto"
	"context"
	"log"

	"aiotools/src/services"
)

type ShortenerServiceHandler struct {
	pb.UnimplementedShortenerServiceServer
	shortenService services.ShortenService
}

func (handler *ShortenerServiceHandler) Shorten(ctx context.Context, in *pb.ShortenRequest) (*pb.ShortenResponse, error) {
	log.Printf("Received: %v", in.GetUrl())
	shortUrl, err := handler.shortenService.Shorten(in.GetUrl())
	return &pb.ShortenResponse{Id: shortUrl}, err
}

func (handler *ShortenerServiceHandler) Expand(ctx context.Context, in *pb.ExpandRequest) (*pb.ExpandResponse, error) {
	log.Printf("Received: %v", in.GetId())
	urlObj, err := handler.shortenService.Expand(in.GetId())
	return &pb.ExpandResponse{Url: urlObj}, err
}

func NewShortenerServiceHandler(shortenService services.ShortenService) pb.ShortenerServiceServer {
	return &ShortenerServiceHandler{shortenService: shortenService}
}
