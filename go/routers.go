package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
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

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"AuthCheckPost",
		strings.ToUpper("Post"),
		"/auth/check",
		AuthCheckPost,
	},

	Route{
		"AuthLoginPost",
		strings.ToUpper("Post"),
		"/auth/login",
		AuthLoginPost,
	},

	Route{
		"GrantsGet",
		strings.ToUpper("Get"),
		"/grants",
		GrantsGet,
	},

	Route{
		"GrantsGrantIdFiltersPut",
		strings.ToUpper("Put"),
		"/grants/{grant_id}/filters",
		GrantsGrantIdFiltersPut,
	},

	Route{
		"GrantsGrantIdGet",
		strings.ToUpper("Get"),
		"/grants/{grant_id}",
		GrantsGrantIdGet,
	},
}
