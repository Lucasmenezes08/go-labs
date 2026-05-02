package iteration

import "testing"

func TestRepeat(t *testing.T) {
	result := Repeat("a", 5)
	esperado := "aaaaa"

	if result != esperado {
		t.Errorf("esperado %s , resultado: %s", esperado, result)
	}
}
