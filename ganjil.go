package main

import "fmt"

func main() {
	var input uint
	var output bool
	
	fmt.Scan(&input)
	
	if (input % 2 == 0) {
		output = true
	} else {
		output = false
	}
	
	fmt.Println(output)
}