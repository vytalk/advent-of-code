package main

/*
The captcha requires you to review a sequence of digits (your puzzle input) and
find the sum of all digits that match the next digit in the list. The list is
circular, so the digit after the last digit is the first digit in the list.
*/
func SolveCaptchaA(code []int) int {
	sum := 0
	for i, j := 0, 1; i < len(code); i, j = i+1, j+1 {
		if j == len(code) {
			j = 0
		}
		if code[i] == code[j] {
			sum += code[i]
		}
	}
	return sum
}

/*
Now, instead of considering the next digit, it wants you to consider the digit
halfway around the circular list. That is, if your list contains 10 items, only
include a digit in your sum if the digit 10/2 = 5 steps forward matches it.
Fortunately, your list has an even number of elements.
*/
func SolveCaptchaB(code []int) int {
	sum := 0
	for i, j := 0, len(code)/2; i < len(code); i, j = i+1, j+1 {
		if j >= len(code) {
			j = 0
		}
		if code[i] == code[j] {
			sum += code[i]
		}
	}
	return sum
}
