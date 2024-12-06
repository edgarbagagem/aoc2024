package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var input [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		var numbers []int

		for _, val := range parts {
			num, err := strconv.Atoi(val)

			if err != nil {
				fmt.Println("Error converting str to num:", err)
				continue
			}
			numbers = append(
				numbers,
				num,
			)
		}
		input = append(input, numbers)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	safeReports := 0

	for _, lineOfNumbers := range input {
		if isSafeWithDampener(lineOfNumbers) {
			safeReports++
		}
	}

	fmt.Printf("%d", safeReports)
}

func isSafe(numbers []int) bool {
	isIncreasing := numbers[1] > numbers[0]

	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]

		if diff > 3 || diff < -3 {
			return false
		}

		if isIncreasing && diff <= 0 {
			return false
		}
		if !isIncreasing && diff >= 0 {
			return false
		}
	}
	return true
}

func isSafeWithDampener(numbers []int) bool {
	if isSafe(numbers) {
		return true
	}

	for i := 0; i < len(numbers); i++ {
		newReport := append([]int{}, numbers[:i]...)
		newReport = append(newReport, numbers[i+1:]...)

		if isSafe(newReport) {
			return true
		}
	}

	return false
}
