package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, visitor")
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{ID: 1, Description: "Teste 1"},
		Todo{ID: 2, Description: "Teste 2"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func GetTodosByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	testID := vars["id"]
	fmt.Fprintf(w, "You are on Todo ID %s", testID)
}

func PostTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not implemented yet.")
}

func PutTodosByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not implemented yet.")
}

func DeleteTodosByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not implemented yet.")
}
