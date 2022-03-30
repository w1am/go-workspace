package main

import (
	"fmt"
	"reflect"
)

func isAnagram(s string, t string) bool {
	// Check if length is same
	var match bool = len(s) == len(t)

	if !match {
		return false
	}

	table_s := make(map[string]int)
	table_t := make(map[string]int)

	for i := 0; i < len(s); i++ {
		var current string = string(s[i])
		_, isPresent := table_s[current]
		if isPresent {
			table_s[current] = table_s[current] + 1
		} else {
			table_s[current] = 1
		}
	}

	for i := 0; i < len(s); i++ {
		var current string = string(t[i])
		_, isPresent := table_t[current]
		if isPresent {
			table_t[current] = table_t[current] + 1
		} else {
			table_t[current] = 1
		}
	}

	eq := reflect.DeepEqual(table_s, table_t)

	return eq
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}
