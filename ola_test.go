package main


import "testing"

func TestOla (t * testing.T){
	resultado := Ola("Lucas")
	esperado := "Olá, Lucas"

	if resultado != esperado{
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}