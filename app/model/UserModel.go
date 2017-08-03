package model

import "golang.org/x/crypto/bcrypt"
import (
	"github.com/AddressBookRevelWithCassandra/app"
	"github.com/gocql/gocql"
)


type User struct {
	Username string
	Email string
	Password string
}


func (info User) LoginDB() bool{

	var dbPasswordHash []byte
	if(info.Username !="") {
		err := app.CassandraSession.Query("SELECT password FROM user_by_name WHERE user_name=? ", info.Username).Scan(&dbPasswordHash)
		return err == nil
	}
	err := bcrypt.CompareHashAndPassword(dbPasswordHash, []byte(info.Password))
	return err == nil
}

func (info User) Exists() bool {

	var id gocql.UUID
	err := app.CassandraSession.Query("SELECT user_id FROM user_by_name WHERE user_name=? " ,info.Username).Scan(&id)
	if err !=nil{
		return false
	}
	if len(id) > 0{
		return true
	}else {
		return false
	}
}

func (signinfo User) SignupDB()error{
	securedpass,err := bcrypt.GenerateFromPassword([]byte(signinfo.Password),bcrypt.DefaultCost)
	uuid,err := gocql.RandomUUID()
	checkErr(err)
	err = app.CassandraSession.Query("INSERT INTO user_by_name(user_id,user_name,email,password) VALUES (?, ?, ?, ?)",
		uuid, signinfo.Username, signinfo.Email,securedpass).Exec()
	return err
}

