package main

import (
	"log"
	"net/http"
)

type ArmazenamentoJogadorMemoria struct{}

func (a *ArmazenamentoJogadorMemoria) ObterPontuacaoJogador(nome string)int{
	return 123
}


func main() {
	servidor := &ServidorJogador{&ArmazenamentoJogadorMemoria{}}
	tratador := http.HandlerFunc(servidor.ServerHTTP)
	if err := http.ListenAndServe(":5000", tratador); err != nil {
		log.Fatalf("não foi possivel escutar na porta 5000 %v", err)
	}
}
