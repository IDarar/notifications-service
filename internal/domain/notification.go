package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID    int                `json:"user_id,omitempty" bson:"user_id,omitempty"`
	From      int                `json:"from,omitempty" bson:"from,omitempty"`       //userID
	Topic     string             `json:"topic,omitempty" bson:"topic,omitempty"`     //ID of topic where mentioned
	Content   string             `json:"content,omitempty" bson:"content,omitempty"` //according to type
	Type      string             `json:"type,omitempty" bson:"type,omitempty"`       // forum notification, friend request, forum ban etc
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	IsRead    bool               `json:"is_read,omitempty" bson:"is_read,omitempty"`
}
