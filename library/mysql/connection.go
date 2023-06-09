package mysql

import (
	//"boilerplate/util/logwrapper"
	"boilerplate/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Client *sql.DB

func NewConnection(Mysqlcongif config.MySqlConfig) error {

	// Open up our database connection.
	// I've set up a database on my local machine using Docker.

	//logwrapper.Logger.Debugln(Mysqlcongif)
	db, err := sql.Open("mysql", Mysqlcongif.URL+"/"+Mysqlcongif.Database)

	// if there is an error opening the connection, handle it
	if err != nil {
		// panic(err.Error())
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	//logwrapper.Logger.Infoln("Connected to MYSQL_URL : ", Mysqlcongif.URL)
	Client = db
	return nil
}
