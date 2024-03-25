package messageRepos

import (
	"chat-app-fiber/internal/domain/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

type messageRepository struct {
	messageCollection *mongo.Collection
}

func NewMessageRepository(client *mongo.Client) MessageRepository {
	messageCollection := client.Database("chat").Collection("messages")
	return &messageRepository{messageCollection: messageCollection}
}

func (m messageRepository) FindMessageByID(id string) (*entities.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m messageRepository) FindAllMessagesByChatID(chatID string, size int, page int) ([]*entities.Message, int, error) {
	//TODO implement me
	panic("implement me")
}

func (m messageRepository) FindChatByContentInChatID(chatID string, content string, size int, page int) ([]*entities.Message, int, error) {
	//TODO implement me
	panic("implement me")
}

func (m messageRepository) CreateMessage(message *entities.Message) (entities.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m messageRepository) DeleteMessageByID(id string) error {
	//TODO implement me
	panic("implement me")
}
