// Assisted by watsonx Code Assistant 
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")
		if len(parts) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}

		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error converting first number:", err)
			continue
		}

		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error converting second number:", err)
			continue
		}

		fmt.Println(num1, num2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}