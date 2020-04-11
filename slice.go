package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	i := 0
	s := make([]int64, 3)
	for {
		var str string
		fmt.Println("enter a number or x to exit")
		_, err := fmt.Scanln(&str)
		//user input
		if err == nil {
			isint := isInt(str)
			if isint { //checking if integer or not
				var num int64
				num, err1 := strconv.ParseInt(str, 10, 64) //converting string to integer
				if err1 == nil {
					if i < 3 {
						s[i] = num
						for temp := 0; temp <= i; temp++ {
							fmt.Printf("%d ", s[temp])
						}
					} else {
						s = append(s, num)
						for temp := 0; temp <= i; temp++ {
							fmt.Printf("%d ", s[temp])
						}
					}
					i++
					fmt.Println("")
				}
			} else {
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
