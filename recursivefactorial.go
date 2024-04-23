// Write a recursive function that returns the factorial of the int passed as parameter.
// Errors (non possible values or overflows) will return 0.for is forbidden for this exercise.
package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 21 {
		return 0
	} else if nb < 1 {
		return 1
	}
	return RecursiveFactorial(nb-1) * nb
}
