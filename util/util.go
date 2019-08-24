package utils

import (
	"encoding/json"
	"net/http"
)

// Message function accepts a bool and a string. Return a hashmap of type [string]interface{}
func Message(status bool, message string) (map[string]interface{}) {
	return map[string]interface{} {"status" : status, "message" : message}
}

// Respond function takes in a response write with HTTP data
func Respond(respWr http.ResponseWriter, data map[string] interface{})  {
	respWr.Header().Add("Content-Type", "application/json")
	json.NewEncoder(respWr).Encode(data)
}