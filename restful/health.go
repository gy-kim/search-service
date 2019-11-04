package restful

import (
	"encoding/json"
	"net/http"
)

func health(response http.ResponseWriter, r *http.Request) {
	res := struct {
		State string `json:"state"`
	}{
		State: "OK",
	}

	byRes, _ := json.Marshal(res)

	response.Write(byRes)
}
