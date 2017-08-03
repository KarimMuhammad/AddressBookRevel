package model


import (
	"github.com/AddressBookRevelWithCassandra/app"
	"github.com/gocql/gocql"
)

type Phone struct {
	Id gocql.UUID
	PhoneNumber string
}

func (number Phone) AddNumDB(contactid gocql.UUID) (gocql.UUID,error){

	uuid,err := gocql.RandomUUID()
	err =app.CassandraSession.Query("INSERT INTO phones( number_id,number,contact_id) VALUES (?,?,?)", uuid,number.PhoneNumber, contactid).Exec()
	return  uuid,err
}
