package main

import (
	"github.com/anCreny/WebIsland"
	"webApp/Controller"
)

func main() {
	var app = WebIsland.PrepareBuilding("localhost", "8000") //Getting object of web application for settings and starting

	var controllerHandler = app.UseControllers() //Including controllers in request flow

	controllerHandler.AddController(Controller.NewHomeController()) //Adding new controller
	controllerHandler.UseDefaultPathPattern()                       //Using default controller and action if request path will be empty

	app.Run() //Start listening
}
