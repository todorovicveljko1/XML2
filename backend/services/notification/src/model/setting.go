package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type OneNotificationSetting struct {
	Type    string `bson:"type"`
	Allowed bool   `bson:"allowed"`
}

type NotificationSetting struct {
	Id       primitive.ObjectID       `bson:"_id"`
	UserId   primitive.ObjectID       `bson:"user_id"`
	Settings []OneNotificationSetting `bson:"settings"`
}
