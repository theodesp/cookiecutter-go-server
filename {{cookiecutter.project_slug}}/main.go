package main

import (
	"net/http"
	"log"
	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/routes/routes"
	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/utils/utils"
)

func main()  {
	mux := http.NewServeMux()

	utils.AttachRoutes(mux, routes.Routes)

	log.Fatal(http.ListenAndServe(":3000", mux))
}

