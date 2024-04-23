// Write a recursive function that returns the value at the position index in the fibonacci sequence.
// The first value is at index 0.
// The sequence starts this way: 0, 1, 1, 2, 3 etc...
// A negative index will return -1. for is forbidden for this exercise.
package piscine

func IterativeFactorial(nb int) int {
	if nb == 0 {
		return 1
	}

	factorial := 1

	if nb > 0 && nb < 21 {
		for i := 1; i <= nb; i++ {
			factorial *= i
		}
	} else {
		return 0
	}

	return factorial
}
