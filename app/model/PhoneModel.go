package model

import "strconv"
import "github.com/AddressBookRevel/app"

type Phone struct {
	Id int
	PhoneNumber string
}

func (number Phone) AddNumDB(contactid string) (int64){

	ID, _ := strconv.ParseInt(contactid, 10, 64)
	add, err := app.DB.Exec("INSERT INTO Phone_Numbers( No_ID,Phone_Number,Contact_Id_FK) VALUES (?,?,?)", nil,number.PhoneNumber, ID)
	checkErr(err)
	id , err := add.LastInsertId()
	checkErr(err)
	return  id
}
