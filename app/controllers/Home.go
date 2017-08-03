package controllers

import (
	"github.com/revel/revel"
	"github.com/AddressBookRevelWithCassandra/app/model"
	"github.com/AddressBookRevelWithCassandra/app"
	"fmt"
	"github.com/gocql/gocql"
)
//in this controller file it will search for app folder in views folder
type Home struct {
	*revel.Controller
}

//in app folder it will search for login file
func (c Home) Home() revel.Result {

	var contact model.ContactInfo
	user_name:= c.Session["user_name"]
	ID, err := FindUserID(user_name)
	h  := contact.GetContacts(ID)
	if(err !=nil) {
		c.Flash.Error(err.Error())
		c.Validation.Keep()
		c.FlashParams()
	}
	return c.Render(h)
}

func (c Home) AddContact() revel.Result{
	var  Contactinfo model.ContactInfo
	var phones model.Phone

	c.Params.Bind(&Contactinfo,"ContactInfo")
	phones.PhoneNumber = c.Params.Form.Get("phone")

	c.Validation.Required(Contactinfo.FirstName)
	c.Validation.Required(Contactinfo.LastName)
	c.Validation.Required(Contactinfo.Email)
	c.Validation.Required(Contactinfo.Company)
	c.Validation.Required(Contactinfo.JobTitle)
	fmt.Println(Contactinfo.Phones)

	if c.Validation.HasErrors(){

		c.Validation.Keep()
		c.FlashParams()

	}else {

		user_name:= c.Session["user_name"]
		ID, err := FindUserID(user_name)
		Contactinfo.Phones = append(Contactinfo.Phones , phones)
		lastinsertid ,err := Contactinfo.AddContactDB(ID)
		err = Contactinfo.AddNumDB(lastinsertid)
		if(err !=nil) {
			c.Flash.Error("DB Error")
			c.Validation.Keep()
			c.FlashParams()
		}
		contact := model.ContactInfo{

			Id: lastinsertid,
			FirstName: Contactinfo.FirstName,
			LastName: Contactinfo.LastName,
			Email: Contactinfo.Email,
			Company:Contactinfo.Company,
			JobTitle: Contactinfo.JobTitle,
			Phones: Contactinfo.Phones,
		}
		   return c.RenderJSON(contact)
	}
	return c.Redirect("/home")

}

func(c Home) DeleteContact() revel.Result{

	var contact model.ContactInfo
	contact_id := c.Params.Get("Contact_ID")
	user_name:= c.Session["user_name"]
	ID, err := FindUserID(user_name)
	contactuuid,err := gocql.ParseUUID(contact_id)
	err = contact.DeleteContactDB(ID , contactuuid)
	if(err !=nil) {
		c.Flash.Error("DB Error ")
		c.Validation.Keep()
		c.FlashParams()
	}
	return c.Redirect("/home")
}

func(c Home) AddNumber() revel.Result{

	contact_id := c.Params.Get("Contact_ID")
	var phone model.Phone
	c.Params.Bind(&phone,"Phone")
	c.Validation.Required(phone.PhoneNumber)
	if c.Validation.HasErrors(){

		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/home")
	}else{
		contactuuid,err := gocql.ParseUUID(contact_id)
		id ,err := phone.AddNumDB(contactuuid)
		if(err !=nil) {
			c.Flash.Error("DB Error ")
			c.Validation.Keep()
			c.FlashParams()
		}
		number := model.Phone{
			Id: id,
			PhoneNumber:phone.PhoneNumber,
		}
		return c.RenderJSON(number)
	}

}

func FindUserID(name string) (gocql.UUID,error){

	var id gocql.UUID
	err := app.CassandraSession.Query("SELECT user_id FROM user_by_name WHERE user_name=? " ,name).Scan(&id)
	return id,err
}