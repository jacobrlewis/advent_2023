package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type game struct {
	Id       int
	MinRed   int
	MinBlue  int
	MinGreen int
}

func (g *game) handleInformation(num string, color string) {
	
	amount, err := strconv.Atoi(num)
	if err != nil {
		panic("error reading cube number: " + num)
	}

	switch color {
	case "red":
		g.MinRed = max(g.MinRed, amount)
	case "blue":
		g.MinBlue = max(g.MinBlue, amount)
	case "green":
		g.MinGreen = max(g.MinGreen, amount)
	}
}

func parseLine(line string) game {

	fmt.Println(line)

	var game game

	parts := strings.Split(line, ":")

	idPart := parts[0]
	idString := strings.Split(idPart, " ")
	id, err := strconv.Atoi(idString[1])
	if err != nil {
		panic("Error reading game id. " + err.Error())
	}

	game.Id = id

	setsList := parts[1]
	sets := strings.Split(setsList, ";")

	for _, set := range sets {
		cubesList := strings.Split(set, ",")
		for _, cubeStr := range cubesList {
			cubeStr = strings.Trim(cubeStr, " ")
			tokens := strings.Split(cubeStr, " ")
			num := tokens[0]
			color := tokens[1]

			game.handleInformation(num, color)
		}
	}

	fmt.Printf("%+v\n", game)

	return game
}

func getGames() []game {
	file, err := os.Open("./day2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var games []game
	for scanner.Scan() {
		line := scanner.Text()
		games = append(games, parseLine(line))
	}
	return games
}

func Part1() {
	games := getGames()

	MAX_RED := 12
	MAX_GREEN := 13
	MAX_BLUE := 14

	sum := 0
	for _, game := range games {
		if game.MinBlue <= MAX_BLUE &&
			game.MinGreen <= MAX_GREEN &&
			game.MinRed <= MAX_RED {
			sum += game.Id
		}
	}
	fmt.Println(sum)
}

func Part2() {
	games := getGames()
	sum := 0
	for _, game := range games {
		sum += game.MinBlue * game.MinGreen * game.MinRed
	}
	fmt.Println(sum)
}
