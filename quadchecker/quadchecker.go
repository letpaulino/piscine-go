// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"os"
	"bufio"
)

// QuadA
func QuadA(x, y int) string {

	ret_str := ""
	if !(x > 0 && y > 0) {
		return ret_str
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if (i == 0 && j == 0) || (i == 0 && j == x-1) || (j == 0 && i == y-1) || (i == y-1 && j == x-1) {
				ret_str = ret_str + "o"
			}
			if (i == 0 || i == y-1) && (j > 0 && j < x-1) {
				ret_str = ret_str + "-"
			}
			if i > 0 && i < y-1 && (j == 0 || j == x-1) {
				ret_str = ret_str + "|"
			}
			if j == x-1 && i != y-1 {
				ret_str = ret_str + "\n"
			}
			if i > 0 && i < y-1 && (j > 0 && j < x-1) {
				ret_str = ret_str + " "
			}
		}
	}
	ret_str = ret_str + "\n"

	return ret_str

}

// QuadB
func QuadB(x, y int) string {

	ret_str := ""
	if !(x > 0 && y > 0) {
		return ret_str
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if (i == 0 && j == 0) || (i == y-1 && j == x-1 && i != 0 && j > 0) {
				ret_str = ret_str + "/"
			}
			if (i == 0 && j == x-1 && j > 0) || (j == 0 && i == y-1 && i != 0) {
				ret_str = ret_str + "\\"
			}
			if (i == 0 || i == y-1) && (j > 0 && j < x-1) {
				ret_str = ret_str + "*"
			}
			if i > 0 && i < y-1 && (j == 0 || j == x-1) {
				ret_str = ret_str + "*"
			}
			if j == x-1 && i != y-1 {
				ret_str = ret_str + "\n"
			}
			if i > 0 && i < y-1 && (j > 0 && j < x-1) {
				ret_str = ret_str + " "
			}
		}
	}
	ret_str = ret_str + "\n"

	return ret_str
}


// QuadC
func QuadC(x, y int) string {
	
	ret_str := ""
	if !(x > 0 && y > 0) {
		return ret_str
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if (i == 0 && j == 0) || (i == 0 && j == x-1) {
				ret_str = ret_str + "A"
			}
			if (j == 0 && i == y-1 && i > 0) || (i == y-1 && j == x-1 && i > 0) {
				ret_str = ret_str + "C"
			}
			if (i == 0 || i == y-1) && (j > 0 && j < x-1) {
				ret_str = ret_str + "B"
			}
			if i > 0 && i < y-1 && (j == 0 || j == x-1) {
				ret_str = ret_str + "B"
			}
			if j == x-1 && i != y-1 {
				ret_str = ret_str + "\n"
			}
			if i > 0 && i < y-1 && (j > 0 && j < x-1) {
				ret_str = ret_str + " "
			}
		}
	}
	ret_str = ret_str + "\n"

	return ret_str
}

// QuadD
func QuadD(x, y int) string {
	
	ret_str := ""
	if !(x > 0 && y > 0) {
		return ret_str
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if (i == 0 && j == 0) || (i == y-1 && j == 0) {
				ret_str = ret_str + "A"
			}
			if (i == 0 && j == x-1 && j > 0) || (i == y-1 && j == x-1 && i > 0 && j > 0) {
				ret_str = ret_str + "C"
			}
			if (i == 0 || i == y-1) && (j > 0 && j < x-1) {
				ret_str = ret_str + "B"
			}
			if i > 0 && i < y-1 && (j == 0 || j == x-1) {
				ret_str = ret_str + "B"
			}
			if j == x-1 && i != y-1 {
				ret_str = ret_str + "\n"
			}
			if i > 0 && i < y-1 && (j > 0 && j < x-1) {
				ret_str = ret_str + " "
			}
		}
	}
	ret_str = ret_str + "\n"

	return ret_str
}

// QuadE
func QuadE(x, y int) string {
	ret_str := ""
	if !(x > 0 && y > 0) {
		return ret_str
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if (i == 0 && j == 0) || (i == y-1 && j == x-1 && i > 0 && j > 0) {
				ret_str = ret_str + "A"
			}
			if (j == 0 && i == y-1 && i > 0) || (i == 0 && j == x-1 && j > 0) {
				ret_str = ret_str + "C"
			}
			if (i == 0 || i == y-1) && (j > 0 && j < x-1) {
				ret_str = ret_str + "B"
			}
			if i > 0 && i < y-1 && (j == 0 || j == x-1) {
				ret_str = ret_str + "B"
			}
			if j == x-1 && i != y-1 {
				ret_str = ret_str + "\n"
			}
			if i > 0 && i < y-1 && (j > 0 && j < x-1) {
				ret_str = ret_str + " "
			}
		}
	}
	ret_str = ret_str + "\n"

	return ret_str
}


