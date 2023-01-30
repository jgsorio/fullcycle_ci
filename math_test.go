package main

import "testing"

func SomaTest(t *testing.T) {
	total := Soma(10, 10)

	if total != 20 {
		t.Errorf("O resultado esperado está incorreto. Esperado: %d. Recebido: %d", total, 20)
	}
}
