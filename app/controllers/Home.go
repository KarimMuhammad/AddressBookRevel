package controllers

import (
	"github.com/revel/revel"
	"github.com/AddressBookRevel/app/model"
	"fmt"
	"strconv"
)
//in this controller file it will search for app folder in views folder
type Home struct {
	*revel.Controller
}

//in app folder it will search for login file
func (c Home) Home() revel.Result {

	var contact model.ContactInfo
	userid,_ := strconv.Atoi(c.Session["user_id"])
	h := contact.HomeDB(int64(userid))
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

        id,_ := strconv.Atoi(c.Session["user_id"])
		fmt.Println(id,c.Session["user_id"])
		Contactinfo.Phones = append(Contactinfo.Phones , phones)
		lastinsertid := Contactinfo.AddContactDB(int64(id))
		Contactinfo.AddNumDB(lastinsertid)
		contact := model.ContactInfo{

			Id:int(lastinsertid),
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
	contact.DeleteContactDB(contact_id)
	return c.Redirect("/home")
}

func(c Home) AddNumber() revel.Result{

	contact_id := c.Params.Get("Contact_ID")
	var phone model.Phone
	c.Params.Bind(&phone,"Phone")
	//fmt.Println(contact_id,phone.PhoneNumber)
	c.Validation.Required(phone.PhoneNumber)
	if c.Validation.HasErrors(){

		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/home")
	}else{
		id := phone.AddNumDB(contact_id)
		number := model.Phone{
			Id:int(id),
			PhoneNumber:phone.PhoneNumber,
		}
		return c.RenderJSON(number)
	}

}