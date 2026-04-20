package arrays

import (
	"reflect"
	"testing"
)


func TestArrays (t *testing.T){

	t.Run("Com tamanho 5 numeros" , func(t *testing.T) {
		numbers := []int{1,2,3,4,5}  
		resultado := SomaArrays(numbers)
		esperado := 15

		if resultado != esperado {
			t.Errorf("esperado %d , resultado %d" , esperado , resultado)
		}
	}) 

	t.Run("Com tamanho variado", func(t *testing.T) {
		numbers := []int{1,2,3}
		resultado := SomaArrays(numbers)
		esperado := 6

		if resultado != esperado {
			t.Errorf("esperado %d , resultado %d" , esperado , resultado)
		}

	})
	
}


func TestSomaTudo (t *testing.T){
	resultado := SomaTudo([]int{1,2}, []int{3,4})
	esperado := []int{3,7}

	if reflect.DeepEqual(resultado,esperado) {
        t.Errorf("resultado %v esperado %v", resultado, esperado)
    }
}