package model

import "fmt"
import (
	"github.com/AddressBookRevel/app"
	"strconv"
)

type ContactInfo struct {

	Id int
	FirstName string
	LastName string
	JobTitle string
	Company string
	Email string
	Phones []Phone

}

func (contact ContactInfo) AddContactDB(user_id int64)(int64){

	fmt.Println(user_id)
	contactquery, err :=app.DB.Exec("INSERT INTO Contact_List( Contact_ID, First_Name, Last_Name, Job_Title, Company, Email, User_Id_FK) VALUES (?,?,?,?,?,?,?)",nil,contact.FirstName,contact.LastName,contact.JobTitle,contact.Company,contact.Email,user_id)
	checkErr(err)
	lastinsertid,err := contactquery.LastInsertId()
	checkErr(err)
	return  lastinsertid
}

func (contact ContactInfo) AddNumDB(contactid int64){

	_, err := app.DB.Exec("INSERT INTO Phone_Numbers( No_ID,Phone_Number,Contact_Id_FK) VALUES (?,?,?)", nil,contact.Phones[0].PhoneNumber, contactid)
	checkErr(err)
}

func  (contact ContactInfo) DeleteContactDB (contactid string){

	ID, _ := strconv.ParseInt(contactid, 10, 64)
	_,err := app.DB.Exec("DELETE FROM Phone_Numbers WHERE Contact_Id_FK=?",ID)
	checkErr(err)
	_,err = app.DB.Exec("DELETE FROM Contact_List WHERE Contact_ID=?",ID)
	checkErr(err)
}


func (contact ContactInfo) HomeDB( user_id int64 ) [] ContactInfo{

	contactrows, err := app.DB.Query("SELECT Contact_ID, First_Name, Last_Name, Job_Title, Company, Email, No_ID, Phone_Number FROM Contact_List join`Phone_Numbers` ON Contact_List.Contact_ID = Phone_Numbers.Contact_Id_FK WHERE User_Id_FK=? ",user_id)
	checkErr(err)
	var contacts []ContactInfo
	var currentcontact ContactInfo
	var no Phone

	for contactrows.Next(){

		contactrows.Scan(&contact.Id, &contact.FirstName, &contact.LastName, &contact.JobTitle, &contact.Company,&contact.Email ,&no.Id,&no.PhoneNumber)
		if contact.Id!=currentcontact.Id && currentcontact.Id != 0{

			contacts = append(contacts, currentcontact)
			currentcontact = contact

		}else if currentcontact.Id == 0{

			currentcontact=contact
		}
		currentcontact.Phones = append(currentcontact.Phones, no)
	}
	contacts = append(contacts, currentcontact)
	fmt.Println(contacts)
	return contacts
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
