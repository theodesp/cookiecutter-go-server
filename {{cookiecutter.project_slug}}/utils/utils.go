package utils

import (
	"net/http"
	_ "github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/routes"
)


func AttachRoutes(mux *http.ServeMux, rts []routes.Route)  {
	for _, r := range rts {
		mux.HandleFunc(r.Path, r.Handler)
	}
}