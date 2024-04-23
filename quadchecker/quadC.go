package main

import (
	"fmt"
	"os"
	"strconv"
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

    // string to int
    x, err := strconv.Atoi(os.Args[1])
	y, err := strconv.Atoi(os.Args[2])

    if err != nil {
        // ... handle error
        panic(err)
    }

    fmt.Printf(QuadC(x, y))
}
