package main

import "fmt"

func main() {
	const tableSize = 8

	result := ""

	for i := 0; i < tableSize; i++ {
		for j := 0; j < tableSize; j++ {
			if i%2 == 0 {
				if j%2 == 0 {
					result += " "
				} else {
					result += "#"
				}
			} else {
				if j%2 == 0 {
					result += "#"
				} else {
					result += " "
				}
			}
		}

		result += "\n"
	}

	fmt.Println(result)
}
