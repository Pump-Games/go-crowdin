package crowdin

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

var (
	mux     *http.ServeMux
	crowdin *Crowdin
	server  *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	crowdin = New("token", "project-name")

	// Fake the API and Stream base URLs by using the test
	// server URL instead.
	url, _ := url.Parse(server.URL)
	crowdin.config.apiBaseURL = url.String() + "/"
	crowdin.config.apiAccountBaseURL = url.String() + "/"
}

func teardown() {
	server.Close()
}
