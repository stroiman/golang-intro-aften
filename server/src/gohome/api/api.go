package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var Router http.Handler

func getPosts(wr http.ResponseWriter, req *http.Request) {
	obj := struct{ Id string }{Id: "42"}
	wr.Header().Add("Content-Type", "application/json")
	wr.WriteHeader(http.StatusOK)
	json.NewEncoder(wr).Encode(obj)
}

func init() {
	router := mux.NewRouter()
	router.Path("/blogs").HandlerFunc(getPosts)
	Router = router
}
