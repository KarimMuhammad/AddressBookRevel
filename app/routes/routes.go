// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tSign struct {}
var Sign tSign


func (_ tSign) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sign.Login", args).URL
}

func (_ tSign) Signup(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sign.Signup", args).URL
}

func (_ tSign) Signin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sign.Signin", args).URL
}

func (_ tSign) Register(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sign.Register", args).URL
}


type tHome struct {}
var Home tHome


func (_ tHome) Home(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Home.Home", args).URL
}

func (_ tHome) AddContact(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Home.AddContact", args).URL
}

func (_ tHome) DeleteContact(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Home.DeleteContact", args).URL
}

func (_ tHome) DeleteNumber(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Home.DeleteNumber", args).URL
}

func (_ tHome) AddNumber(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Home.AddNumber", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


