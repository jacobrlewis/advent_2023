package day1

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func runeToInt(c rune) (int, error) {
	if (c >= '0') && (c <= '9') {
		return int(c - '0'), nil
	}
	return 0, errors.New("Invalid rune")
}

func Part1() {
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		var first int
		var last int
		line := scanner.Text()

		for _, c := range line {
			digit, err := runeToInt(c)
			if err == nil {
				first = digit * 10
				break
			}
		}

		for i := range line {
			c := rune(line[len(line)-(i+1)])
			digit, err := runeToInt(c)
			if err == nil {
				last = digit
				break
			}
		}

		sum += first + last
	}

	fmt.Println(sum)
}

func Part2() {
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		first := getFirstDigit(line) * 10
		last := getLastDigit(line)
		fmt.Println("= " + fmt.Sprint(first+last))
		sum += first+last
	}
	fmt.Println(sum)
}

var wordList = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

// startWithNumberWord returns true and the integer value if a string begins with a number spelled out in lowercase
func startsWithNumberWord(s string) (bool, int) {
	for i, word := range wordList {
		if strings.HasPrefix(s, word) {
			return true, i+1
		}
	}
	return false, 0
}

func getFirstDigit(line string) int {
	if len(line) == 0 {
		panic("getFirstDigit found no digits")
	}

	digit, err := runeToInt(rune(line[0]))
	if err == nil {
		return digit
	}

	found, digit := startsWithNumberWord(line)
	if found {
		return digit
	}

	return getFirstDigit(line[1:])
}

func getLastDigit(line string) int {
	for i := range line {

		c := rune(line[len(line)-(i+1)])
		digit, err := runeToInt(c)
		if err == nil {
			return digit
		}

		endOfLine := line[len(line)-(i+1):]
		found, digit := startsWithNumberWord(endOfLine)
		if found {
			return digit
		}
	}
	panic("getLastDigit found no digits")
}
