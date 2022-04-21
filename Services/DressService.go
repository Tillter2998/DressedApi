package Services

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dress struct {
	Id          primitive.ObjectID `bson:"_id" json:"Id,omitempty"`
	Price       float32            `bson:"price" json:"Price"`
	Name        string             `bson:"name" json:"Name"`
	Description string             `bson:"description" json:"Description"`
}

type IDressService interface {
	GetDresses() ([]*Dress, error)
	GetDress(id string) (Dress, error)
	AddDress(dress *Dress) (string, error)
	UpdateDress(dress *Dress) (string, error)
	DeleteDress(id string) (string, error)
}

type DressService struct {
	db *Database
}

func NewDressService(database *Database) DressService {
	return DressService{db: database}
}

func (ds *DressService) GetDresses() ([]*Dress, error) {
	fmt.Println(ds.db.Client)
	fmt.Println(ds.db.Collection)
	fmt.Println(ds.db.Name)
	fmt.Println(ds.db.Password)
	fmt.Println(ds.db.User)
	var dresses = []*Dress{
		{Id: primitive.NewObjectID(), Price: 10.99, Name: "Dress 1", Description: "First dress"},
		{Id: primitive.NewObjectID(), Price: 9.99, Name: "Dress 2", Description: "Second Dress"},
	}

	return dresses, nil
}
