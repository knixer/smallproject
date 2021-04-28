package servertest

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
)

// Customizer is a type for request-modifying functions that can be used to
// customize requests created by the shorthand request functions (Do, Get, Post, Delete)
type Customizer func(*http.Request)

// Do performs given method on url with handler and returns an
// httptest.ResponseRecorder to inspect the resulting response
func Do(handler http.Handler, method, url string, v interface{}, fs ...Customizer) *httptest.ResponseRecorder {
	var body bytes.Buffer
	if v != nil {
		if err := json.NewEncoder(&body).Encode(&v); err != nil {
			log.Fatalf("encode json: %v", err)
		}
	}
	req, err := http.NewRequest(method, url, &body)
	if err != nil {
		log.Fatalf("create request: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	for _, f := range fs {
		f(req)
	}

	res := httptest.NewRecorder()

	handler.ServeHTTP(res, req)

	return res
}

// Get is a shorthand for Do(handler, "GET", url, nil)
func Get(handler http.Handler, url string, fs ...Customizer) *httptest.ResponseRecorder {
	return Do(handler, "GET", url, nil, fs...)
}
