package WebIsland

import (
	"net/http"
)

type IController interface {
	GetRoute() string
	GetActions() map[string]*func(resp http.ResponseWriter, req *http.Request, services *ServicesHandler)
}
