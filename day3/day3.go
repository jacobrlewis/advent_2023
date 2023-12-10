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

// seekNumber takes an x coordinate of a digit and seeks left and right to get the entire value of the number, and the x coordinate of the last digit
func seekNumber(line string, x int) (int, int) {
	strNum := string(line[x])
	rightEdge := x

	// seek left
	for i := x - 1; i >= 0; i-- {
		c := line[i]
		if c >= '0' && c <= '9' {
			strNum = string(c) + strNum
		} else {
			break
		}
	}

	// seek right
	for i := x + 1; i <= len(line)-1; i++ {
		c := line[i]
		if c >= '0' && c <= '9' {
			strNum = strNum + string(c)
			rightEdge = i
		} else {
			break
		}
	}

	value, err := strconv.Atoi(strNum)
	if err != nil {
		panic(err)
	}
	return value, rightEdge
}

// gearValue returns the value of a gear (*) at the given x,y coord
func gearValue(matrix []string, x int, y int) int {

	var neighbors []int
	leftEdge := max(0, x-1)
	rightEdge := min(len(matrix[0])-1, x+1)

	if y != 0 {
		// check top
		for i, c := range matrix[y-1][leftEdge:rightEdge+1] {
			if c >= '0' && c <= '9' {
				fmt.Printf("seekNumber at %d, %d\n", x, y-1)
				value, lastPlace := seekNumber(matrix[y-1], x+i-1)
				neighbors = append(neighbors, value)

				// avoid adding the same number twice
				if lastPlace >= x {
					break
				}
			}
		}
	}

	if x != 0 {
		// check left
		c := matrix[y][x-1]
		if c >= '0' && c <= '9' {
			fmt.Printf("seekNumber at %d, %d\n", x-1, y)
			value, _ := seekNumber(matrix[y], x-1)
			neighbors = append(neighbors, value)
		}
	}

	if x != len(matrix[0])-1 {
		// check right
		c := matrix[y][x+1]
		if c >= '0' && c <= '9' {
			fmt.Printf("seekNumber at %d, %d\n", x+1, y)
			value, _ := seekNumber(matrix[y], x+1)
			neighbors = append(neighbors, value)
		}
	}

	if y != len(matrix)-1 {
		// check bottom
		for i, c := range matrix[y+1][leftEdge:rightEdge+1] {
			if c >= '0' && c <= '9' {
				fmt.Printf("seekNumber at %d, %d\n", x+i-1, y+1)
				value, lastPlace := seekNumber(matrix[y+1], x+i-1)
				neighbors = append(neighbors, value)

				// avoid adding the same number twice
				if lastPlace >= x {
					break
				}
			}
		}
	}

	if len(neighbors) > 1 {
		fmt.Printf("Valid gear at %d, %d. Neighbors: %v\n", x, y, neighbors)
		result := 1
		for _, n := range neighbors {
			result *= n
		}
		return result
	}
	return 0
}

func Part2() {
	matrix := getMatrix()

	sum := 0

	for y, line := range matrix {
		for x, c := range line {
			if c == '*' {
				fmt.Printf("Found * at %d, %d\n", x, y)
				sum += gearValue(matrix, x, y)
			}
		}
	}

	fmt.Println(sum)
}
