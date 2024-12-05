package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toIntArray(strArray []string) []int {
	// Output: int array
	intArray := make([]int, len(strArray))

	// Convert each string to an integer
	for i, str := range strArray {
		num, err := strconv.Atoi(str) // Convert string to int
		if err != nil {
			fmt.Println("Error converting:", str, err)
			os.Exit(1)
		}
		intArray[i] = num
	}
	return intArray
}

// If the value pair is a valid sequence return true if its unsafe return false
func validDiff(diff int) bool {
	if diff == 0 {
		return false
	} else if diff > 3 {
		return false
	} else if diff < -3 {
		return false
	}
	return true
}

func validateRow(array []int) bool {
	dampend := false

	var increase bool
	var decrease bool
	for i := 0; i < len(array)-1; i++ {
		diff := array[i] - array[i+1]
		valid := validDiff(diff)
		if (valid == false) && (dampend == true) {
			return false
		} else if (valid == false) && (dampend == false) {
			dampend = true
			// Because we are already looping over the array the amount of loops is already set so we increase i with 1
			// So we set the unsafe value to the value we want to try in the next loop
			array[i+1] = array[i]
		} else if (increase == true) && (decrease == true) {
			return false
		} else if diff > 0 { // Its decreasing because 20 - 10 = 10 (positive number)
			decrease = true
		} else if diff < 0 {
			increase = true
		}
	}
	return true
}

func main() {
	// file, err := os.Open("test_data.txt")
	file, err := os.Open("../data.txt")

	if err != nil {
		fmt.Println(err)
		panic(1)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	// Create array that will contain our data
	var fileLines []string

	// Read the data into our array
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	var safeReports int

	for _, v := range fileLines {
		if validateRow(toIntArray(strings.Fields(v))) {
			safeReports++
		}
	}
	fmt.Println("There are ", safeReports, " safe reports")
}
