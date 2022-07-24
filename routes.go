package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"MovieList",
		"GET",
		"/peliculas",
		MovieList,
	},
	Route{
		"MovieShow",
		"GET",
		"/peliculas/{id}",
		MovieShow,
	},
	Route{
		"MovieAdd",
		"POST",
		"/pelicula",
		MovieAdd,
	},
	Route{
		"MovieUpdate",
		"POST",
		"/pelicula/{id}",
		MovieUpdate,
	},
	Route{
		"MovieRemove",
		"POST",
		"/remove/{id}",
		MovieRemove,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandleFunc)
	}
	return router
}
