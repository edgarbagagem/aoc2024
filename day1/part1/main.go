package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

	var list1, list2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "  ")

		if len(parts) != 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		val1, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		val2, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err1 != nil || err2 != nil {
			fmt.Println("Error parsing line:", line)
			continue
		}

		list1 = append(list1, val1)
		list2 = append(list2, val2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sort.Ints(list1)
	sort.Ints(list2)

	sum := 0

	for i := 0; i < len(list1); i++ {
		sum += int(math.Abs(float64(list1[i] - list2[i])))
	}

	fmt.Printf("%d", sum)
}
