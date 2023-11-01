package internal

import "strings"

func ValidateCpf(cpf string) bool {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")
	cpf = strings.ReplaceAll(cpf, " ", "")

	if len(cpf) != 11 {
		return false
	}
	for i := 1; i < 11; i++ {
		if cpf[i] != cpf[0] {
			break
		}
		if i == 10 {
			return false
		}
	}
	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(cpf[i]-'0') * (10 - i)
	}
	remainder := sum % 11
	digit := 0
	if remainder >= 2 {
		digit = 11 - remainder
	}
	if digit != int(cpf[9]-'0') {
		return false
	}
	sum = 0
	for i := 0; i < 10; i++ {
		sum += int(cpf[i]-'0') * (11 - i)
	}
	remainder = sum % 11
	digit = 0
	if remainder >= 2 {
		digit = 11 - remainder
	}
	if digit != int(cpf[10]-'0') {
		return false
	}
	return true
}
