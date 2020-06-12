package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = []string{"john", "jane", "doe"}
	var dataContainsJ = filter(data, func(each string) bool {
		return strings.Contains(each, "j")
	})

	fmt.Println("Daftar nama: ", data)
	fmt.Println("Nama dengan huruf j: ", dataContainsJ)
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, value := range data {
		if validateFilter := callback(value); validateFilter {
			result = append(result, value)
		}
	}

	return result
}
