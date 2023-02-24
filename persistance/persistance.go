package persistance

type Persistance struct {
	db *MongoDB
}

func NewPersistance(dbConfig MongoDBConfig) *Persistance {
	p := &Persistance{}
	p.db = NewMongoDB(dbConfig)
	return p
}
