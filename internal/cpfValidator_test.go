package internal

import "testing"

func TestCpf(t *testing.T) {
	cpf := "01234567890"

	isValidCpf := ValidateCpf(cpf)

	if !isValidCpf {
		t.Errorf("Must be a valid CPF!")
	}
}

func TestCpfWithMask(t *testing.T) {
	cpf := "012.345.678-90"

	isValidCpf := ValidateCpf(cpf)

	if !isValidCpf {
		t.Errorf("Must be a valid CPF!")
	}
}

func TestCpfInvalidDigit(t *testing.T) {
	cpf := "958.187.055-00"

	isValidCpf := ValidateCpf(cpf)

	if isValidCpf {
		t.Errorf("Must be a invalid CPF digit!")
	}
}

func TestCpfInvalid(t *testing.T) {
	cpf := "012.345.678"

	isValidCpf := ValidateCpf(cpf)

	if isValidCpf {
		t.Errorf("Must be a invalid CPF!")
	}
}
