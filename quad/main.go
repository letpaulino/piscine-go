package main

import (
	piscinea "piscine/quada"
	piscineb "piscine/quadb"
	piscinec "piscine/quadc"
	piscined "piscine/quadd"
	piscinee "piscine/quade"

	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune('A')
	z01.PrintRune('\n')
	piscinea.QuadA(5, 3)
	piscinea.QuadA(5, 1)
	piscinea.QuadA(1, 1)
	piscinea.QuadA(1, 5)
	z01.PrintRune('\n')

	z01.PrintRune('B')
	z01.PrintRune('\n')
	piscineb.QuadB(5, 3)
	piscineb.QuadB(5, 1)
	piscineb.QuadB(1, 1)
	piscineb.QuadB(1, 5)
	z01.PrintRune('\n')

	z01.PrintRune('C')
	z01.PrintRune('\n')
	piscinec.QuadC(5, 3)
	piscinec.QuadC(5, 1)
	piscinec.QuadC(1, 1)
	piscinec.QuadC(1, 5)
	z01.PrintRune('\n')

	z01.PrintRune('D')
	z01.PrintRune('\n')
	piscined.QuadD(5, 3)
	piscined.QuadD(5, 1)
	piscined.QuadD(1, 1)
	piscined.QuadD(3, 5)
	z01.PrintRune('\n')

	z01.PrintRune('E')
	z01.PrintRune('\n')
	piscinee.QuadE(5, 3)
	piscinee.QuadE(5, 1)
	piscinee.QuadE(1, 1)
	piscinee.QuadE(1, 5)

}
