package main

import (
	"log"
	"net/http"
)

type Route struct {
	//path of the http route
	//must follow path patter of url
	path string
	//handler function to handle http request on the specific path
	handler func(res http.ResponseWriter, req *http.Request)
}

func main() {

	//Array of Route
	var routes = []Route{
		{
			path: "/",
			handler: func(res http.ResponseWriter, req *http.Request) {
				res.WriteHeader(200)
				res.Header().Add("content-type", "application/json")
				res.Write([]byte("{\"hello\":\"world\"}"))
			},
		},
	}

	//Adding Route as Handle Functions
	for i := 0; i < len(routes); i++ {
		http.HandleFunc(
			routes[i].path,
			routes[i].handler,
		)
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
