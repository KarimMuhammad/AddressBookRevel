package model


import (
	"github.com/AddressBookRevelWithCassandra/app"
	"github.com/gocql/gocql"
)

type Phone struct {
	NoId gocql.UUID
	PhoneNumber string
}

func (number Phone) AddNumDB(contactid gocql.UUID) (gocql.UUID,error){

	uuid,err := gocql.RandomUUID()
	err =app.CassandraSession.Query("INSERT INTO phones( number_id,number,contact_id) VALUES (?,?,?)", uuid,number.PhoneNumber, contactid).Exec()
	return  uuid,err
}

func (number Phone) DeleteNumDB(contactid gocql.UUID , numberid gocql.UUID) (error){

	err :=app.CassandraSession.Query("DELETE FROM phones WHERE contact_id= ? and number_id = ?", contactid, numberid).Exec()
	return  err
}
