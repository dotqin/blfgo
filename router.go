package blfgo

import (
	"fmt"
	"net/http"
)

type BlfHandler struct {
}

type BlfRouter interface {
	Route(url string) string
}

var Router BlfRouter

func (h *BlfHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	content := Router.Route(req.RequestURI)
	fmt.Fprint(w, content)
	return
}
