package webApp

import (
	"WebIsland/webApp/Controllers"
	"WebIsland/webApp/ServicesSystem"
	"log"
	"net/http"
	"strings"
)

type App struct {
	address   string
	executors []string

	middlewares        map[string]*IMiddleware
	controllersHandler *Controllers.ControllerHandler

	Services *ServicesSystem.Handler

	servicesUpdateChan chan int
}

func PrepareBuilding(ipAddress string, port string) (app *App) {
	var servicesUpdateChan = make(chan int)

	app = &App{ipAddress + ":" + port, []string{}, map[string]*IMiddleware{}, nil, ServicesSystem.NewHandler(servicesUpdateChan), servicesUpdateChan}
	return
}

func (this *App) AddMiddleware(middleware IMiddleware) {
	var identifier = string(rune(len(this.middlewares) + 1))

	this.executors = append(this.executors, identifier)
	this.middlewares[identifier] = &middleware
}

func (this *App) UseControllers() *Controllers.ControllerHandler {
	this.controllersHandler = Controllers.NewControllerHandler()
	this.executors = append(this.executors, "controllersHandler")
	return this.controllersHandler
}

func (this *App) Run() {
	log.Fatal(http.ListenAndServe(this.address, this))
}

func (this *App) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	var accpets = req.Header.Get("Accept")
	var formattedAccepts = strings.Split(accpets, ",")
	if formattedAccepts[0] != "image/avif" {

		this.Services.StartRequest()

	Loop:
		for _, val := range this.executors {
			switch val {
			case "controllersHandler":
				this.controllersHandler.HandleThread(resp, req, this.Services)
			default:
				if middleware, ok := this.middlewares[val]; ok {
					if !(*middleware).Handle(resp, req, this.Services) {
						break Loop
					}
				}
			}
		}
	}
}
