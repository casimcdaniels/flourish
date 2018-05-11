package flourish

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)


func GetStrainEndpoint(service StrainService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		reqId := vars["id"]
		id, err := strconv.ParseUint(reqId, 10, 64)

		if err != nil {
			return
		}

		strain, err := service.Get(id)

		if err != nil {
			return
		}

		b, err := json.Marshal(strain)

		if err != nil {
			return
		}

		w.Write(b)
	}
}
type createStrainRequest struct {
	Name    string        `json:"name"`
	Race    string        `json:"race"`
	Flavors []string      `json:"flavors"`
	Effects StrainEffects `json:"effects"`
}

func CreateStrainEndpoint(service StrainService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createStrainRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)

		if err != nil {
			return
		}

		strain, err := service.Create(req.Name, req.Race, req.Flavors, req.Effects)

		response, err := json.Marshal(strain)
		if err != nil {
			return
		}

		w.Write(response)
	}
}

func DeleteStrainEndpoint(service StrainService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		reqId := vars["id"]
		id, err := strconv.ParseUint(reqId, 10, 64)

		if err != nil {
			return
		}

		err = service.Remove(id)
		if err != nil {
			return
		}
	}
}

type updateStrainRequest struct {
	Name    *string        `json:"name"`
	Race    *string        `json:"race"`
	Flavors *[]string      `json:"flavors"`
	Effects *StrainEffects `json:"effects"`
}

func UpdateStrainEndpoint(service StrainService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req updateStrainRequest

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)

		vars := mux.Vars(r)
		reqId := vars["id"]
		id, err := strconv.ParseUint(reqId, 10, 64)

		if err != nil {
			return
		}

		err = service.Update(id, req.Name, req.Race, req.Flavors, req.Effects)

		if err != nil {
			return
		}

	}
}

func SearchStrainsEndpoint(service StrainService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
