package main

import (
	"./controller"

	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	routes(r)
	err := http.ListenAndServe("localhost:4444", r)
	if err != nil {
		log.Fatal(err)
	}
}

func routes(r *httprouter.Router) {
	r.ServeFiles("/public/*filepath", http.Dir("public"))
	r.GET("/", controller.StartPage)
	r.GET("/users", controller.GetUsers)
}
