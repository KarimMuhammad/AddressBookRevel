package controllers

import (
	"github.com/revel/revel"
	"github.com/AddressBookRevelWithCassandra/app/model"
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
		return c.Redirect("/home")
	}
}


func (c Sign) Register() revel.Result {


	var signup model.User
	c.Params.Bind(&signup,"user")

	c.Validation.Required(signup.Username)
	c.Validation.Required(signup.Email)
	c.Validation.Required(signup.Password)
	fmt.Println(signup)

	if c.Validation.HasErrors(){

		c.Validation.Keep()
		c.FlashParams()

	}else if(signup.Exists())  {
		fmt.Println("exist" , signup.Exists() )
		c.Flash.Error("user name already exist")
		c.Validation.Keep()
		c.FlashParams()

	}else {
		err := signup.SignupDB()
		if(err !=nil){
			c.Flash.Error("DB Error")
			c.Validation.Keep()
			c.FlashParams()
		}
		c.Flash.Success("Successful Sign up")
	}
	return c.RenderTemplate("Sign/Signup.html")
}


