package register

import (
	MysqlWrapper "boilerplate/library/mysql"
	"boilerplate/util/logwrapper"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Register struct {
	ID        int    `json:"id"`
	Username  string `json:"Username" validate:"required,min=4,max=30,customUsernameValidation"`
	Firstname string `json:"Firstname" validate:"required,alpha,customFirstnameValidation"`
	Lastname  string `json:"Lastname"`
	Email     string `json:"Email" validate:"required,email"`
	Pan       string `json:"Pan"`
	Adhar     string `json:"Adhar"`
	Mobile    string `json:"Mobile" validate:"required,len=10|len=12`
	Password  string `json:"Password" validate:"required,min=6,max=15,customPasswordValidation"`
}
type Login struct {
	Username string `json:"Username" validate:"required,min=4,max=30,customUsernameValidation"`
	Password string `json:"Password" validate:"required,min=8,alphanum"`
}

func createUserTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS registers 
		ID INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(255) NOT NULL,
		firstname varchar(255) not null,
		lastname varchar(255),
		email VARCHAR(255) NOT NULL,
		pan varchar (30),
		adhar varchar(30),
		mobile varchar(12) not null,
		password VARCHAR(255) NOT null
	)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	rows, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating album table", err)
		return err
	}
	log.Printf("Rows affected when creating table: %d", rows)
	return nil
}

// Query Executer For user info Insert
func InsertUser(s Register) (bool, error) {
	fmt.Println("request pay load ", s)
	/*err := createUserTable(MysqlWrapper.Client)
	if err != nil {
		log.Printf("Create register table failed with error %s", err)
		return false, err
	}*/
	args := []interface{}{s.Username, s.Firstname, s.Lastname, s.Email, s.Pan, s.Adhar, s.Mobile, s.Password}
	_, q_err := MysqlWrapper.Client.Exec("INSERT INTO registers (username, firstname, lastname, email, pan, adhar, mobile, password) VALUES (?, ?, ?, ?, ?, ?, ?,?);", args...)

	// if there is an error inserting, handle it

	if q_err != nil {
		logwrapper.Logger.Debugln(q_err)
		return false, q_err
	}

	return true, nil
}
