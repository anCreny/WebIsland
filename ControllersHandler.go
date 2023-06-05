package WebIsland

import (
	"net/http"
	"strings"
)

type ControllerHandler struct {
	controllers map[string]*IController

	defaultControllerRoute string
	defaultActionRoute     string
}

func NewControllerHandler() *ControllerHandler {
	return &ControllerHandler{
		controllers:            map[string]*IController{},
		defaultControllerRoute: "",
		defaultActionRoute:     "",
	}
}

func (this *ControllerHandler) HandleThread(resp http.ResponseWriter, req *http.Request, services *ServicesHandler) bool {
	var path = req.URL.Path
	pathParts := strings.Split(path, "/")
	pathParts = wipeSpaces(pathParts)

	if len(pathParts) == 0 {
		if controller, ok := this.controllers[this.defaultControllerRoute]; ok {
			actions := (*controller).GetActions()

			if action, ok := actions[this.defaultActionRoute]; ok {
				(*action)(resp, req, services)
			}
		}
	} else if len(pathParts) == 1 {
		if controller, ok := this.controllers[pathParts[0]]; ok {
			actions := (*controller).GetActions()

			if action, ok := actions[this.defaultActionRoute]; ok {
				(*action)(resp, req, services)
			}
		}
	} else {
		if controller, ok := this.controllers[pathParts[0]]; ok {
			actions := (*controller).GetActions()

			if action, ok := actions[pathParts[1]]; ok {
				(*action)(resp, req, services)
			}
		}
	}

	return true
}

func (this *ControllerHandler) AddController(controller IController) {
	route := controller.GetRoute()
	this.controllers[route] = &controller
}

func (this *ControllerHandler) UseDefaultPathPattern() {
	this.defaultControllerRoute = "home"
	this.defaultActionRoute = "index"
}

func (this *ControllerHandler) SetOwnPathPattern(defaultController string, defaultAction string) {
	this.defaultControllerRoute = defaultController
	this.defaultActionRoute = defaultAction
}

func wipeSpaces(array []string) []string {
	var result []string
	for _, value := range array {
		if value != " " && value != "" {
			result = append(result, value)
		}
	}

	return result
}
