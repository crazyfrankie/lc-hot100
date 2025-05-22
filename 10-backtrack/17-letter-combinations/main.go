package main

import "fmt"

func letterCombinations(digits string) []string {
	m := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	if digits == "" {
		return []string{""}
	}

	var result []string
	var path []byte
	var backtrack func(start int)
	backtrack = func(start int) {
		if len(path) == len(digits) {
			temp := make([]byte, len(path))
			copy(temp, path)
			result = append(result, string(temp))
			return
		}
		for i := start; i < len(digits); i++ {
			als := m[digits[i]]
			for j := 0; j < len(als); j++ {
				path = append(path, als[j])
				backtrack(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(0)

	return result
}

func main() {
	fmt.Println(letterCombinations("23"))
}
