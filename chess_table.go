package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	tableSize, error := strconv.ParseUint(os.Args[1], 10, 64)

	if error != nil {
		fmt.Println(error)
		return
	}

	result := ""

	var i uint64
	for i = 0; i < tableSize; i++ {

		var j uint64
		for j = 0; j < tableSize; j++ {
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
