package pointers

import "testing"

func TestPointer(t *testing.T){

	confirmarSaldo := func(t *testing.T , carteira Carteira, esperado Bitcoin){
		t.Helper()
		resultado := carteira.Saldo()

		if resultado != esperado{
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}

	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))
		confirmarSaldo(t,carteira,Bitcoin(10))
	})

	t.Run("Retirar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}
		carteira.Retirar(Bitcoin(10))
		confirmarSaldo(t, carteira, Bitcoin(10))
	})


	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {

		saldoInicial := Bitcoin(20)

		carteira := Carteira{saldoInicial}
		err := carteira.Retirar(Bitcoin(100))
		confirmarSaldo(t, carteira, saldoInicial)
		
		if err == nil {
			t.Error("Esperava um erro mas nenhum ocorreu")
		}
	})
	
}




