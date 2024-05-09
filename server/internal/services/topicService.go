package services

import (
	"aiotools/proto"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type topic struct {
	connections []*chan string
}

type TopicService interface {
	Subscribe(id uuid.UUID, subServer proto.PubSubService_SubscribeServer) error
	Publish(id uuid.UUID, message string) error
	CreateTopic() (uuid.UUID, error)
}

type topicServiceImpl struct {
	topics map[uuid.UUID]*topic
}

func (service *topicServiceImpl) CreateTopic() (uuid.UUID, error) {
	id := uuid.New()
	topic := &topic{
		connections: make([]*chan string, 0),
	}
	service.topics[id] = topic
	go expireTopic(*topic, 30*time.Minute)
	return id, nil
}

func (service *topicServiceImpl) Subscribe(id uuid.UUID, subServer proto.PubSubService_SubscribeServer) error {

	topic, ok := service.topics[id]
	if !ok {
		return fmt.Errorf("topic with id %s not found", id)
	}

	subChan := make(chan string)
	topic.connections = append(topic.connections, &subChan)

	for message := range subChan {
		if err := subServer.Send(&proto.SubscribeResponse{Message: message}); err != nil {
			close(subChan)
			return err
		}
	}

	return nil
}

func (service *topicServiceImpl) Publish(id uuid.UUID, message string) error {
	topic, ok := service.topics[id]
	if !ok {
		return fmt.Errorf("topic with id %s not found", id)
	}
	for _, conn := range topic.connections {
		*conn <- message
	}
	return nil
}

func expireTopic(tCol topic, duration time.Duration) {
	time.Sleep(duration)
	for _, conn := range tCol.connections {
		close(*conn)
	}
}

func NewTopicService() TopicService {
	return &topicServiceImpl{
		topics: make(map[uuid.UUID]*topic),
	}
}
