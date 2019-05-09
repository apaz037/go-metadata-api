package test

import (
	"bytes"
	"github.com/apaz037/go-metadata-api/api"
	"github.com/apaz037/go-metadata-api/api/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateApplicationHandler(t *testing.T) {
	api.NewServer()

	testAppStr := `title: Valid Application 1
version: 0.0.1
maintainers:
- name: firstmaintainer app1
  email: firstmaintainer@hotmail.com
- name: secondmaintainer app1
  email: secondmaintainer@gmail.com
company: Random Inc.
website: https://website.com
source: https://github.com/random/repo
license: Apache-2.0
description: |
 ### Interesting Title
 Some application content, and description`

	req, err := http.NewRequest("POST", "/application", bytes.NewReader([]byte(testAppStr)))
	if err != nil {
		t.Errorf("Could not create request")
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateApplicationHandler)
	handler.ServeHTTP(rr, req)

	// check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned unexpected status code: got %v want %v", status, http.StatusOK)
	}

	// TODO: check body
}
