package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMatrix() []string {
	file, err := os.Open("./day3/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var matrix = []string{}

	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}

	return matrix
}

// isValidPart takes the start of a number and the length, and determines if it is adjacent to a symbol
func isValidPart(matrix []string, x int, y int, length int) bool {

	symbols := "!@#$%^&*/=+-\\"
	topEdge := max(0, y-1)
	bottomEdge := min(len(matrix)-1, y+1)
	leftEdge := max(0, x-1)
	rightEdge := min(len(matrix[0])-1, x+length+1)

	contents := matrix[topEdge][leftEdge:rightEdge] + matrix[y][leftEdge:rightEdge] + matrix[bottomEdge][leftEdge:rightEdge]

	return strings.ContainsAny(contents, symbols)
}

// getPartValue takes the x,y of the start of a number and determines the int value
// 0 is returned if the part is not adjacent to a symbol
func getPartvalue(matrix []string, x int, y int) (int, int) {

	// find entire part number
	var partNumber string
	for _, c := range matrix[y][x:] {
		if (c >= '0') && (c <= '9') {
			partNumber += string(c)
		} else {
			// number has ended
			break
		}
	}

	num, err := strconv.Atoi(partNumber)
	if err != nil {
		panic("error reading number (x,y): (" + fmt.Sprint(x) + ", " + fmt.Sprint(y) + ")")
	}

	valid := isValidPart(matrix, x, y, len(partNumber))

	if !valid {
		fmt.Println("invalid:", num)
		num = 0
	}

	return num, len(partNumber)
}

func Part1() {
	matrix := getMatrix()

	sum := 0

	for y, line := range matrix {
		for x := 0; x < len(line); x++ {
			c := line[x]
			if (c >= '0') && (c <= '9') {
				val, length := getPartvalue(matrix, x, y)
				sum += val
				x += length
			}
		}
	}
	fmt.Println(sum)
}
