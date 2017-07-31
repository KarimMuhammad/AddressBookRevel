package model

import "golang.org/x/crypto/bcrypt"
import "github.com/AddressBookRevel/app"


type User struct {
	Username string
	Email string
	Password string
}


func (info User) LoginDB() bool{

	rows, err := app.DB.Query("SELECT User_Pass FROM Users WHERE User_Name=? " ,info.Username)
	defer rows.Close()
	checkErr(err)
	var dbPasswordHash []byte
	if rows.Next() {
		rows.Scan(&dbPasswordHash)
	}
	err = bcrypt.CompareHashAndPassword(dbPasswordHash, []byte(info.Password))
	return err == nil
}

func (info User) Exists() bool {

	row, err := app.DB.Query("SELECT * FROM Users WHERE User_Name=?", info.Username )
	defer row.Close()
	checkErr(err)
	return row.Next()
}

func (signinfo User) SignupDB(){

	securedpass,err := bcrypt.GenerateFromPassword([]byte(signinfo.Password),bcrypt.DefaultCost)
	_,err = app.DB.Exec("INSERT INTO Users(User_ID,User_Name,User_Email,User_Pass) VALUES (?, ?, ?, ?)",
		nil, signinfo.Username, signinfo.Email,securedpass)
	checkErr(err)
}

