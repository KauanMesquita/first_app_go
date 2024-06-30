package main

import (
	"log"
	"net/http"
)

func main() {
	servidor := &ServidorJogador{}

	if err := http.ListenAndServe(":500", servidor); err != nil {
		log.Fatalf("Não foi possível escutar na porta 500 %v", err)
	}
}