func main() {

	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Read input line by line
	text:= ""
	for scanner.Scan() {
		text = text + scanner.Text() + "\n" // Get the current line of text. '\n' is needed since Text supress from input
	}
	
	// fmt.Println("You entered:", text) // Uncomment to see what you have input
	
	// Exception handling - always good in case something goes wrong...
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	input_str := text
	//input_str := "o-o\n| |\no-o\n" // --> should output [quadA] [3] [3]
	// We should read this string from the standard input
	// uncomment or comment to see the output (leave only one uncommented!)
	//input_str := "o---o\n|   |\no---o\n" // --> should output: [quadA] [5] [3]
	//input_str := "o-o\n| |\no-o\n" // --> should output [quadA] [3] [3]
	//input_str := "o---o\n|   |\n|   |\n|   |\n|   |\no---o\n" // --> should output: [quadA] [5] [6]
	//input_str := "duifhoasduih" //--> should output: "Not a quad function\n"
	//input_str := "o\n" \\ --> should output [quadB] [1] [1]
	//input_str := "/***\\\n*   *\n\\***/\n" // --> should output: [quadB] [5] [3]
	//input_str := "/*\\\n* *\n* *\n\\*/\n" // --> should output: [quadB] [3] [4]
	//input_str := "ABA\nB B\nCBC\n" //--> should output: [quadC] [3] [3]
	//input_str := "ABC\nB B\nABC\n"
	//input_str := "ABC\nB B\nCBA\n"  
	//input_str := "A\n"
	//input_str := "A\nC\n"
	//go run main.go $'ABC\nB B\nCBA\n'
	

	/*
	if len(os.Args) == 2 {
		input_str = os.Args[1]
	} else {
		fmt.Println("Incorrect arguments")
		fmt.Println(os.Args[1])
		return 
	}
	*/

	// lembre de checar que tem um so argumento na entrada e printar "Too many arguments" se for >1 ou "Not enough arguments" se for 0
	// We'll find the height of square counting the \n s in the string
	// count "\n"s:
	y := 0 // initialize with zero, but we should increment when we find an \n in the string
	for i := 0; i < len(input_str); i++ {
		if input_str[i] == '\n' {
			y = y + 1 // read: the new value of "y" is y+1
		}
	}

	// We'll find the width of the square by counting in the first line. So we scan the string until we find a \n
	x := 0
	// read: for each index i the string...
	for j := 0; j < len(input_str); j++ {
		// read: if the jth element in the string is equal to \n, break the for loop (but assign j to y first)
		if input_str[j] == '\n' {
			x = j
			break
		}
	}

	// now we have in x and y variables the values we assume that the input_str has...
	//fmt.Printf(QuadA(x, y))
	//fmt.Printf(QuadB(x, y))
	//fmt.Printf(QuadC(x, y))
	//fmt.Printf(QuadD(x, y))
	//fmt.Printf(QuadE(x, y))
	// Now for each type of quad, we generate a square, based on the widht and height we found and then we compare with
	// what received from the user.
	// comparar pra saber qual o tipo: A B C D ou E
	
	unique := true
	ret_str := ""
	if input_str == QuadA(x, y) {
		ret_str = ret_str + fmt.Sprintf("[quadA] [%d] [%d]", x, y)
		unique = false
	}
	
	if input_str == QuadB(x, y) {
		ret_str = ret_str + fmt.Sprintf("[quadB] [%d] [%d]", x, y)
		unique = false
	} 
	if input_str == QuadC(x, y) {
		ret_str = ret_str + fmt.Sprintf("[quadC] [%d] [%d]", x, y)
		unique = false
	}
	
	if input_str == QuadD(x, y) {
		if unique == false {
			ret_str = ret_str + " || "
		}
		ret_str = ret_str + fmt.Sprintf("[quadD] [%d] [%d]", x, y)
		unique = false
	}
	
	if input_str == QuadE(x, y) {
		if unique == false {
			ret_str = ret_str + " || "
		}
		ret_str = ret_str + fmt.Sprintf("[quadE] [%d] [%d]", x, y)
		unique = false
	} 
	ret_str = ret_str + "\n"
	
	if unique == true {
		fmt.Printf("Not a quad function\n")
	} else {
		fmt.Printf(ret_str)
	}
}
