package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model interface {
	CollectionName() string
	GetId() string
	GetCreatedAt() primitive.DateTime
	GetUpdatedAt() primitive.DateTime
	SetCreatedAt()
	SetUpdatedAt()
}

type BaseModel struct {
	ID        string             `json:"id" bson:"_id,omitempty"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

func (b *BaseModel) GetId() string {
	return b.ID
}
func (b *BaseModel) GetCreatedAt() primitive.DateTime {
	return b.CreatedAt
}

func (b *BaseModel) GetUpdatedAt() primitive.DateTime {
	return b.UpdatedAt
}

func (b *BaseModel) SetCreatedAt() {
	b.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
}

func (b *BaseModel) SetUpdatedAt() {
	b.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
}

// func NewBaseModel() BaseModel {
// 	return BaseModel{
// 		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
// 		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
// 	}
// }
