package main


import "testing"

func TestOla (t * testing.T){
	t.Run("diz olá para as pessoas" , func(t * testing.T) {
		resultado := Ola("Lucas")
		esperado := "Olá, Lucas"

		if resultado != esperado{
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}) 

	t.Run("diz olá mundo quando a string estiver vazia", func(t *testing.T) {
		resultado := Ola("")
		esperado := "Olá, mundo"

		if resultado != esperado{
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})
}