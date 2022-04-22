package Services

import c "DressedApi/Config"

type Database struct {
	// Client will be updated to the proper type once the db connection is added
	Client     string
	Name       string
	User       string
	Collection string
	Password   string
}

func NewDatabase(config *c.Configuration) Database {
	return Database{
		Client:     "testClient",
		Name:       "test",
		User:       "test",
		Collection: "test",
		Password:   "test"}
}
