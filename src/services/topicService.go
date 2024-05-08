package services

import (
	"aiotools/proto"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type topicCollection struct {
	messageStream *chan string
}

type TopicService interface {
	Subscribe(id uuid.UUID, subServer proto.PubSubService_SubscribeServer) error
	Publish(id uuid.UUID, message string) error
	CreateTopic() (uuid.UUID, error)
}

type topicServiceImpl struct {
	topics map[uuid.UUID]*topicCollection
	wg     sync.WaitGroup
}

func (service *topicServiceImpl) CreateTopic() (uuid.UUID, error) {
	id := uuid.New()
	topicChan := make(chan string)
	topic := &topicCollection{
		messageStream: &topicChan,
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

	for message := range *topic.messageStream {
		if err := subServer.Send(&proto.SubscribeResponse{Message: message}); err != nil {
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

	*topic.messageStream <- message
	return nil
}

func expireTopic(tCol topicCollection, duration time.Duration) {
	time.Sleep(duration)
	close(*tCol.messageStream)
}

func NewTopicService() TopicService {
	return &topicServiceImpl{
		topics: make(map[uuid.UUID]*topicCollection),
		wg:     sync.WaitGroup{},
	}
}
