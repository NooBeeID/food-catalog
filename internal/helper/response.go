package helper

import (
	"encoding/json"
	"net/http"
)

// format response
type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

// function ini berfungsi untuk mengirim response ke client
func (resp APIResponse) WriteJsonResponse(w http.ResponseWriter) {
	// untuk menambahkan response header
	// dengan nama content-type dengan isi nya adalah application/json
	// selalu set content type duluan sebelum menulis status code
	w.Header().Set("Content-Type", "application/json")

	// set CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Agent")

	// untuk membuat http status code
	w.WriteHeader(resp.Status)

	// proses kirim response
	json.NewEncoder(w).Encode(resp)
}
