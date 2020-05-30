package main

import "strings"

func toLetters(num int) string {
	power := num / 26
	mod := num % 26

	var output string

	if mod > 0 {
		output = string('A' - 1 + mod)
	} else {
		power--
		output = string('Z')
	}

	var returnValue string

	if power > 0 {
		returnValue = toLetters(power) + output
	} else {
		returnValue = output
	}

	return strings.ToLower(returnValue)
}