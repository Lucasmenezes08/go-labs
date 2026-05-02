package main

import (
	"fmt"
	"net/http"
)

type ArmazenamentoJogador interface {
	ObterPontuacaoJogador(string) int
}

type ServidorJogador struct {
	Armazenamento ArmazenamentoJogador
}

func (s *ServidorJogador) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.registrarVitoria(w)
	case http.MethodGet:
		s.mostrarPontuacaoJogador(w,r)
	}
}


func (s *ServidorJogador) mostrarPontuacaoJogador(w http.ResponseWriter, r *http.Request){
	jogador := r.URL.Path[len("/jogadores/"):]

	pontuacao := s.Armazenamento.ObterPontuacaoJogador(jogador)

	if pontuacao == 0{
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, pontuacao)
}

func (s *ServidorJogador) registrarVitoria(w http.ResponseWriter){
	w.WriteHeader(http.StatusAccepted)
}

func ObterPontuacaoJogador(nome string) string {
	if nome == "maria" {
		return ("20")
	}
	if nome == "pedro" {
		return ("10")
	}

	return ""
}
