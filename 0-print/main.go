package main

import (
  "fmt"
  "time"
)

func main() {
  p := fmt.Println
  now := time.Now()
  p("Hello, we are Holberton School")
  p("the date is", now)
  p("the year is", now.Year())
  p("the month is", now.Month())
  p("the day is", now.Day())
}
