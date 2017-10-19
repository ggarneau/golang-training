package reply

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type NameResponse struct {
	Name string `json:"Hello"`
}

type StatusResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func JSON(w http.ResponseWriter, j interface{}) {
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(j)
	if err != nil {
		fmt.Println(err)
		InternalError(w)
		return
	}
	fmt.Fprintf(w, "%s", jsonResp)
}

func InternalError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "{\"status\":%d,\"error\":\"%s\"}", http.StatusInternalServerError, "Internal Error!")
}
