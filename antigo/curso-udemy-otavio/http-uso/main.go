package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) { //Rota criada para requisições para a porta 5000
		w.Write([]byte("Página Home"))
	})

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) { //Rota criada para requisições para a porta 5000
		w.Write([]byte("Página de Usuários"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil)) //Escutando requisições para a porta 5000 do servidor(no caso minha máquina)

}
