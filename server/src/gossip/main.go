package main

import (
	"fmt"
	"gossip/application"
	"gossip/dataaccess"
	"gossip/queueing"
	"net/http"

	"github.com/facebookgo/inject"
)

func main() {
	// handler := createRootHandler()
	fmt.Println("Starting")
	handler, err := CreateRootObj()
	must(err)
	handler.Init()
	fmt.Println("Starting")
	http.ListenAndServe("0.0.0.0:9000", logHandler{handler})
}

type logHandler struct {
	handler http.Handler
}

func (l logHandler) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	fmt.Printf("Request: %s %s\n", req.Method, req.URL)
	l.handler.ServeHTTP(wr, req)
}

type RootObj struct {
	App          *application.Application `inject:""`
	MessageHub   *application.MessageHub  `inject:""`
	*HttpHandler `inject:""`
}

func (r RootObj) Init() {
	r.HttpHandler.Init()
	r.MessageHub.Init()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func catch(perr *error) {
	if r := recover(); r != nil {
		var ok bool
		*perr, ok = r.(error)
		if !ok {
			panic(r)
		}
	}
}

var dbUrl = "postgres://gossip:gossip@127.0.0.1/gossip?sslmode=disable"
var amqpUrl = "amqp://localhost/" // "amqp://guest:guest@localhost"

func CreateRootObj() (result *RootObj, err error) {
	result = new(RootObj)
	defer catch(&err)
	graph := inject.Graph{}
	dbConn, err := dataaccess.NewConnection(dbUrl)
	must(err)
	amqpConn, err := queueing.NewConnection(amqpUrl)
	must(err)
	must(graph.Provide(&inject.Object{Value: result}))
	must(graph.Provide(&inject.Object{Value: &dbConn}))
	must(graph.Provide(&inject.Object{Value: &amqpConn}))
	err = graph.Populate()
	return
}
