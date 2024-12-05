// Assisted by watsonx Code Assistant 
package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day1/input.txt")
//	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	arr := [][2]int{}
	
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

		intnums := [2]int{num1, num2}
		arr = append(arr, intnums)

		// fmt.Println("Reading: ", num1, num2)
	}

	// fmt.Println("Array: ", arr)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// input read into arr now go through and take lowest from left and right and find difference
	// max := 99999999
	minLeft := 0
	minRight := 0
	total := 0.0

	maxLeft := 0
	maxRight := 0
	for i := 0; i < len(arr); i++ {
		arrayEntry := arr[i]
		leftNum := arrayEntry[0]
		rightNum := arrayEntry[1]
		if (leftNum > maxLeft) {
			maxLeft = leftNum
		}
		if (rightNum > maxRight) {
			maxRight = rightNum
		}
	}

	// fmt.Println("maxLeft: ", maxLeft)
	// fmt.Println("maxRight: ", maxRight)

	xpos := 0
	ypos := 0

	// Check if not changed previously so can exit loop as processed everything
	for x := 0; x < len(arr); x++  {
		oldMinLeft := minLeft
		oldMinRight := minRight

		minLeft = maxLeft
		minRight = maxRight

		// get lowest left and right numbers above the current oldMinLeft and oldMinRight
		for i := 0; i < len(arr); i++ {
			arrayEntry := arr[i]

			leftNum := arrayEntry[0]
			rightNum := arrayEntry[1]

			// fmt.Println("leftNum: ", leftNum)
			// fmt.Println("rightNum: ", rightNum)

			if (leftNum >= oldMinLeft) {
				if (leftNum <= minLeft) {
					minLeft = leftNum
					xpos = i
				}
			}
			if (rightNum >= oldMinRight) {
				if (rightNum <= minRight) {
					minRight = rightNum
					ypos = i
				}
			}
		}

		// fmt.Println("minLeft: ", minLeft)
		// fmt.Println("minRight: ", minRight)

		// find diff between min left and right
		difference := math.Abs(float64(minLeft) - float64(minRight))
		total = total + difference

		// Blank in the array as duplicates in y!
		arr[xpos][0] = 0
		arr[ypos][1] = 0

	}

	fmt.Println("TOTAL: ", total)

}