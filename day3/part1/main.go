package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input string
	for scanner.Scan() {
		input += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	pattern := `mul\(\d+,\d+\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAll([]byte(input), -1)

	sum := 0
	re = regexp.MustCompile(`\d+`)
	for _, val := range matches {
		nums := re.FindAllString(string(val), -1)
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Printf("Error converting %s to integer: ", nums[0])
		}
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Printf("Error converting %s to integer: ", nums[1])
		}

		sum += num1 * num2
	}
	fmt.Print(sum)
}
