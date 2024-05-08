package handlers

import (
	"aiotools/proto"
	"aiotools/src/services"
	"context"

	"github.com/google/uuid"
)

type PubSubServiceHandler struct {
	proto.UnimplementedPubSubServiceServer
	topicService services.TopicService
}

func NewPubSubServiceHandler(
	topicService services.TopicService) proto.PubSubServiceServer {
	return &PubSubServiceHandler{
		topicService: topicService,
	}
}

func (pubSubHandler *PubSubServiceHandler) CreateTopic(ctx context.Context, CreateTopicRequest *proto.CreateTopicRequest) (*proto.CreateTopicResponse, error) {
	id, err := pubSubHandler.topicService.CreateTopic()
	if err != nil {
		return nil, err
	}
	return &proto.CreateTopicResponse{Topic: id.String()}, nil
}

func (pubSubHandler *PubSubServiceHandler) Subscribe(subRequest *proto.SubscribeRequest, subServer proto.PubSubService_SubscribeServer) error {
	id, err := uuid.Parse(subRequest.Topic)
	if err != nil {
		return err
	}
	return pubSubHandler.topicService.Subscribe(id, subServer)
}

func (pubSubHandler *PubSubServiceHandler) Publish(ctx context.Context, publishRequest *proto.PublishRequest) (*proto.PublishResponse, error) {
	id, err := uuid.Parse(publishRequest.Topic)

	if err != nil {
		return nil, err
	}
	return &proto.PublishResponse{}, pubSubHandler.topicService.Publish(id, publishRequest.Message)
}
