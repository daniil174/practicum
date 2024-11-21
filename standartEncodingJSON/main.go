package main

import (
	"encoding/json"
	"net/http"
)

type Visitor struct {
	ID      int      `json:"id"`
	Name    string   `json:"name,omitempty"`
	Phones  []string `json:"phones,omitempty"`
	Company string   `json:"company,omitempty"`
}

var visitors = map[string]Visitor{
	"1": {
		ID:   1,
		Name: "Guest",
		Phones: []string{
			`789-673-56-90`,
			`612-934-77-23`,
		}},
}

func JSONHandler(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	resp, err := json.Marshal(visitors[id])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func main() {
	http.ListenAndServe("localhost:8080", http.HandlerFunc(JSONHandler))
}
