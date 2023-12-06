package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileRecord struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FileId    string             `bson:"file_id" json:"file_id"`
	UserID    string             `bson:"user_id" json:"user_id"`
	FileName  string             `bson:"file_name" json:"file_name"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
	FileSize  int64              `bson:"file_size" json:"file_size"`
	Format    string             `bson:"format,omitempty" json:"format,omitempty"`
}
