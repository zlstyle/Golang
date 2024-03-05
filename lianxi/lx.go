package main

import "fmt"

var sz = [5]string{"I", "am", "stupid", "and", "weak"}

func main() {
	for index, value := range sz {
		if value == "weak" {
			sz[index] = "strong"
		}
	}
	fmt.Println(sz)
}
