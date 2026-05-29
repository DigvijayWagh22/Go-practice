package main

import "strings"

func countWordFrequency(text string) map[string]int {
	result := map[string]int{}
	text = strings.ToLower(text)
	specialChars := []string{".", ",", "!", "?", ";", ":", "\"", "'"}

	for _, ch := range specialChars {
		text = strings.ReplaceAll(text, ch, "")
	}
	text = strings.ReplaceAll(text, "-", " ")
	words := strings.Fields(text)

	for _, word := range words {
		if word != "" {
			result[word]++
		}
	}
	return result
}
func main() {

}
