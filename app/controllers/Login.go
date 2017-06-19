package controllers

import (
	"github.com/revel/revel"
	"github.com/AddressBookRevel/app/model"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	"github.com/AddressBookRevel/app"

	"strconv"
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

	var login model.Signin
	var user_id int
	c.Params.Bind(&login, "signin")

	c.Validation.Required(login.Username)
	c.Validation.Required(login.Password)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
	} else {
		rows, _ := app.DB.Query("SELECT User_ID, User_Name, User_Pass FROM Users WHERE User_Name=? ", login.Username)
		for rows.Next() {
			var user model.Signin
			rows.Scan(&user_id, &user.Username, &user.Password)
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
			if err != nil {
				c.Flash.Data["log"] = "Wrong Password"
			} else {

				c.Session["user_name"] = user.Username
				c.Session["user_id"] = strconv.Itoa(user_id)
				return c.Redirect("/home")
				//return c.Redirect("/home/%d", user_id)

			}
		}
	}
	return c.RenderTemplate("Sign/Login.html")
}
func (c Sign) Register() revel.Result {

	var signup model.Signup
	c.Params.Bind(&signup,"signup")

	c.Validation.Required(signup.Username)
	c.Validation.Required(signup.Email)
	c.Validation.Required(signup.Password)
	fmt.Println(signup)

	if c.Validation.HasErrors(){

		c.Validation.Keep()
		c.FlashParams()

	}else {
		Upassword ,_ := bcrypt.GenerateFromPassword([]byte(signup.Password),bcrypt.DefaultCost)
		data,err :=app.DB.Prepare("INSERT INTO Users(User_ID,User_Name,User_Email,User_Pass) VALUES (?, ?, ?, ?)")
		res,err := data.Exec(nil, signup.Username, signup.Email,Upassword)
		fmt.Println(err,res)
		c.Flash.Success("Successful Sign up")
	}
	return c.RenderTemplate("Sign/Signup.html")
}
