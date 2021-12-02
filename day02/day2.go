package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(Coordinates []string) int {
	x, y := 0, 0

	for _, Coords := range Coordinates{
		splitString := strings.Fields(Coords)
		depth, _ := strconv.Atoi(splitString[1])
		switch splitString[0] {
		case "forward":
			x += depth
		case "up":
			y -= depth
		case "down":
			y += depth
		}
	}
	return x * y
}

func part2(Coordinates []string) int {
	x, y, aim := 0, 0, 0

	for _, Coords := range Coordinates{
		splitString := strings.Fields(Coords)
		depth, _ := strconv.Atoi(splitString[1])
		switch splitString[0] {
		case "forward":
			x += depth
			y += aim * depth
		case "up":
			aim -= depth
		case "down":
			aim += depth
		}
	}
	return x * y
}

func getInput() []string {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func main() {
	Input := getInput()
	fmt.Println("Part one:", part1(Input))
	fmt.Println("Part two:", part2(Input))
	
}
