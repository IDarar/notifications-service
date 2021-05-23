package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Templates for different notifications
const (
	ForumNotification  = "Wrote [%v] to you in [%v]" //v will refer to post, in ... to topic
	FirstFriendRequest = "Added you to friendlist. Add to your friend list as well?"
	FriendRequest      = "Added you to friendlist"                            //if user is alreaydfi in your friend list
	Ban                = "You've been banned for [%v] for commenting in [%v]" //v -> time, in -> post
)

type Notification struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID    int                `json:"user_id,omitempty" bson:"user_id,omitempty"` //owner
	From      int                `json:"from,omitempty" bson:"from,omitempty"`       //userID
	Topic     string             `json:"topic,omitempty" bson:"topic,omitempty"`     //ID of topic where mentioned
	Content   string             `json:"content,omitempty" bson:"content,omitempty"`
	Text      string             `json:"text,omitempty" bson:"text,omitempty"` // text according to type contating id's of what and where
	Type      string             `json:"type,omitempty" bson:"type,omitempty"` // forum notification, friend request, forum ban etc
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	IsRead    bool               `json:"is_read,omitempty" bson:"is_read,omitempty"`
}
