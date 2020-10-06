package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
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
	http.HandleFunc("/api/v1/callsomething", l.CallSomeExternalМожноИСпользоватьКириллицуMagicServiceRightFromHere)
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

// CallSomeExternalМожноИСпользоватьКириллицуMagicServiceRightFromHere is doing ...
func (r HandlerVersion) CallSomeExternalМожноИСпользоватьКириллицуMagicServiceRightFromHere(w http.ResponseWriter, req *http.Request) {

	request := "give me please"
	resultFromLogic := r.LogicVersion.CallSomeExternalМожноИСпользоватьКириллицуMagicServiceRightFromHere(request)
	result, err := json.Marshal(resultFromLogic)
	if err != nil {
		io.WriteString(w, "error-6677f788: some error")
		return
	}

	io.WriteString(w, string(result))
}

// Layer3: Logic--------------------------------------------------------------------------------

// LogicVersion ...
type LogicVersion struct {
	Storage StorageVersion
}

// NewLogic constructor
func NewLogic() LogicVersion {
	return LogicVersion{
		Storage: NewStorage(),
	}
}

// GetVersion ...
func (r LogicVersion) GetVersion(request string) ModelVersion {

	result := r.Storage.GetVersion(request)
	return result
}

// CallSomeExternalМожноИСпользоватьКириллицуMagicServiceRightFromHere ...
func (r LogicVersion) CallSomeExternalМожноИСпользоватьКириллицуMagicServiceRightFromHere(request string) ModelVersion {

	res, err := http.Get("https://www.voccent.com/api/v1/version")
	if err != nil {
		log.Fatal(err)
	}
	version, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return ModelVersion{
		Version: string(version),
	}
}

// Layer4: Storage/Queue--------------------------------------------------------------------------------

// StorageVersion ...
type StorageVersion struct {
	DataBaseURL string
}

// NewStorage constructor
func NewStorage() StorageVersion {
	return StorageVersion{}
}

// GetVersion ...
func (r StorageVersion) GetVersion(request string) ModelVersion {

	// Imaging this is coming back from the database.... :-P
	result := ModelVersion{
		Version: "ver1.0.0",
	}

	return result
}

// Layer5: Storage/Queue--------------------------------------------------------------------------------
