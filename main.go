package main

import (
  "github.com/fatih/color"
  "rsc.io/quote"
)

func main() {
  color.Cyan("Testing Go Modules")
  color.Cyan(quote.Hello())
}
