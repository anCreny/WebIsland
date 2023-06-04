package webApp

import (
	"WebIsland/webApp/ServicesSystem"
	"net/http"
)

type IMiddleware interface {
	Handle(resp http.ResponseWriter, req *http.Request, services *ServicesSystem.Handler) bool
}
