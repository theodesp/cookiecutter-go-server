package main

import (
	"net/http"
	"log"
	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/routes"
	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/utils"
	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/config"
	"time"
)

func main()  {
	mux := http.NewServeMux()

	utils.AttachRoutes(mux, routes.Routes)
	config := config.LoadConfig()

	server := &http.Server{
		Addr:           config.Host + ":" + config.Port,
		Handler:        mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}

