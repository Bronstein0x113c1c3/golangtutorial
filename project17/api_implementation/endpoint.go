package api_implementation

import (
	"encoding/json"
	"net/http"
	"project17/"
)
var student_request 
func addHandler(w http.ResponseWriter, r *http.Request){
	
	json.NewDecoder(r.Body).Decode()
}
