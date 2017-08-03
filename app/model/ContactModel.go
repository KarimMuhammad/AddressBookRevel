package model

import "fmt"
import (
	"github.com/AddressBookRevelWithCassandra/app"
	"github.com/gocql/gocql"
)

type ContactInfo struct {

	Id gocql.UUID
	FirstName string
	LastName string
	JobTitle string
	Company string
	Email string
	Phones []Phone

}

func (contact ContactInfo) AddContactDB(user_id gocql.UUID) (gocql.UUID, error){

	uuid,err := gocql.RandomUUID()
	err =app.CassandraSession.Query("INSERT INTO contact_by_user( contact_id, first_name, last_name, job , company, email, user_id ) VALUES (?,?,?,?,?,?,?)",uuid,contact.FirstName,contact.LastName,contact.JobTitle,contact.Company,contact.Email,user_id).Exec()
	return uuid,err
}

func (contact ContactInfo) AddNumDB(contactid gocql.UUID) error{

	uuid,err := gocql.RandomUUID()
	err =app.CassandraSession.Query("INSERT INTO phones( number_id,number,contact_id) VALUES (?,?,?)", uuid,contact.Phones[0].PhoneNumber, contactid).Exec()
	return err
}

func  (contact ContactInfo) DeleteContactDB (userid gocql.UUID , contactid gocql.UUID) error{

	deleteContactStat := "DELETE FROM phones WHERE contact_id= ?"
	deletePhonesStat := "DELETE FROM contact_by_user WHERE user_id =? And contact_id=?"
	batch := gocql.NewBatch(gocql.LoggedBatch)
	batch.Query(deleteContactStat, contactid)
	batch.Query(deletePhonesStat, userid, contactid)
	err := app.CassandraSession.ExecuteBatch(batch)
	return err
}


func (contact ContactInfo) GetContacts( user_id gocql.UUID ) ([] ContactInfo){

	var contacts []ContactInfo
	var no Phone

	iter := app.CassandraSession.Query("SELECT contact_id, first_name, last_name, job , company, email FROM contact_by_user  WHERE user_id=? ",user_id).Iter()

	for iter.Scan(&contact.Id, &contact.FirstName, &contact.LastName, &contact.JobTitle, &contact.Company,&contact.Email ){
		iter2:=app.CassandraSession.Query("SELECT Number_id, number FROM phones WHERE contact_id = ? ALLOW FILTERING" , &contact.Id).Iter()
		for iter2.Scan(&no.Id, &no.PhoneNumber) {
			fmt.Println(no)
			contact.Phones = append(contact.Phones, no)
		}
		contacts = append(contacts, contact)
	}

	return contacts
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
