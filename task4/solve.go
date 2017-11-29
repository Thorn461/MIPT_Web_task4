package main

import (
	"unicode"
)

func RemoveEven(array [] int) (answer [] int) {
	for i := 0; i < len(array); i++ {
		if array[i] % 2 == 1 {
			answer = append(answer, array[i])
		}
	}
	return
}

func PowerGenerator(number int) (func() int) {
	power := 1
	return func() int {
		power =  number * power
		return power
	}
}

func DifferentWordsCount (s string) int {
	allWords := make(map[string] int)
	currentWord := ""
	isWord := true
	for i := 0; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) {
			isWord = true
			currentWord += string(unicode.ToLower(rune(s[i])))
		} else if isWord {
			isWord = false
			allWords[currentWord]++
			currentWord = ""
		}
	}
	if isWord {
		allWords[currentWord]++
	}
	return len(allWords)
}
