package main

import "testing"

func TestSoma(t *testing.T) {
	if somar(10, 10) != 20 {
		t.Error("O valor deve ser 20")
	}
}
