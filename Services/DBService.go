package Services

type Database struct {
	// This will be updated to the proper type once the db connection is added
	Client     string
	Name       string
	User       string
	Collection string
	Password   string
}

func NewDatabase() Database {
	return Database{
		Client:     "testClient",
		Name:       "test",
		User:       "test",
		Collection: "test",
		Password:   "test"}
}
