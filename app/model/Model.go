package model

type Signup struct {
	Username string
	Email string
	Password string
}

type Signin struct {
	Username string
	Password string
}

type Contactinfo struct {
	First_Name string
	Last_Name string
	Job_Title string
	Company string
	Email string

}
type ShowContacts struct {
	Contact_ID int
	First_Name string
	Last_Name string
	Job_Title string
	Company string
	Email string
	Phones []PhoneNo
}
type PhoneNo struct {
	No_ID int
	Phone_Number string
}


