package main

import (
	"crypto/sha256"
	"fmt"
)

const salt = "go-2024"

func typeNames(values ...any) []string {
	result := make([]string, 0, len(values))
	for _, value := range values {
		result = append(result, fmt.Sprintf("%T", value))
	}
	return result
}

func concatenateValues(values ...any) string {
	result := ""
	for _, value := range values {
		result += fmt.Sprint(value)
	}
	return result
}

func hashRunesWithSalt(runes []rune, salt string) [32]byte {
	middle := len(runes) / 2
	withSalt := make([]rune, 0, len(runes)+len([]rune(salt)))
	withSalt = append(withSalt, runes[:middle]...)
	withSalt = append(withSalt, []rune(salt)...)
	withSalt = append(withSalt, runes[middle:]...)

	return sha256.Sum256([]byte(string(withSalt)))
}

func main() {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	values := []any{
		numDecimal,
		numOctal,
		numHexadecimal,
		pi,
		name,
		isActive,
		complexNum,
	}

	fmt.Println("Types:", typeNames(values...))

	combined := concatenateValues(values...)
	fmt.Println("Combined:", combined)

	runes := []rune(combined)
	fmt.Println("Runes:", runes)

	hash := hashRunesWithSalt(runes, salt)
	fmt.Printf("SHA-256: %x\n", hash)
}
