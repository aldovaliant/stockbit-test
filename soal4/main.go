package main

import "fmt"

func main() {
	listStr := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	fmt.Println(anagramListCheck(listStr))
}

func anagramListCheck(listStr []string) [][]string {
	res := [][]string{}
	check := make([]bool, len(listStr))
	for i := 0; i < len(listStr); i++ {
		if check[i] {
			continue
		}
		anagramGroup := []string{listStr[i]}
		for j := i + 1; j < len(listStr); j++ {
			if check[j] {
				continue
			}
			anagram := anagramCheck(listStr[i], listStr[j])
			if anagram {
				check[i] = true
				check[j] = true
				anagramGroup = append(anagramGroup, listStr[j])
			}
		}
		res = append(res, anagramGroup)
	}
	return res
}

func anagramCheck(str1 string, str2 string) bool {
	aCharMap := buildCharMap(str1)
	bCharMap := buildCharMap(str2)
	if len(aCharMap) != len(bCharMap) {
		return false
	}
	for key := range aCharMap {
		if aCharMap[key] != bCharMap[key] {
			return false
		}
	}
	return true
}

func buildCharMap(str string) map[rune]int {
	charMap := map[rune]int{}
	for _, char := range str {
		charMap[char] = charMap[char] + 1 | 1
	}
	return charMap
}
