package main

import (
	"github.com/facebookgo/inject"
	"net/http"
)

func main() {
	handler := createRootHandler()
	http.ListenAndServe("0.0.0.0:9000", handler)
}

type RootObj struct{}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateRootObj() (result *RootObj, err error) {
	result = new(RootObj)
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				panic(r)
			}
		}
	}()
	graph := inject.Graph{}
	must(graph.Provide(&inject.Object{Value: result}))
	return
}
