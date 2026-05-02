package helloworld

import "testing"

func TestOla(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Lucas")
		esperado := "Olá, Lucas"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz olá mundo quando a string estiver vazia", func(t *testing.T) {
		resultado := Ola("")
		esperado := "Olá, Mundo"

		verificaMensagemCorreta(t, resultado, esperado)
	})
}
