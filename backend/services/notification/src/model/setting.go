package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type OneNotificationSetting struct {
	Type    string `bson:"type,omitempty"`
	Allowed bool   `bson:"allowed,omitempty"`
}

type NotificationSetting struct {
	Id       primitive.ObjectID       `bson:"_id,omitempty"`
	UserId   primitive.ObjectID       `bson:"user_id,omitempty"`
	Settings []OneNotificationSetting `bson:"settings,omitempty"`
}
