package rootrequesthandler

import (
	"fmt"
	"net/http"
)

type RootRequestHandler struct{}

func (rqh *RootRequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rqh.handleGET(w)
	case http.MethodPost:
		rqh.handlePOST(w)
	}
}

func (rqh *RootRequestHandler) handleGET(w http.ResponseWriter) {
	fmt.Fprintf(w, "Hello, world from get method")
}

func (rqh *RootRequestHandler) handlePOST(w http.ResponseWriter) {
	fmt.Fprintf(w, "Hello, world from POST")
}
