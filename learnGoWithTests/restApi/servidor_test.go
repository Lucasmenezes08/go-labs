package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ListenAndServer(endereco string, tratador *Handler) error {
	return nil
}

type Handler interface {
	ServerHTTP(http.ResponseWriter, *http.Request)
}

type MockArmazenamentoJogador struct{
	scores map[string]int
}

func (m *MockArmazenamentoJogador) ObterPontuacaoJogador(nome string)int{
	score := m.scores[nome]
	return score
}

func TestObterJogadores(t *testing.T) {
	armazenamento := MockArmazenamentoJogador{
		map[string]int{
			"maria": 20,
			"pedro": 10,
		},
	}
	servidor := &ServidorJogador{Armazenamento: &armazenamento}
	t.Run("Deve retornar resultado de maria", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("maria")
		resposta := httptest.NewRecorder()

		servidor.ServerHTTP(resposta, requisicao)

		verificarStatusCodigo(t, resposta.Code, http.StatusOK)
		verificarRequisicaoCorreta(t, resposta.Body.String(), "20")
	})

	t.Run("Deve retornar resultado de pedro", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("pedro")
		resposta := httptest.NewRecorder()

		servidor.ServerHTTP(resposta, requisicao)

		verificarStatusCodigo(t, resposta.Code, http.StatusOK)
		verificarRequisicaoCorreta(t, resposta.Body.String(), "10")	
	})

	t.Run("Deve retornar not found para rota não encontrada", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("jorge")
		resposta := httptest.NewRecorder()

		servidor.ServerHTTP(resposta, requisicao)

		recebido := resposta.Code
		esperado := http.StatusNotFound

		if recebido != esperado{
			t.Errorf("recebido status %d esperado %d", recebido, esperado)
		}
	})
}



func novaRequisicaoObterPontuacao(nome string) *http.Request {
	requisicao, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/jogadores/%s", nome), nil)
	return requisicao
}

func verificarRequisicaoCorreta(t *testing.T, recebido string, esperado string) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("recebido '%s', esperado '%s'", recebido, esperado)
	}
}

func verificarStatusCodigo(t *testing.T, recebido , esperado int) {
	t.Helper()
	if recebido != esperado {
    	t.Errorf("não recebeu código de status HTTP esperado, recebido %d, esperado %d", recebido, esperado)
    }
}


