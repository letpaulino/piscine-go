package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else {
		if n < 0 {
			z01.PrintRune('-')
		}
		PrintInt(n)
	}
}

func PrintInt(number int) {
	i := '0'
	for j := 1; j <= number%10; j++ {
		i++
	}
	for j := -1; j >= number%10; j-- {
		i++
	}
	if number/10 != 0 {
		PrintInt(number / 10)
	}
	z01.PrintRune(i)
}
