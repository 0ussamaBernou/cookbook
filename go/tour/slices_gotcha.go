package main

import (
	"fmt"
	"os"
	"regexp"
)

var digitRegexp = regexp.MustCompile("[0-9]+")

func CopyDigits(filename string) []byte {
	b, _ := os.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

func main() {
	c := CopyDigits("test.txt")
	fmt.Println(c)
}
