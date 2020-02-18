package data

type Database interface {
	Connect(collection string) error
	Aggregate(query interface{}, result interface{}) (err error)
}
