package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1(numbers []int) int {
	result := 0
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[i-1] {
			result++
		}
	}
	
	return result
}

func part2(numbers []int) int {
	result := 0
	for i := 0; i < len(numbers); i++ {
		if len(numbers)-i == 3 {
			break
		} else if numbers[i]+numbers[i+1]+numbers[i+2] < numbers[i+1]+numbers[i+2]+numbers[i+3] {
			result++
		}
	}
	return result
}

func getMeasurements() []int {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var numbers []int
	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, x)
	}
	return numbers
}

func main() {
	numbers := getMeasurements()
	fmt.Println("Part one", part1(numbers))
	fmt.Println("Part two", part2(numbers))
}
