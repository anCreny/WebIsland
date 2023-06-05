package WebIsland

import (
	"net/http"
)

type IMiddleware interface {
	Handle(resp http.ResponseWriter, req *http.Request, services *ServicesHandler) bool
}
