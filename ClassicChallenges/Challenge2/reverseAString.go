package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		output := ReverseString(input)
		fmt.Println(output)
	}
}

func ReverseString(s string) string {
	runeArray := []rune(s)
	n := len(runeArray)
	low := 0
	high := n - 1

	for low < high {
		runeArray[low], runeArray[high] = runeArray[high], runeArray[low]
		low++
		high--
	}

	result := string(runeArray)
	return result
}
