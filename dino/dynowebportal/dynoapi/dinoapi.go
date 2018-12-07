package dynoapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gy-kim/modern-golang-programming/dino/databaselayer"
)

type DinoRESTAPIHandler struct {
	dbhandler databaselayer.DinoDBHandler
}

func newDinoRESTAPIHandler(db databaselayer.DinoDBHandler) *DinoRESTAPIHandler {
	return &DinoRESTAPIHandler{
		dbhandler: db,
	}
}

func (handler *DinoRESTAPIHandler) searchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	criteria, ok := vars["SearchCriteria"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `No search criteria found, you can either search by nickname via /api/dinos/nickname/rex, or
						to search by type via /api/dinos/type/velociraptor`)
		return
	}

	searchkey, ok := vars["search"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `No search criteria found, you can either search by nickname via /api/dinos/nickname/rex, or
						to search by type via /api/dinos/type/velociraptor`)
		return
	}
	var animal databaselayer.Animal
	var animals []databaselayer.Animal
	var err error
	switch strings.ToLower(criteria) {
	case "nickname":
		animal, err = handler.dbhandler.GetDynoByNickname(searchkey)
	case "type":
		animals, err = handler.dbhandler.GetDynosByType(searchkey)
	}

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error occured while querying animals %v", err)
		return
	}
	if len(animals) > 0 {
		json.NewEncoder(w).Encode(animals)
		return
	}
	json.NewEncoder(w).Encode(animal)
}
