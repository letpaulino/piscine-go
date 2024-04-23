// Write a function that will be formatted as below.
// unc DivMod(a int, b int, div *int, mod *int) { }
// This function will divide the int a and b.
// The result of this division will be stored in the int pointed by div.
// The remainder of this division will be stored in the int pointed by mod.
package piscine

// DivMod divide o número a pelo número b e armazena o resultado da divisão no int apontado por div.
// Armazena o resto da divisão no int apontado por mod.
func DivMod(a int, b int, div *int, mod *int) {
	// Divide a pelo b e atribui o resultado ao int apontado por div.
	*div = a / b
	// Calcula o resto da divisão de a por b e atribui o resultado ao int apontado por mod.
	*mod = a % b
}
