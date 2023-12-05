package models

import "time"

type FileTransaction struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    string    `bson:"user_id" json:"user_id"`
	FileName  string    `bson:"file_name" json:"file_name"`
	Operation string    `bson:"operation" json:"operation"` // "upload" or "download"
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	FileSize  int64     `bson:"file_size" json:"file_size"`
	Format    string    `bson:"format,omitempty" json:"format,omitempty"`
}
