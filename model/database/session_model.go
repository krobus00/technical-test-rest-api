package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type Session struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string             `bson:"username"`
	RefreshToken string             `bson:"refresh_token"`
	IsBlocked    bool               `bson:"is_blocked"`
	DateColumn
}
