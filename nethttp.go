package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Layer0: Model

// ModelVersion ...
type ModelVersion struct {
	Version string `json:"version"`
}

// Layer1: Endpoint --------------------------------------------------------------------------------

func main() {

	l := HandlerVersion{
		LogicVersion: NewLogic(), // Call NewLogic factory (or constructor)
	}

	http.HandleFunc("/api/v1/version", l.GetVersion)
	log.Fatal(http.ListenAndServe(":9191", nil))
}

// Layer2: Handler --------------------------------------------------------------------------------

// HandlerVersion ...
type HandlerVersion struct {
	LogicVersion LogicVersion
}

// GetVersion is doing ...
func (r HandlerVersion) GetVersion(w http.ResponseWriter, req *http.Request) {

	request := "some random request from http"
	resultFromLogic := r.LogicVersion.GetVersion(request)
	result, err := json.Marshal(resultFromLogic)
	if err != nil {
		io.WriteString(w, "error-ed7869c4: some error")
		return
	}

	io.WriteString(w, string(result))
}

// Layer3: Logic--------------------------------------------------------------------------------

// LogicVersion ...
type LogicVersion struct {
	DataBaseURL string
}

// NewLogic constructor
func NewLogic() LogicVersion {
	return LogicVersion{}
}

// GetVersion ...
func (r LogicVersion) GetVersion(request string) ModelVersion {
	result := ModelVersion{
		Version: "ver1.0.0",
	}

	return result
}

// Layer4: Storage/Queue--------------------------------------------------------------------------------

// Layer5: Storage/Queue--------------------------------------------------------------------------------
