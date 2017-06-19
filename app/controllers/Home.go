package controllers

import (
	"github.com/revel/revel"
	"github.com/AddressBookRevel/app/model"
	"github.com/AddressBookRevel/app"
	"fmt"
	"strconv"

)
//i this controller file it will search for app folder in views folder
type Home struct {
	*revel.Controller
}

//in app folder it will search for login file
func (c Home) Home() revel.Result {
	h := []model.ShowContacts{}
	userid,_ := strconv.Atoi(c.Session["user_id"])
	contactrows,_ :=app.DB.Query("SELECT Contact_ID, First_Name, Last_Name, Job_Title, Company, Email FROM Contact_List WHERE User_Id_FK=? ",userid)


	for contactrows.Next(){
		var contact model.ShowContacts
		contactrows.Scan(&contact.Contact_ID, &contact.First_Name, &contact.Last_Name, &contact.Job_Title, &contact.Company,&contact.Email)
		numbersrows,_ :=app.DB.Query("SELECT No_ID, Phone_Number FROM Phone_Numbers WHERE Contact_Id_FK=? ",&contact.Contact_ID)
		for numbersrows.Next(){
			var no model.PhoneNo
			numbersrows.Scan(&no.No_ID,&no.Phone_Number)
			contact.Phones = append(contact.Phones,no)
		}
		h = append(h,contact)


	}
	return c.Render(h)
}

func (c Home) AddContact() revel.Result{
	var  Contactinfo model.Contactinfo

	c.Params.Bind(&Contactinfo,"Contactinfo")

	c.Validation.Required(Contactinfo.First_Name)
	c.Validation.Required(Contactinfo.Last_Name)
	c.Validation.Required(Contactinfo.Email)
	c.Validation.Required(Contactinfo.Company)
	c.Validation.Required(Contactinfo.Job_Title)
	fmt.Println(Contactinfo)


	if c.Validation.HasErrors(){

		c.Validation.Keep()
		c.FlashParams()

	}else {

        id,_ := strconv.Atoi(c.Session["user_id"])
		fmt.Println(id,c.Session["user_id"])
		query,_:=app.DB.Exec("INSERT INTO Contact_List( Contact_ID, First_Name, Last_Name, Job_Title, Company, Email, User_Id_FK) VALUES (?,?,?,?,?,?,?)", nil, Contactinfo.First_Name, Contactinfo.Last_Name, Contactinfo.Job_Title, Contactinfo.Company, Contactinfo.Email,id)
		pk,_ := query.LastInsertId()
		contact := model.ShowContacts{

			Contact_ID:int(pk),
			First_Name: Contactinfo.First_Name,
			Last_Name: Contactinfo.Last_Name,
			Email: Contactinfo.Email,
			Company:Contactinfo.Company,
			Job_Title: Contactinfo.Job_Title,

		}
		   return c.RenderJSON(contact)

	}
	return c.Redirect("/home")

}

func(c Home) DeleteContact() revel.Result{
	contact_id := c.Params.Get("Contact_ID")
	app.DB.Exec("DELETE FROM Phone_Numbers WHERE Contact_Id_FK=?",contact_id)
	app.DB.Exec("DELETE FROM Contact_List WHERE Contact_ID=?",contact_id)
	return c.Redirect("/home")
}

func(c Home) AddNumber() revel.Result{

	contact_id := c.Params.Get("Contact_ID")
	//number := c.Params.Get("no")
	var phone model.PhoneNo
	c.Params.Bind(&phone,"PhoneNo")
	fmt.Println(contact_id,phone.Phone_Number)
	query,_:=app.DB.Exec("INSERT INTO Phone_Numbers( No_ID,Phone_Number,Contact_Id_FK) VALUES (?,?,?)",nil,phone.Phone_Number,contact_id)
	pk,_ := query.LastInsertId()

	number := model.PhoneNo{

		No_ID:int(pk),
		Phone_Number:phone.Phone_Number,
	}
	return c.RenderJSON(number)

}