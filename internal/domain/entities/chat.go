package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chat struct {
	ID               primitive.ObjectID `json:"id" bson:"_id"`
	NameConversation string             `json:"name_conversation,omitempty" bson:"name_conversation,omitempty"`
	// Type: 1 - Private Chat, 2 - Group Chat
	Type      int                  `json:"type" bson:"type"`
	Members   []primitive.ObjectID `json:"members" bson:"members"`
	CreatedAt time.Time            `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time            `json:"updated_at" bson:"updated_at"`
}

type Message struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	// ChatID is the ID of the chat
	ChatID primitive.ObjectID `json:"chat_id" bson:"chat_id"`
	// SenderID is the ID of the user who sent the message
	SenderID          primitive.ObjectID   `json:"sender_id" bson:"sender_id"`
	IsAttach          bool                 `json:"is_attach" bson:"is_attach"`
	Content           string               `json:"content" bson:"content"`
	SendAt            time.Time            `json:"send_at" bson:"send_at"`
	DeleteFromSenders []primitive.ObjectID `json:"delete_from_senders" bson:"delete_from_senders"`
}

type Attachment struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	MessageID primitive.ObjectID `json:"message_id" bson:"message_id"`
	FilePaths []string           `json:"file_paths" bson:"file_paths"`
	IsDelete  bool               `json:"is_delete" bson:"is_delete"`
}
