package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "slices"
  "strconv"
)


func main(){
  file, err := os.Open("../data.txt")

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  var firstColumn[]int
  var secondColumn[]int

  fileScanner := bufio.NewScanner(file)
  fileScanner.Split(bufio.ScanLines)
  // Read the data into our arrays
  for fileScanner.Scan() {
    lineArray := strings.Fields(fileScanner.Text())
    
    v0, err := strconv.Atoi(lineArray[0])
    if err != nil {
      return
    }
    v1, err :=  strconv.Atoi(lineArray[1])
    if err != nil {
      return
    }

    firstColumn = append(firstColumn, v0)
    secondColumn = append(secondColumn, v1)
  }
  slices.Sort(firstColumn)
  slices.Sort(secondColumn)

  var result int

  if len(firstColumn) == len(secondColumn) {
    for i := 0; i < len(firstColumn); i++ {
      if firstColumn[i] > secondColumn[i] {
        result += firstColumn[i] - secondColumn[i] 
      } else {
        result += secondColumn[i] - firstColumn[i]
      }
    }
  } else {
    fmt.Println("The lists is not the same lenght. firstColumn is", len(firstColumn), "and secondColumn is ", secondColumn)
    os.Exit(1)
  }
  fmt.Println("The total difference is: ", result)
}
