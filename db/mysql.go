package db

var dbname string

func init() {

	dbname = "Mysql"
}

func GetDbName() string {
	return dbname
}
