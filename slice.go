package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	numbers := make([]int, 0, 3)
	var str string
	for {
		fmt.Println("Enter an integer into the slice ( x to quit ): ")
		fmt.Scanf("%s\n", &str)
		inputNum, err := strconv.Atoi(str)
		if str == "x" || err != nil {
			fmt.Println("Exiting ...")
			break
		}
		numbers = append(numbers, inputNum)
	}
	sort.Ints(numbers)
	fmt.Println(numbers)
}
