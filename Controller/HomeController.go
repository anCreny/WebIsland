package Controller

import (
	"fmt"
	"github.com/anCreny/WebIsland"
	"html/template"
	"net/http"
)

type HomeController struct {
}

func (h *HomeController) GetRoute() string { //Implementing IController interface
	return "home" //value to be associated with in request path (ex: localhost/home)
}

func (h *HomeController) GetActions() map[string]*func(resp http.ResponseWriter, req *http.Request, services *WebIsland.ServicesHandler) { //Implementing IController interface
	var index = h.Index

	return map[string]*func(resp http.ResponseWriter, req *http.Request, services *WebIsland.ServicesHandler){ //Creating a map of controller's actions
		"index": &index, //Key value is a value to be associated with in request path after using controller value (ex: localhost/home/index)
	}
}

//P.S.: If u use UseDefaultPathPattern() in controller settings,
//your default request path route, without any path (ex: localhost), will be with home controller and index action (ex: localhost == localhost/home/index)
//remember to realize it or change PathPattern to avoid errors

func NewHomeController() *HomeController { //Constructor (optional)
	return &HomeController{}
}

func (h *HomeController) Index(resp http.ResponseWriter, req *http.Request, services *WebIsland.ServicesHandler) { //Realize of index action
	var message = req.URL.Query().Get("message")
	if message == "" {
		message = "Hello world!"
	}

	data := ViewData{Message: message}
	tmpl, err := template.ParseFiles("Views/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	tmplErr := tmpl.Execute(resp, data)
	if tmplErr != nil {
		fmt.Println(err)
		return
	}
}

type ViewData struct {
	Message string
}
