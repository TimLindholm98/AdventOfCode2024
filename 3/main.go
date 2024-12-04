package main

import (
  "bufio"
  "fmt"
  "os"
  "regexp"
  "strings"
  "strconv"
)

func stringToInt(str string) int {
  outputInt, err := strconv.Atoi(str) // Convert string to int
  if err != nil {
    fmt.Println("Error converting:", str, err)
    os.Exit(1)
  }
  return outputInt
}

func main()  {
  file, err := os.Open("data.txt")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  fileScanner := bufio.NewScanner(file)
  fileScanner.Split(bufio.ScanRunes)
  
  var mulString string
  var finalResult int

  for fileScanner.Scan() {
    c := fileScanner.Text()
    if c != "\n" {
      mulString += c
    }
  }
  // fmt.Println(mulString)
  // https://stackoverflow.com/questions/37316806/golang-extract-data-with-regex
  // RegexExpressions
  re1 := regexp.MustCompile("mul\\([1-9]{1,3}\\,[1-9]{1,3}\\)")
  re2 := regexp.MustCompile("[0-9]{1,3}\\,[0-9]{1,3}")
  
  matches := re1.FindAllString( mulString, -1)

  //fmt.Println(matches)
  for i := 0 ; i < len(matches) ; i++ {
    // fmt.Println(matches[i])

    // FindAllString() returns string array with matches. 
    // Then strings.Split() returns a string array with two numbers
    var valuesToMultiply []string = strings.Split(re2.FindAllString(matches[i],-1)[0],",")
    // fmt.Println(valuesToMultiply)
    var result int = stringToInt(valuesToMultiply[0]) * stringToInt(valuesToMultiply[1])
    finalResult += result
    // fmt.Println("Values: ", valuesToMultiply[0], "x", valuesToMultiply[1], result)
  }
  fmt.Println(finalResult)
}
