package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignUpUser struct {
	ID            primitive.ObjectID `bson:"_id"`
	Name          *string            `json:"name" validate:"required,min=2,max=150"`
	Password      *string            `json:"password,omitempty" validate:"required,min=6"`
	Email         *string            `json:"email" validate:"email,required"`
	Phone         *string            `json:"phone" validate:"required"`
	Token         *string            `json:"token"`
	Refresh_token *string            `json:"refresh_token"`
	User_id       string             `json:"user_id"`
	// Created_at    *time.Time         `json:"Created_at,omitempty"`
	// Updated_at    *time.Time         `json:"updated_at,omitempty"`
}
