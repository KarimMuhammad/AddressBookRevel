package controllers

import (
	"github.com/revel/revel"
	"strconv"
	"github.com/AddressBookRevel/app/model"
	"github.com/AddressBookRevel/app"

	"fmt"
)
//in this controller file it will search for Sign folder in views folder

type Sign struct {
	*revel.Controller
}
//in Sign folder it will search for login file

func (c Sign) Login() revel.Result {
	return c.Render()
}

func (c Sign) Signup() revel.Result {
	return c.Render()
}

func (c Sign) Signin() revel.Result {

	var login model.User
	c.Params.Bind(&login, "user")

	c.Validation.Required(login.Username)
	c.Validation.Required(login.Password)

	if c.Validation.HasErrors() && ! login.LoginDB() {
		c.Flash.Error("unsuccessful login")
		c.Validation.Keep()
		c.FlashParams()
		return c.RenderTemplate("Sign/Login.html")

	} else {
		c.Session["user_name"] = login.Username
		c.Session["user_id"] = strconv.Itoa(FindUserID(login.Username))
		fmt.Println(c.Session["user_id"])
		return c.Redirect("/home")
	}
}


func (c Sign) Register() revel.Result {

	var signup model.User
	c.Params.Bind(&signup,"user")

	c.Validation.Required(signup.Username)
	c.Validation.Required(signup.Email)
	c.Validation.Required(signup.Password)

	if c.Validation.HasErrors(){

		c.Validation.Keep()
		c.FlashParams()

	}else if(signup.Exists())  {

		c.Flash.Error("user name already exist")
		c.Validation.Keep()
		c.FlashParams()

	}else {
		signup.SignupDB()
		c.Flash.Success("Successful Sign up")
	}
	return c.RenderTemplate("Sign/Signup.html")
}

func FindUserID(name string) int{

	rows, _:= app.DB.Query("SELECT User_ID FROM Users WHERE User_Name=? " ,name)
	defer rows.Close()
	var id int
	if rows.Next() {
		rows.Scan(&id)
	}
	return id
}
