package messageRepos

import "chat-app-fiber/internal/domain/entities"

type MessageRepository interface {
	// queries
	FindMessageByID(id string) (*entities.Message, error)
	FindAllMessagesByChatID(chatID string, size int, page int) ([]*entities.Message, int, error)
	FindChatByContentInChatID(chatID string, content string, size int, page int) ([]*entities.Message, int, error)
	// commands
	CreateMessage(message *entities.Message) (entities.Message, error)
	DeleteMessageByID(id string) error
}
