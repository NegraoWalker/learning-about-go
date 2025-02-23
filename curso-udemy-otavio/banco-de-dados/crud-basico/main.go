package main

import (
	"fmt"
	"log"
	"mod-crud-basico/servidor"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriaUsuario).Methods(http.MethodPost)

	fmt.Println("Escutando na porta 5000...")
	log.Fatal(http.ListenAndServe(":5000", router))

}
