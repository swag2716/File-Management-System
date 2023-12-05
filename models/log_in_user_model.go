package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type LogInUser struct {
	ID            primitive.ObjectID `bson:"_id"`
	Password      *string            `json:"password,omitempty" validate:"required,min=6"`
	Email         *string            `json:"email" validate:"email,required"`
	Token         *string            `json:"token"`
	Refresh_token *string            `json:"refresh_token"`
	User_id       string             `json:"user_id"`
	// Created_at    *time.Time         `json:"Created_at,omitempty"`
	// Updated_at    *time.Time         `json:"updated_at,omitempty"`
}
