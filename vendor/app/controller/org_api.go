package controller

import (
	"app/model"
	"app/shared"
	"encoding/json"
	"net/http"
)

type OrgCreateApiForm struct {
	OrgName string
}

type OrgCreateApiResponse struct {
	OrgName string
	Key     string
}

func OrgCreateApi(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)

	var form OrgCreateApiForm
	err := dec.Decode(&form)
	if err != nil {
		jsonError(w, "invalid json", http.StatusBadRequest)
		return
	}
	if len(form.OrgName) == 0 {
		jsonError(w, "OrgName is required", http.StatusBadRequest)
		return
	}
	key := shared.RandomKey()
	_, err = model.OrgCreate(form.OrgName, key)
	if err != nil {
		jsonError(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

	jsonOk(w, OrgCreateApiResponse{
		OrgName: form.OrgName,
		Key:     key,
	})
}
