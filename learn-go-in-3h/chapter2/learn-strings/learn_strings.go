package main

import "fmt"

func main() {
	s := "Hello, \t\"World! \\" // Hello,  "World! \
	b := s[0]                   // 72
	b_string := string(b)       // H
	b2 := s[4]                  // 111
	b2_string := string(b2)     // o
	fmt.Println(s, b, b_string, b2, b2_string)

	var vals [3]int
	vals[0] = 2

	fmt.Println(vals)
}
