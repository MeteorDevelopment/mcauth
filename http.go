package main

import (
	"encoding/json"
	"log"
	"mcauth/codes"
	"net/http"
	"strconv"
)

func setupHttpServer() {
	http.HandleFunc("GET /retrieve/{code}", onRetrieve)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func onRetrieve(w http.ResponseWriter, r *http.Request) {
	code, err := strconv.Atoi(r.PathValue("code"))
	if err != nil {
		sendJson(w, map[string]string{"error": codes.ErrInvalidCode.Error()}, http.StatusBadRequest)
		return
	}

	id, err := codes.Retrieve(code)
	if err != nil {
		sendJson(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	sendJson(w, map[string]string{"uuid": id.String()}, http.StatusOK)
}

func sendJson(w http.ResponseWriter, value any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(value)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
