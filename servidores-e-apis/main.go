package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const port = 8000

type Client struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Address *Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

var clients = make([]Client, 0)

// utilities

func checkForFields(c Client) error {
	if c.Name == "" {
		return fmt.Errorf(`Missing field "name"`)
	}
	if c.Address == nil {
		return fmt.Errorf(`Missing field "address"`)
	}
	if c.Address.City == "" {
		return fmt.Errorf(`Missing field "address"."city"`)
	}
	if c.Address.State == "" {
		return fmt.Errorf(`Missing field "address"."state"`)
	}
	return nil
}

func errorJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(struct {
		Code   int    `json:"code"`
		Detail string `json:"detail"`
	}{
		Code:   code,
		Detail: message,
	})
}

// handlers

func getClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)
}

func getClientByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil || id < 0 {
		errorJSON(w, "Invalid id, must be a positive integer", http.StatusBadRequest)
		return
	}

	for _, c := range clients {
		if c.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(c)
			return
		}
	}

	errorJSON(w, "Client not found", http.StatusNotFound)
}

func createClient(w http.ResponseWriter, r *http.Request) {
	newID := 1

	if len(clients) > 0 {
		newID = clients[len(clients)-1].ID + 1
	}

	var client Client
	err := json.NewDecoder(r.Body).Decode(&client)

	if err != nil {
		errorJSON(w, "Invalid request body", http.StatusUnprocessableEntity)
		return
	}

	err = checkForFields(client)

	if err != nil {
		errorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	client.ID = newID
	clients = append(clients, client)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)
}

func updateClientByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil || id < 0 {
		errorJSON(w, "Invalid id, must be a positive integer", http.StatusBadRequest)
		return
	}

	var client Client
	err = json.NewDecoder(r.Body).Decode(&client)

	if err != nil {
		errorJSON(w, "Invalid request body", http.StatusUnprocessableEntity)
		return
	}

	err = checkForFields(client)

	if err != nil {
		errorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	targetIndex := -1
	for i, c := range clients {
		if c.ID == id {
			targetIndex = i
			client.ID = c.ID
			break
		}
	}

	if targetIndex == -1 {
		errorJSON(w, "Client not found", http.StatusNotFound)
		return
	}

	clients[targetIndex] = client

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(client)
}

func deleteClientByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		errorJSON(w, "Invalid id, must be an integer", http.StatusBadRequest)
		return
	}

	for i, c := range clients {
		if c.ID == id {
			clients = append(clients[:i], clients[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	errorJSON(w, "Client not found", http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/clients", getClients).Methods(http.MethodGet)
	r.HandleFunc("/clients", createClient).Methods(http.MethodPost)
	r.HandleFunc("/clients/{id}", getClientByID).Methods(http.MethodGet)
	r.HandleFunc("/clients/{id}", updateClientByID).Methods(http.MethodPut)
	r.HandleFunc("/clients/{id}", deleteClientByID).Methods(http.MethodDelete)

	fmt.Printf("Listening on http://localhost:%d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		log.Fatal(err)
	}
}
