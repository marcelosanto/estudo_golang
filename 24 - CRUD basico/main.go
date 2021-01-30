package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//CRUD - CREATE - READ - UPDATE - DELETE

	router := mux.NewRouter()

	fmt.Println("Escutando na porta: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
