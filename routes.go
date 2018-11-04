// Package main is the main package
// -------------------------------------
// .../restauranteapi/routes.go
// -------------------------------------
package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route is
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes are
type Routes []Route

// VNewRouter is
func VNewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

// XNewRouter is
func XNewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{

	Route{"savetomysql", "GET", "/savetomysql", CopyOrdersToMySQL},
	Route{"orderlist", "GET", "/orderlist", OrderListV2},
	Route{"ordercompleted", "GET", "/ordercompleted", ordercompleted},
	Route{"orderstatus", "GET", "/orderstatus", orderstatus},
	Route{"orderadd", "POST", "/orderadd", Horderadd},
	Route{"APIorderadd", "POST", "/APIorderadd", HAPIorderadd},
	Route{"orderadd", "POST", "/orderupdate", Horderupdate},
	Route{"orderfind", "GET", "/orderfind", Horderfind},
	// -------------------------------------------------------------------
	Route{"getcachedvalues", "GET", "/getcachedvalues", getcachedvalues},
	Route{"savetomysql", "GET", "/savetomysql", CopyOrdersToMySQL},
}
