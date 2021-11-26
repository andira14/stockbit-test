package main

import (
	"fmt"
	"strings"
)

// A function to sum the ascii value
func asciiSum(charSet []byte) int {
	var total byte
	for _, element := range charSet {
		total += element
	}
	return int(total)
}

func anagram(stringArray []string) (result [][]string) {
	mapper := make(map[string][]string)

	for _, element := range stringArray {
		// To sanitize the ascii value
		charSets := []byte(strings.ToLower(element))

		// hash map naming convention
		mapping := fmt.Sprintf("%d,%d", asciiSum(charSets), len(element))

		mapper[mapping] = append(mapper[mapping], element)
	}

	for _, element := range mapper {
		result = append(result, element)
	}
	return result
}

func main() {
	fmt.Println(anagram([]string{"kita", "atik", "tika", "aku", "kua", "makan", "kia"}))
}
