package group_anagrams

import (
	"sort"
)

func GroupAnagrams(val []string) [][]string {
	aMap := map[string][]string{}

	for _, str := range val {
		r := []rune(str)
		sort.Slice(
			r, func(i, j int) bool {
				return r[i] < r[j]
			})
		key := string(r)
		aMap[key] = append(aMap[key], str)
	}
	result := [][]string{}
	for _, v := range aMap {
		result = append(result, v)
	}
	return result
}
