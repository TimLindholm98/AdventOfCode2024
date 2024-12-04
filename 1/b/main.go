package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"
)

func main(){
	file, err := os.Open("../data.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var firstColumn[]int
	var similarityScore int
	occurrenceCount := make(map[int]int)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	// Read the data into our arrays
	for fileScanner.Scan() {
		lineArray := strings.Fields(fileScanner.Text()) // Array with both columns
		// Error handling
		v0, err := strconv.Atoi(lineArray[0])
		if err != nil {
			return
		}
		v1, err :=  strconv.Atoi(lineArray[1])
		if err != nil {
			return
		}
		// Populate our array for calculating the similarity score.
		firstColumn = append(firstColumn, v0)
		// Create map for the seconds column that contains how many times a number has ocured.
		occurrenceCount[v1] = occurrenceCount[v1]+1
	}
	fmt.Println(occurrenceCount)
	for i := 0; i < len(firstColumn); i++ {
		similarityScore += firstColumn[i] * occurrenceCount[i]
	}
	fmt.Println("Similarity Score: ", similarityScore)
}