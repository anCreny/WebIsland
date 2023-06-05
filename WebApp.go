package WebIsland

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type App struct {
	address   string
	executors []string

	middlewares        map[string]*IMiddleware
	controllersHandler *ControllerHandler

	Services *ServicesHandler

	servicesUpdateChan chan int
}

func PrepareBuilding(ipAddress string, port string) (app *App) {
	var servicesUpdateChan = make(chan int)

	app = &App{ipAddress + ":" + port, []string{}, map[string]*IMiddleware{}, nil, NewHandler(servicesUpdateChan), servicesUpdateChan}
	return
}

func (this *App) AddMiddleware(middleware IMiddleware) {
	var identifier = string(rune(len(this.middlewares) + 1))

	this.executors = append(this.executors, identifier)
	this.middlewares[identifier] = &middleware
}

func (this *App) UseControllers() *ControllerHandler {
	this.controllersHandler = NewControllerHandler()
	this.executors = append(this.executors, "controllersHandler")
	return this.controllersHandler
}

func (this *App) Run() {
	fmt.Println("[" + Green + "success" + Reset + "]: Server is listening now on " + Yellow + this.address + Reset)
	log.Fatal(http.ListenAndServe(this.address, this))
}

func (this *App) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	var accpets = req.Header.Get("Accept")
	var formattedAccepts = strings.Split(accpets, ",")
	if formattedAccepts[0] != "image/avif" {

		this.Services.startRequest()

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
