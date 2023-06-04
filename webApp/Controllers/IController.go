package Controllers

import (
	"WebIsland/webApp/ServicesSystem"
	"net/http"
)

type IController interface {
	GetRoute() string
	GetActions() map[string]*func(resp http.ResponseWriter, req *http.Request, services *ServicesSystem.Handler)
}
