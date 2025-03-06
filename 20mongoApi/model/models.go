package model

import (
	"go.mongodb.org/mongo-driver/v2/bson/primitive"
)

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson`
	Movie   string             `json:"movie,omitempty"`
	watched bool               `json:"watched,omitempty"`
}
