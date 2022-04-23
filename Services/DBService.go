package Services

import (
	c "DressedApi/Config"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
