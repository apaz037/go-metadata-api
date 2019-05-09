package test

import (
	"bytes"
	"fmt"
	"github.com/apaz037/go-metadata-api/api"
	"github.com/apaz037/go-metadata-api/api/handlers"
	"github.com/apaz037/go-metadata-api/db"
	"github.com/apaz037/go-metadata-api/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteApplicationHandler(t *testing.T) {
	t.Skip()
	// set up fresh server
	api.NewServer()

	// create testApp to insert into DB
	testApp := models.Application{
		Title:   "Valid Application 1",
		Version: "0.0.1",
		Maintainers: []models.Maintainer{{
			Name:  "aaron paz",
			Email: "aaron_paz@gmail.com",
		},
			{
				Name:  "another maintainer",
				Email: "another_maintainer@gmail.com",
			},
		},
		Company:     "Random Inc.",
		Website:     "https://testwebsite.com/",
		Source:      "https://github.com/random/repo",
		License:     "Apaache-2.0",
		Description: "a very long string with many special characters should go here.",
	}

	// Insert testApp and retrieve UUID
	appUUID := db.Insert(db.Database, testApp)

	route := "/application/" + appUUID.String()

	req, err := http.NewRequest("DELETE", route, bytes.NewReader([]byte("")))
	fmt.Println(req.RequestURI)

	if err != nil {
		t.Errorf("Could not create request")
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.DeleteApplicationHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned unexpected status code: got %v want %v", status, http.StatusOK)
	}

	err = db.Get(db.Database, appUUID)
	if err == nil { // check if err is nil because we should have already deleted this application entry
		t.Errorf("handler did not delete application from DB")
	}
}

// TODO: test for 404
