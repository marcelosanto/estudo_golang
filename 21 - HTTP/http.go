package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar pagina de usuarios"))
}

func main() {
	// HTTP é um protocolo de comunicação - base da comunicação web

	// cliente (faz a requisição) <-> servidor (processa requisição e envia resposta)

	// Request - Response

	// Rotas
	// URI - Identificador do Recurso
	// Método - GET / Post / Put / Delete

	http.HandleFunc("/", home)
	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
