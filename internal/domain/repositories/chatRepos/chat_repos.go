package chatRepos

import "chat-app-fiber/internal/domain/entities"

type ChatRepository interface {
	// queries
	FindChatByID(id string) (*entities.Chat, error)
	FindChatByMemberID(memberID string) ([]*entities.Chat, error)
	FindChatByMemberIDs(memberIDs []string) ([]*entities.Chat, error)
	//commands
	CreateChat(chat *entities.Chat) (entities.Chat, error)
	LeaveChat(chatID string, memberID string) error
}
