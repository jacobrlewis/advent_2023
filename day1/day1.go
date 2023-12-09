package day1

import (
	"bufio"
	"fmt"
	"os"
)

func ParseFile() {
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
			if (c >= '0') && (c <= '9') {
				first = int(c-'0') * 10
				break
			}
		}

		for i := range line {
			c := line[len(line)-(i+1)]
			if (c >= '0') && (c <= '9') {
				last = int(c - '0')
				break
			}
		}

		sum += first + last
	}

	fmt.Println(sum)
}
