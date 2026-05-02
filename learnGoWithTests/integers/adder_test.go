package integers

import (
	"testing"
)

func TestAdd(t *testing.T) {
	soma := Adder(2, 2)
	esperado := 4

	if soma != esperado {
		t.Errorf("esperado '%d , resultado '%d", esperado, soma)
	}
}
