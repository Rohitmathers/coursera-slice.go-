package main

import (
	"fmt"
	"sort"
	"strconv"
	"unicode"
)

func main() {
	s := make([]int, 3)
	s1 := make([]int, 0)
	for {
		var str string
		fmt.Println("enter a number or x to exit")
		_, err := fmt.Scanln(&str)
		//user input
		if err == nil {
			isint := isInt(str)
			if isint { //checking if integer or not
				var num int
				num, err1 := strconv.Atoi(str) //converting string to integer
				if err1 == nil {
					s1 = append(s1, num)
					fmt.Println("")
					sort.Ints(s1)
				}
				fmt.Println(s1)
				s = s1
			} else {
				fmt.Println(s)
				if str == "x" { //checking if user wants to exit or not
					break
				} else {
					fmt.Println("Invalid Input")
					continue
				}

			}
		}
	}
}

//Function to check whether the string is numeric or not.
func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
package main

import (
	"fmt"
	"sort"
	"strconv"
	"unicode"
)

func main() {
	s := make([]int, 3)
	s1 := make([]int, 0)
	for {
		var str string
		fmt.Println("enter a number or x to exit")
		_, err := fmt.Scanln(&str)
		//user input
		if err == nil {
			isint := isInt(str)
			if isint { //checking if integer or not
				var num int
				num, err1 := strconv.Atoi(str) //converting string to integer
				if err1 == nil {
					s1 = append(s1, num)
					fmt.Println("")
					sort.Ints(s1)
				}
				fmt.Println(s1)
				s = s1
			} else {
				fmt.Println(s)
				if str == "x" { //checking if user wants to exit or not
					break
				} else {
					fmt.Println("Invalid Input")
					continue
				}

			}
		}
	}
}

//Function to check whether the string is numeric or not.
func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
