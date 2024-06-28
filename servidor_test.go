package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestObterJogadores(t *testing.T) {
	t.Run("retornar resultado de Maria", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Maria")
		resposta := httptest.NewRecorder()

		ServidorJogador(resposta, requisicao)

		verificarCorpoRequisicao(t, resposta.Body.String(), "20")
	})

	t.Run("returns Pedro's score", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Pedro")
		resposta := httptest.NewRecorder()

		ServidorJogador(resposta, requisicao)

		verificarCorpoRequisicao(t, resposta.Body.String(), "10")
	})
}

func novaRequisicaoObterPontuacao(nome string) *http.Request {
	requisicao, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/jogadores/%s", nome), nil)
	return requisicao
}

func verificarCorpoRequisicao(t *testing.T, recebido, esperado string) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("corpo da requisição é inválido, obtive '%s' esperava '%s'", recebido, esperado)
	}
}
