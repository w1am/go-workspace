package main

import (
	"bytes"
	"fmt"
	"sort"
)

func isAnagram(s string, t string) bool {
	sb, tb := []byte(s), []byte(t)

	sort.Slice(sb, func(i int, j int) bool { return sb[i] < sb[j] })
	sort.Slice(tb, func(i int, j int) bool { return tb[i] < tb[j] })

	return bytes.Equal(sb, tb)
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}
