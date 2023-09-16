package main

import (
	"regexp"
)

func main() {

	check_regexp()
}

func check() {
	var arr1 = []bool{
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true,
	}

	var result = CountSheeps(arr1)
	println(result)
}

func CountSheeps(numbers []bool) int {
	var result = 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] {
			result++
		}
	}
	return result
}

func check_regexp() {
	println(ReplaceDots("one.two.three"))
}

func ReplaceDots(str string) string {
	return regexp.MustCompile("\\.").ReplaceAllString(str, "-")
}
