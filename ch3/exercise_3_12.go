package ch3

import "reflect"

func IsScrambledStr(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	hashmap1 := make(map[rune]int)
	for _, k := range str1 {
		_, isHave := hashmap1[k]
		if isHave {
			hashmap1[k]++
		} else {
			hashmap1[k] = 1
		}
	}
	hashmap2 := make(map[rune]int)
	for _, k := range str1 {
		_, isHave := hashmap2[k]
		if isHave {
			hashmap2[k]++
		} else {
			hashmap2[k] = 1
		}
	}
	return reflect.DeepEqual(hashmap1, hashmap2)
}
