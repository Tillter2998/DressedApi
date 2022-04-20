package Models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dress struct {
	Id          primitive.ObjectID `bson:"_id" json:"Id,omitempty"`
	Price       float32            `bson:"price" json:"Price"`
	Name        string             `bson:"name" json:"Name"`
	Description string             `bson:"description" json:"Description"`
}
