package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado esperado: %d, Resultado obtido: %d", 30, total)
	}
}

func TestSubtrai(t *testing.T) {
	total := Subtrai(30, 15)

	if total != 15 {
		t.Errorf("Resultado da subtração é inválido: Resultado esperado: %d, Resultado obtido: %d", 15, total)
	}
}

func TestMultiplica(t *testing.T) {
	total := Multiplica(5, 6)

	if total != 30 {
		t.Errorf("Resultado da multiplicação é inválido: Resultado esperado: %d, Resultado obtido: %d", 30, total)
	}
}

func TestDivide(t *testing.T) {
	total := Divide(30, 2)

	if total != 15 {
		t.Errorf("Resultado da divisão é inválido: Resultado esperado: %d, Resultado obtido: %d", 15, total)
	}
}

func TestSoma2(t *testing.T) {
	total := Soma2(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado esperado: %d, Resultado obtido: %d", 30, total)
	}
}

func TestSubtrai2(t *testing.T) {
	total := Subtrai2(30, 15)

	if total != 15 {
		t.Errorf("Resultado da subtração é inválido: Resultado esperado: %d, Resultado obtido: %d", 15, total)
	}
}

func TestMultiplica2(t *testing.T) {
	total := Multiplica2(5, 6)

	if total != 30 {
		t.Errorf("Resultado da multiplicação é inválido: Resultado esperado: %d, Resultado obtido: %d", 30, total)
	}
}

func TestDivide2(t *testing.T) {
	total := Divide2(30, 2)

	if total != 15 {
		t.Errorf("Resultado da divisão é inválido: Resultado esperado: %d, Resultado obtido: %d", 15, total)
	}
}
