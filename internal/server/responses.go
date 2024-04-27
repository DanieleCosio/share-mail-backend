package server

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct{}

func (r JsonResponse) Dispatch(w *http.ResponseWriter, data interface{}) {
	(*w).Header().Set("Content-Type", "application/json")
	json.NewEncoder(*w).Encode(data)
}

type PingResponse struct {
	JsonResponse
	Success bool `json:"success"`
}
