package db

import (
	"errors"
	"github.com/apaz037/go-metadata-api/models"
	"github.com/google/uuid"
)

// TODO: make thread safe
var Database map[string]models.Application

// type alias for models.Application
type Application = models.Application

// db functions
func Insert(db map[string]Application, app Application) uuid.UUID {
	// TODO: uuid collisions possible here.
	// app is already validated at this point, no need to perform twice
	uuid := uuid.New()
	db[uuid.String()] = app
	return uuid
}

func Get(db map[string]Application, uuid uuid.UUID) error {
	err := db[uuid.String()].Validate() // validate entry before returning to user since we're using a map and it could just return empty struct
	if err != nil {
		return errors.New(uuid.String() + " not found")
	} else {
		return nil
	}
}

func Delete(db map[string]Application, uuid string) Application {
	app := db[uuid]  // store entry we are about to delete so we can return it to the user, questionable, but fits for simple uses
	delete(db, uuid) // remove entry
	return app       // return original entry
}

func Update(db map[string]Application, uuid string, app Application) Application {
	db[uuid] = app
	return app
}
