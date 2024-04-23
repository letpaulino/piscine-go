package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			z01.PrintRune('\\')
			z01.PrintRune('n')
		} else {
			z01.PrintRune(rune(s[i]))
		}
	}
}
