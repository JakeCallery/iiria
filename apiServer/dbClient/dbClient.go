package dbClient

type DbClient interface {
	Init()
	DataFromTime(string) (string, error)
	CheckConnection() error
}
