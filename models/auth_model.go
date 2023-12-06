package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type LogInUser struct {
	ID            primitive.ObjectID `bson:"_id"`
	Password      *string            `bson:"password,omitempty" json:"password,omitempty" validate:"required,min=6"`
	Email         *string            `bson:"email" json:"email" validate:"email,required"`
	Token         *string            `bson:"token" json:"token"`
	Refresh_token *string            `bson:"refresh_token" json:"refresh_token"`
	User_id       string             `bson:"user_id" json:"user_id"`
	// Created_at    *time.Time         `json:"Created_at,omitempty"`
	// Updated_at    *time.Time         `json:"updated_at,omitempty"`
}

type SignUpUser struct {
	ID            primitive.ObjectID `bson:"_id"`
	Name          *string            `bson:"name" json:"name" validate:"required,min=2,max=150"`
	Password      *string            `bson:"password,omitempty" json:"password,omitempty" validate:"required,min=6"`
	Email         *string            `bson:"email" json:"email" validate:"email,required"`
	Phone         *string            `bson:"phone" json:"phone" validate:"required"`
	Token         *string            `bson:"token" json:"token"`
	Refresh_token *string            `bson:"refresh_token" json:"refresh_token"`
	User_id       string             `bson:"user_id" json:"user_id"`
	// Created_at    *time.Time         `json:"Created_at,omitempty"`
	// Updated_at    *time.Time         `json:"updated_at,omitempty"`
}
