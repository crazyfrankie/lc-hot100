package main

import (
	"fmt"
	"strings"
)

func defangIPaddr(address string) string {
	var res strings.Builder
	start := 0
	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			res.WriteString(address[start:i] + "[.]")
			start = i + 1
		}
	}

	res.WriteString(address[start:])

	return res.String()
}

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
}
