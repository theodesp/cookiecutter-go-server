package routes

import "net/http"

// Convenient struct to save routes
type Route struct {
	Path    string
	Handler func(w http.ResponseWriter, req *http.Request)
}

var Routes = []Route {
	{
		Path: "/fizz",
		Handler: FizzHandler,
	},
}

func FizzHandler(w http.ResponseWriter, _ *http.Request)  {
	w.Write([]byte("buzz"))
}
