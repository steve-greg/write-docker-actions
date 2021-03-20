package main

import (
	"fmt"
	"os"
)

func main() {

  // Access Inputs as environment vars
  firstGreeting := os.Getenv("INPUT_FIRSTGREETING")
  secondGreeting := os.Getenv("INPUT_SECONDGREETING")
  thirdGreeting := os.Getenv("INPUT_THIRDGREETING")

  fmt.Println("Hello " + firstGreeting)
  fmt.Println("Hello " + secondGreeting)

  // Someimes inputs are not "required"
  if thirdGreeting != "" {
    fmt.Println("Hello " + thirdGreeting)
    }

}