package structs

import "testing"


func TestPerimetro(t *testing.T){
	retangulo := Retangulo{10.0 , 20.0}
	resultado := Perimetro(retangulo)
	esperado := 60.0

	if resultado != esperado {
        t.Errorf("resultado %.2f esperado %.2f", resultado, esperado)
    }
}

func TestArea (t *testing.T){
	retangulo := Retangulo{10.0 , 20.0}
	resultado := Area(retangulo)
	esperado := 200.0

	if resultado != esperado {
        t.Errorf("resultado %.2f esperado %.2f", resultado, esperado)
    }
}