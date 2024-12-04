package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"
)

func validateDecrease(array[]int) bool {
  var endOfArray = len(array)
  for i := 0; i < len(array); i++ { // Dont check the one we already checked, thats why i = 1
    nextIndex := i+1
    // If the nextIndex is bigger than the arrays values we are on the last value and return true 
    if nextIndex >= endOfArray {
      return true
    }
    currentValue := array[i]
    nextValue := array[nextIndex]
    diffValue := currentValue - nextValue
    //fmt.Println("validateDecrease: ", diffValue, currentValue, nextValue)
    if (diffValue <= 0) || (diffValue > 3 ) {
      return false
    }
  }
  return true
}

func validateIncrease(array[]int) bool {
  var endOfArray = len(array)
  for i := 0; i < len(array); i++ { // Dont check the one we already checked, thats why i = 1
    nextIndex := i+1
    // If the nextIndex is bigger than the arrays values we are on the last value and return true 
    if nextIndex >= endOfArray {
      return true
    }
    currentValue := array[i]
    nextValue := array[nextIndex]
    diffValue := currentValue - nextValue
    //fmt.Println("validateIncrease: ", diffValue, currentValue, nextValue)
    if (diffValue >= 0) || (diffValue < -3) {
      return false
    }
  }
  return true
}

func toIntArray (strArray[]string) []int {
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



func main() {
  // file, err := os.Open("test_data.txt")
  file, err := os.Open("data.txt")
  
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
    lineArray := strings.Fields(v)

    firstValue, err := strconv.Atoi(lineArray[0])
    secondValue, err := strconv.Atoi(lineArray[1])

    if err != nil { 
      fmt.Println("Error converting:", firstValue, err)
      return
		}
    if err != nil {
			fmt.Println("Error converting:", secondValue, err)
			return
		}
    // Loop over the numbers in lineArray
    trajectoryValue := firstValue - secondValue
    // if the trajectoryValue is 0 its unsafe.
    if trajectoryValue == 0 {
      continue

    } else if trajectoryValue > 0 { // Its decreasing, because its bigger then 0. 20 - 10 = 10
      if validateDecrease(toIntArray(lineArray)) { // If the validation succeds increase safeReports value
        safeReports++ 
      }

    } else if trajectoryValue < 0 { // Its increasing, because its less than 0. 10 - 20 = -10
      if validateIncrease(toIntArray(lineArray)) { // If the validation succeds increase safeReports value
        safeReports++ 
      }

    } else {
      continue
    }
  }
  fmt.Println("There are ", safeReports, " safe reports")
}
