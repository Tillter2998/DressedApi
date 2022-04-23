package Services

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dress struct {
	Id          primitive.ObjectID `bson:"_id" json:"Id,omitempty"`
	Price       float64            `bson:"price" json:"Price"`
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

	var dresses []*Dress
	fmt.Println("Getting Collection...")
	coll := ds.db.Client.Database(ds.db.Collection).Collection(ds.db.Documents)

	fmt.Println("Getting documents...")
	cursor, err := coll.Find(ds.db.Context, bson.M{})
	if err != nil {
		log.Fatal("Getting documents failed with error: ", err)
	}
	fmt.Println("Documents successfully retrieved")
	fmt.Println(cursor.RemainingBatchLength())

	for cursor.Next(ds.db.Context) {
		var dress *Dress
		if err = cursor.Decode(&dress); err != nil {
			log.Fatal("Failed to decode dress with error: ", err)
		}
		dresses = append(dresses, dress)
		fmt.Println(len(dresses))
	}

	return dresses, nil
}
