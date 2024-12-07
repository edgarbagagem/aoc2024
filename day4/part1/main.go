package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := loadData("../input.txt")
	fmt.Print(parseData(input))
}

func loadData(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		panic("Error opening!")
	}
	data := [][]string{}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		for _, v := range line {
			row = append(row, string(v))
		}
		data = append(data, row)
	}
	return data
}

func parseData(data [][]string) int {
	ans := 0
	for x := range len(data) {
		for y := range len(data[0]) {
			if data[x][y] == "X" {
				result := dfs(data, 0, "XMAS", x, y)
				ans += result
			}
		}
	}
	return ans
}

var dir = []string{
	"U", "UR", "R", "DR", "D", "DL", "L", "UL",
}

func dfs(data [][]string, currIndex int, word string, x int, y int) int {
	ans := 0
	for _, dir := range dir {
		result := dfsHelper(data, currIndex, word, dir, x, y)
		ans += result
	}
	return ans
}

func dfsHelper(data [][]string, currIndex int, word string, dir string, x int, y int) int {
	if currIndex >= len(word) {
		return 1
	}

	if x < 0 || x >= len(data) || y < 0 || y >= len(data[0]) {
		return 0
	}

	if data[x][y] != string(word[currIndex]) {
		return 0
	}

	switch dir {
	case "U":
		return dfsHelper(data, currIndex+1, word, dir, x-1, y)
	case "UR":
		return dfsHelper(data, currIndex+1, word, dir, x-1, y+1)
	case "R":
		return dfsHelper(data, currIndex+1, word, dir, x, y+1)
	case "DR":
		return dfsHelper(data, currIndex+1, word, dir, x+1, y+1)
	case "D":
		return dfsHelper(data, currIndex+1, word, dir, x+1, y)
	case "DL":
		return dfsHelper(data, currIndex+1, word, dir, x+1, y-1)
	case "L":
		return dfsHelper(data, currIndex+1, word, dir, x, y-1)
	default:
		return dfsHelper(data, currIndex+1, word, dir, x-1, y-1)
	}
}
