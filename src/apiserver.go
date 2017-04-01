package main

import (
	"encoding/json"
	"net/http"

	"data"
	"handler"
	"logger"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")
	apiRequest := &data.ApiRequest{}
	err := json.Unmarshal([]byte(query), apiRequest)
	if nil != err {
		w.Write([]byte(err.Error()))
		return
	}
	response := handler.HandlerApiRequest(apiRequest)
	w.Write(response)

}

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello shital"))
}

func main() {
	loggerConfig := &logger.LoggerConfig{}
	logger.Init(loggerConfig)
	http.HandleFunc("/apiserver", handleRequest)
	http.HandleFunc("/hello", handleHello)
	err := http.ListenAndServe(":3030", nil)
	logger.Debug("%s", err)
	logger.Debug("API server started successfully.")

}
