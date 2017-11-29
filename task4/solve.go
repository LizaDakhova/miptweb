package main

import (
	"unicode"
	"strings"
	//"fmt"
)

func RemoveEven(input_slice []int)(out_slice [] int) {
	for _, num := range(input_slice) {
		if num % 2 == 1 {
			out_slice = append(out_slice, num)
		}
	}
	return
}

func PowerGenerator(number int) (func() int) {
	pow := 1
	return func() (result int) {
		result = 1
		for i := 0; i < pow; i++ {
			result *= number
		}
		pow++
		return
	}
}

func DifferentWordsCount(str string) int {
	diff_words := make(map[string]bool)
	word := ""
	for _, char := range(str) {
		if unicode.IsLetter(char) {
			word = word + string(char)
			word = strings.ToLower(word)
		} else {
			if word != "" {
				diff_words[word] = true
			}
			word = ""
		}
	}
	if word != "" {
		diff_words[word] = true;
	}
	return len(diff_words)
}

/*func main() {
	input := []int{0, 3, 2, 5}
	result := RemoveEven(input)
	fmt.Println(result)
	gen := PowerGenerator(2)
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(DifferentWordsCount("Hello, world!HELLO  wOrlD...12"))*/
}
