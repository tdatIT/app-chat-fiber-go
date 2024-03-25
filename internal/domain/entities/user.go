package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	Firstname string             `json:"firstname" bson:"firstname"`
	Birthday  time.Time          `json:"birthday" bson:"birthday"`
	Avatar    string             `json:"avatar" bson:"avatar"`
	Email     string             `json:"email" bson:"email,omitempty"`
	Phone     string             `json:"phone" bson:"phone,omitempty"`
	IsActive  bool               `json:"is_active" bson:"is_active"`
	Roles     []string           `json:"roles" bson:"roles"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
