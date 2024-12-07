package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction struct {
	X int
	Y int
}

var dirs = []Direction{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}

func main() {
	input := loadData("../input.txt")
	fmt.Print(solve(input))
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

func solve(data [][]string) int {
	ans := 0
	for x := 1; x < len(data)-1; x++ {
		for y := 1; y < len(data[0])-1; y++ {
			if data[x][y] == "A" {
				s := ""
				for _, dir := range dirs {
					s += data[x+dir.X][y+dir.Y]
				}

				if s == "MMSS" || s == "MSSM" || s == "SSMM" || s == "SMMS" {
					ans++
				}
			}
		}
	}
	return ans
}
