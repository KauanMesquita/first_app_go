package main

import (
	"log"
	"net/http"
)

func main() {
	tratador := http.HandlerFunc(ServidorJogador)
	if err := http.ListenAndServe(":500", tratador); err != nil {
		log.Fatalf("Não foi possível escutar na porta 500 %v", err)
	}
}
