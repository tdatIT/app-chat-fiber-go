package chatRepos

import (
	"chat-app-fiber/internal/domain/entities"
	"chat-app-fiber/internal/utils/db/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type chatRepository struct {
	chatCollection *mongo.Collection
}

func NewChatRepository(client *mongodb.MongoClient) ChatRepository {
	chatCollection := client.GetDB().Collection("chats")
	return &chatRepository{chatCollection: chatCollection}
}

func (c chatRepository) FindChatByID(id string) (*entities.Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (c chatRepository) FindChatByMemberID(memberID string) ([]*entities.Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (c chatRepository) FindChatByMemberIDs(memberIDs []string) ([]*entities.Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (c chatRepository) CreateChat(chat *entities.Chat) (entities.Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (c chatRepository) LeaveChat(chatID string, memberID string) error {
	//TODO implement me
	panic("implement me")
}
