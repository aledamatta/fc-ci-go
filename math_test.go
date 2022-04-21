package main

import "testing"

func TestSoma(t *testing.T) {
	if Somar(10, 10) != 20 {
		t.Error("O valor deve ser 20")
	}
}
