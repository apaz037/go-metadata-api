package test

import (
	"github.com/apaz037/go-metadata-api/api"
	"github.com/apaz037/go-metadata-api/models"
	"testing"
)

func TestGetAllApplicationsHandler(t *testing.T) {
	t.Skip()
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

	// create testApp2 to insert into DB
	testApp2 := models.Application{
		Title:   "Valid Application 2",
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

}
