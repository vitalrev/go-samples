package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("(..?)")
	codes := re.FindAllString("The rain in SPAIN stays mainly in the plain", -1)
	fmt.Printf("%v", codes)
}
