package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Contains(str string, ch rune) bool {
	for _, c := range str {
		if c == ch {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}

	args := os.Args[1:]
	vowels := "aeiouAEIOU"
	var vowelIndices []int
	var nonVowelChars []rune

	// Encontra os índices das vogais em cada argumento
	for _, arg := range args {
		for i, char := range arg {
			if Contains(vowels, char) {
				vowelIndices = append(vowelIndices, i)
			} else {
				nonVowelChars = append(nonVowelChars, char)
			}
		}
	}

	// Inicializa os argumentos modificados com as vogais trocadas
	modifiedArgs := make([]string, len(args))

	// Troca as vogais nas palavras e mantém a ordem dos outros caracteres
	for i, arg := range args {
		modifiedArg := []rune(arg)
		for j, charIndex := range vowelIndices {
			if j >= len(modifiedArg) {
				break
			}
			if charIndex < len(modifiedArg) {
				modifiedArg[charIndex] = nonVowelChars[j]
			}
		}
		modifiedArgs[i] = string(modifiedArg)
	}

	// Imprime os argumentos modificados
	for _, arg := range modifiedArgs {
		z01.PrintRune(' ')
		for _, char := range arg {
			z01.PrintRune(char)
		}
	}

	z01.PrintRune('\n')
}
