package Services

import (
	"fmt"
	"log"

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

	dresses, err := ds.db.GetDresses()
	if err != nil {
		log.Fatal(err)
	}

	return dresses, nil

}

func (ds *DressService) GetDress(id string) (Dress, error) {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("Failed with error: ", err)
	}

	dress, err := ds.db.GetDress(objID)
	if err != nil {
		log.Fatal(err)
	}

	return dress, nil
}

func (ds *DressService) AddDress(dress *Dress) (string, error) {

	objID := primitive.NewObjectID()

	dress.Id = objID

	errors := validateDress(dress)
	if len(errors) > 0 {
		log.Fatal(errors)
	}

	result, err := ds.db.AddDress(dress)
	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}

func (ds *DressService) UpdateDress(dress *Dress) (string, error) {

	errors := validateDress(dress)
	if len(errors) > 0 {
		log.Fatal(errors)
	}

	fmt.Println("Sending update to dbservice")
	result, err := ds.db.UpdateDress(dress)
	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}

func (ds *DressService) DeleteDress(id string) (string, error) {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("Failed with error: ", err)
	}

	result, err := ds.db.DeleteDress(objID)
	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}

func validateDress(dress *Dress) string {

	var errors string

	if dress.Id.IsZero() {
		errors = errors + "\nId cannot be empty"
	}

	if len(dress.Name) == 0 {
		errors = errors + "\nName cannot be empty"
	}

	if len(errors) > 0 {
		log.Fatal("Posted Dress has errors: ", errors)
	}

	return errors
}
