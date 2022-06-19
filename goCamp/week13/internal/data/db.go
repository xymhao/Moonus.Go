package data

type db interface {
	connect() bool
	query() string
}

type MySqlDbCon struct {
}

func NewMySqlDb() db {
	return &MySqlDbCon{}
}

func (db MySqlDbCon) connect() bool {
	return true
}

func (db MySqlDbCon) query() string {
	return "mysql"
}
