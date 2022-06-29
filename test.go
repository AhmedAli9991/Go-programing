package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int
	Scanner := bufio.NewScanner(os.Stdin)
	Scanner.Scan()
	input := Scanner.Text()

	N, err := strconv.Atoi(input)
	if err != nil {
		main()
	}
	if N < 1 || N > 100 {
		return
	}
	results := repeatForX(N, Scanner)
	if strings.Contains(results, "err") {
		return
	}

	writer := bufio.NewWriter(os.Stdout)
	display(results, writer)
	writer.Flush()

}
func repeatForX(N int, Scanner *bufio.Scanner) string {
	if N != 0 {
		Scanner.Scan()
		i := Scanner.Text()
		X, err := strconv.Atoi(i)
		if err != nil {
			return "err"
		}
		if X < 1 || X > 100 {
			return "err"
		}
		result, err1 := repeatForYn(X, Scanner)
		if err1 == "err" {
			return "err"
		}
		return strconv.Itoa(result) + " " + repeatForX(N-1, Scanner)
	}
	return ""

}
func repeatForYn(X int, Scanner *bufio.Scanner) (int, string) {
	Scanner.Scan()
	yn := Scanner.Text()
	if X > 1 {
		sums, err2 := square(yn, X, 0, Scanner)
		if err2 == "err" {
			return 0, "err"
		}
		return sums, " "

	}

	sqrt, err := strconv.Atoi(yn)
	if err != nil {
		return 0, "err"
	}
	if sqrt < -100 || sqrt > 100 {
		return 0, "err"
	}
	sqrt = sqrt * sqrt
	return sqrt, " "

}
func square(yn string, X int, Y int, Scanner *bufio.Scanner) (int, string) {
	if !strings.Contains(yn, " ") {
		integer, err := strconv.Atoi(yn)
		if err != nil {
			return 0, "err"
		}
		Y = Y + 1
		if integer < -100 || integer > 100 || Y != X {
			return 0, "err"
		}
		if integer < 1 {
			integer = 0
		}
		return integer * integer, " "
	}
	index := strings.Index(yn, " ")
	integer, err := strconv.Atoi(yn[0:index])
	if err != nil {
		return 0, "err"
	}
	if string(yn[index+1]) == " " {
		return 0, "err"
	}
	if integer < -100 || integer > 100 {
		return 0, "err"
	}
	if integer < 1 {
		integer = 0
	}
	yn = yn[index+1:]
	val, err2 := square(yn, X, Y+1, Scanner)
	if err2 == "err" {
		return 0, "err"
	}

	return integer*integer + val, " "
}
func display(out string, writer *bufio.Writer) string {
	if !strings.Contains(out, " ") {
		return ""
	}

	index := strings.Index(out, " ")
	writer.WriteString(out[0:index] + "\n")
	out = out[index+1:]
	return display(out, writer)
}
