package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func stringToInt(str string) int {
	outputInt, err := strconv.Atoi(str) // Convert string to int
	if err != nil {
		fmt.Println("Error converting:", str, err)
		os.Exit(1)
	}
	return outputInt
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close() // Ensure the file is closed properly

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanRunes)

	var mulString string
	var finalResult []int

	for fileScanner.Scan() {
		c := fileScanner.Text()
		if c == "\n" {
			mulString += " "
		} else {
			mulString += c
		}
	}

	// Check for scanning errors
	if err := fileScanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Regex expressions
	re1 := regexp.MustCompile("mul\\([1-9][0-9]{0,2},[1-9][0-9]{0,2}\\)") // Match full pattern
	matches := re1.FindAllString(mulString, -1)

	fmt.Println("Matches:", matches)

	if len(matches) == 0 {
		fmt.Println("No matches found.")
		os.Exit(0)
	}

	re2 := regexp.MustCompile("[1-9][0-9]{0,2},[1-9][0-9]{0,2}") // Extract numbers within the brackets
	for _, match := range matches {
		// Extract the numbers part
		numberString := re2.FindString(match)
		if numberString == "" {
			fmt.Println("Error: No numbers found in match:", match)
			continue
		}

		// Split the numbers
		valuesToMultiply := strings.Split(numberString, ",")
		if len(valuesToMultiply) != 2 {
			fmt.Println("Error: Unexpected format for match:", match)
			continue
		}

		// Multiply the numbers and append to results
		finalResult = append(finalResult, stringToInt(valuesToMultiply[0])*stringToInt(valuesToMultiply[1]))
	}

	// Calculate the total sum
	var total int
	for _, result := range finalResult {
		total += result
	}
	fmt.Println("Total:", total)
}
