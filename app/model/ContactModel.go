package model

import "github.com/gocql/gocql"


import (
	"github.com/AddressBookRevelWithCassandra/app"
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
	numberUuid , err := gocql.RandomUUID()
	contactInsertion := "INSERT INTO contact_by_user( contact_id, first_name, last_name, job , company, email, user_id ) VALUES (?,?,?,?,?,?,?)"
	phoneInsertion := "INSERT INTO phones( number_id,number,contact_id) VALUES (?,?,?)"
	batch := gocql.NewBatch(gocql.LoggedBatch)
	batch.Query(contactInsertion ,uuid,contact.FirstName,contact.LastName,contact.JobTitle,contact.Company,contact.Email,user_id )
	batch.Query(phoneInsertion , numberUuid , contact.Phones[0].PhoneNumber , uuid )
	err =app.CassandraSession.ExecuteBatch(batch)
	return uuid,err
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
	var newcontact ContactInfo
	var no Phone

	rows := app.CassandraSession.Query("SELECT contact_id, first_name, last_name, job , company, email FROM contact_by_user  WHERE user_id=? ",user_id)
	scanner :=rows.Iter().Scanner()
		for scanner.Next(){
	 scanner.Scan(&newcontact.Id, &newcontact.FirstName, &newcontact.LastName, &newcontact.JobTitle, &newcontact.Company,&newcontact.Email )
	 res :=app.CassandraSession.Query("SELECT Number_id, number FROM phones WHERE contact_id = ? " , newcontact.Id)
	 numberScanner :=res.Iter().Scanner()
		for numberScanner.Next(){
			numberScanner.Scan(&no.NoId, &no.PhoneNumber)
			newcontact.Phones = append(newcontact.Phones, no)
		}
		contacts = append(contacts, newcontact)
			newcontact.Phones = []Phone{}

	}
	rows.Iter().Close()
	return contacts
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
