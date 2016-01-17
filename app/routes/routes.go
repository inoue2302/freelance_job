// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApiBaseController struct {}
var ApiBaseController tApiBaseController


func (_ tApiBaseController) HandleBadRequestError(
		text string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "text", text)
	return revel.MainRouter.Reverse("ApiBaseController.HandleBadRequestError", args).Url
}

func (_ tApiBaseController) HandleNotFoundError(
		text string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "text", text)
	return revel.MainRouter.Reverse("ApiBaseController.HandleNotFoundError", args).Url
}

func (_ tApiBaseController) HandleInternalServerError(
		text string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "text", text)
	return revel.MainRouter.Reverse("ApiBaseController.HandleInternalServerError", args).Url
}


type tGorpController struct {}
var GorpController tGorpController


func (_ tGorpController) Begin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GorpController.Begin", args).Url
}

func (_ tGorpController) Commit(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GorpController.Commit", args).Url
}

func (_ tGorpController) Rollback(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GorpController.Rollback", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
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
	return revel.MainRouter.Reverse("Static.Serve", args).Url
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
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tJobController struct {}
var JobController tJobController


func (_ tJobController) Detail(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("JobController.Detail", args).Url
}

func (_ tJobController) List(
		page int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "page", page)
	return revel.MainRouter.Reverse("JobController.List", args).Url
}


