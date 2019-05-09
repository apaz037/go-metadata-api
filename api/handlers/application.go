package handlers

import (
	"encoding/json"
	"github.com/apaz037/go-metadata-api/api/utils"
	"github.com/apaz037/go-metadata-api/db"
	"github.com/apaz037/go-metadata-api/models"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func CreateApplicationHandler(w http.ResponseWriter, r *http.Request) {
	application, err := models.Unmarshal(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	appUUID := db.Insert(db.Database, application) // insert application into DB, get entry UUID

	retrievedApp := db.Database[appUUID.String()]

	uuidString := appUUID.String()

	resp := models.ApplicationWithID{
		ID:          uuidString,
		Application: retrievedApp,
	}

	json, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(utils.JsonPrettyPrint(json)))
	return
}

func GetApplicationHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")    // get UUID of app to retrieve from URL
	appUUID, err := uuid.Parse(id) // ensure valid UUID
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not parse id: " + id + " as uuid"))
		return
	}

	err = db.Get(db.Database, appUUID) // retrieve entry from DB based on UUID
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Could not find application with ID: " + appUUID.String()))
		return
	}

	app := db.Database[appUUID.String()] // grab app, doesnt need to be here but makes code more readable for a limited memory allocation increase

	json, err := json.Marshal(app)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not marshal json response"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(utils.JsonPrettyPrint(json))
	return
}

func GetAllApplicationsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: paginate
	// might not have time for this, would be nice to have, add later
	json, err := json.Marshal(db.Database)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error retrieving applications"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(utils.JsonPrettyPrint(json)))
	return
}

func DeleteApplicationHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id") // get ID to delete from URL

	if _, ok := db.Database[id]; !ok { // check if entry exists in DB
		resp := id + ": not found"
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(resp))
		return
	}

	app := db.Delete(db.Database, id) // delete, store app for response

	json, err := json.Marshal(app)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(utils.JsonPrettyPrint(json))
	return
}

func UpdateApplicationHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id") // get ID to update from URL
	if _, ok := db.Database[id]; !ok {
		resp := id + ": not found"
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(resp))
		return
	}

	application, err := models.Unmarshal(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	updatedApp := db.Update(db.Database, id, application)

	resp := models.ApplicationWithID{
		ID:          id,
		Application: updatedApp,
	}

	json, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(utils.JsonPrettyPrint(json)))
	return
}

func SearchApplicationHandler(w http.ResponseWriter, r *http.Request) {
	// grab the full query string
	rawQueryString := r.URL.RawQuery // debug
	log.Println(rawQueryString)      // debug

	queryStringMap := make(map[string]string) // create map to hold all possible query params

	// TODO: lowercase all values in query string
	queryStringMap["Title"] = r.URL.Query().Get("title")
	queryStringMap["Version"] = r.URL.Query().Get("version")
	//queryStringMap["Maintainers"] = r.URL.Query().Get("maintainers") // this one will require some extra work
	queryStringMap["Company"] = r.URL.Query().Get("company")
	queryStringMap["Website"] = r.URL.Query().Get("website")
	queryStringMap["Source"] = r.URL.Query().Get("source")
	queryStringMap["License"] = r.URL.Query().Get("license")
	queryStringMap["Description"] = r.URL.Query().Get("description")

	searchResults := make(map[string]models.Application) // holds our search results

	// search
	for uuid, application := range db.Database { // loop over database
		reflectedApp := reflect.ValueOf(application)                    // grab termValue of application struct
		reflectedValues := make([]interface{}, reflectedApp.NumField()) // grab num values in struct

		// populate reflectedValues
		for i := 0; i < reflectedApp.NumField(); i++ {
			reflectedValues[i] = reflectedApp.Field(i).String()
		}

		for _, termValue := range queryStringMap {
			if termValue != "" {
				for i := 0; i < len(reflectedValues); i++ {
					t, ok := reflectedValues[i].(string)
					if ok {
						if strings.Contains(t, termValue) { // check if query string termValue and reflectedApp's termValue are the same for key
							searchResults[uuid] = application // add to search results
						}
					}
				}
			}
		}
	}

	json, err := json.Marshal(searchResults)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(utils.JsonPrettyPrint(json))
	return
}
