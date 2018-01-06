package main

import (
	"io/ioutil"
	"net/http/httptest"
	"net/http"
	"testing"
	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/routes"
)

func TestFizzRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(routes.FizzHandler))
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}

	checkBody(t, resp, "buzz")
}

func checkBody(t *testing.T, res *http.Response, body string) {
	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Error("reading reponse body: %v, want %q", err, body)
	}

	if g, w := string(b), body; g != w {
		t.Errorf("request body mismatch: got %q, want %q", g, w)
	}
}
