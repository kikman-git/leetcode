package _14

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	// get shortest string
	shortestStr := strs[0]
	shortestIndex := 0
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(shortestStr) {
			shortestStr = strs[i]
			shortestIndex = i
		}
	}

	// check if shortestString is a common prefix
	for len(shortestStr) > 0 {
		isCommonPrefix := true
		for i, str := range strs {
			if i == shortestIndex {
				continue
			}
			if !strings.HasPrefix(str, shortestStr) {
				isCommonPrefix = false
				break
			}
		}

		if isCommonPrefix {
			return shortestStr
		}
		
		shortestStr = shortestStr[:len(shortestStr)-1]
	}

	return shortestStr
}
