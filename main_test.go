package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goseed/models/service"
	"goseed/routers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func initServer() *httptest.Server {
	router := routers.InitRoute()
	ts := httptest.NewServer(router)
	return ts
}

func createUser(URL string) (*http.Response, error) {

	// makes a POST request to signup
	postBody := map[string]interface{}{
		"email":    "abc@xyz.com",
		"name":     "test",
		"password": "test123",
	}
	body, _ := json.Marshal(postBody)
	resp, err := http.Post(fmt.Sprintf("%s/signup", URL), "Application/JSON", bytes.NewReader(body))
	return resp, err
}

func TestSignupRoute(t *testing.T) {
	ts := initServer()

	// Shut down the server and block until all requests have gone through
	defer ts.Close()
	// The setupServer method, that we previously refactored
	// is injected into a test server
	resp, err := createUser(ts.URL)

	// Make a request to our server with the {base url}/ping
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	// Assert that the "content-type" header is actually set
	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	// Assert that it was set as expected
	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
	userservice := service.Userservice{}
	defer userservice.Delete("abc@xyz.com")
}

func TestLoginRoute(t *testing.T) {
	// The setupServer method, that we previously refactored
	// is injected into a test server

	ts := initServer()
	// Shut down the server and block until all requests have gone through
	defer ts.Close()

	// Let's create a user first
	userCreationResp, _ := createUser(ts.URL)
	userservice := service.Userservice{}
	defer userservice.Delete("abc@xyz.com")
	if userCreationResp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", userCreationResp.StatusCode)
	}
	// makes a POST request to signup
	postBody := map[string]interface{}{
		"email":    "abc@xyz.com",
		"password": "test123",
	}
	body, _ := json.Marshal(postBody)
	resp, err := http.Post(fmt.Sprintf("%s/login", ts.URL), "Application/JSON", bytes.NewReader(body))

	// Make a request to our server with the {base url}/ping
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
}
