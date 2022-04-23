package Services

import (
	c "DressedApi/Config"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IDatabaseService interface {
	GetDresses() ([]*Dress, error)
	GetDress(id primitive.ObjectID) (Dress, error)
	AddDress(dress *Dress) (string, error)
	UpdateDress(dress *Dress) (string, error)
	DeleteDress(id primitive.ObjectID) (string, error)
}

type Database struct {
	// Client will be updated to the proper type once the db connection is added
	Client     *mongo.Client
	Context    context.Context
	Name       string
	Collection string
	Documents  string
}

func NewDatabase(config *c.Configuration) Database {

	dbUri := fmt.Sprintf(
		"mongodb+srv://%v:%v@%v.iri9a.mongodb.net/%v?retryWrites=true&w=majority",
		config.DB_USERNAME, config.DB_PASSWORD, config.DB_NAME, config.DB_COLLECTION)

	// fmt.Println(dbUri)

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().
		ApplyURI(dbUri).SetServerAPIOptions(serverAPIOptions)

	ctx := context.Background()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to db successfully")

	return Database{
		Client:     client,
		Context:    ctx,
		Name:       config.DB_NAME,
		Collection: config.DB_COLLECTION,
		Documents:  config.DB_DOCUMENTS}
}

func (db *Database) GetDresses() ([]*Dress, error) {

	var dresses []*Dress

	coll := getCollection(db)

	fmt.Println("Getting documents...")
	cursor, err := coll.Find(db.Context, bson.M{})
	if err != nil {
		log.Fatal("Getting documents failed with error: ", err)
	}
	fmt.Println("Documents successfully retrieved")

	for cursor.Next(db.Context) {
		var dress *Dress
		if err = cursor.Decode(&dress); err != nil {
			log.Fatal("Failed to decode dress with error: ", err)
		}
		dresses = append(dresses, dress)
	}

	return dresses, nil
}

func (db *Database) GetDress(id primitive.ObjectID) (Dress, error) {

	var dress Dress

	coll := getCollection(db)

	fmt.Println("Getting dress with id: ", id)
	result := coll.FindOne(db.Context, bson.M{"_id": id})
	if err := result.Decode(&dress); err != nil {
		log.Fatal("Failed to decode dress with error: ", err)
	}
	return dress, nil
}

func (db *Database) AddDress(dress *Dress) (string, error) {

	coll := getCollection(db)

	fmt.Println("Adding dress with Id: ", dress.Id)
	result, err := coll.InsertOne(db.Context, dress)
	if err != nil {
		log.Fatal("failed inserting with error: ", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func getCollection(db *Database) *mongo.Collection {

	fmt.Println("Getting Collection...")
	coll := db.Client.Database(db.Collection).Collection(db.Documents)

	return coll
}
