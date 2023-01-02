package _13

import (
	"log"
	"strings"
)

// My dumb solution...
func romanToInt(s string) int {
	s = strings.ToUpper(s)
	result := 0
	symbolMap := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	if len(s) <= 1 {
		return symbolMap[s]
	}

	i := 0
	for i < len(s) {
		if len(s)-i >= 2 {
			specialCase, ok := symbolMap[string(s[i])+string(s[i+1])]
			log.Println(string(s[i])+string(s[i+1]), ok)
			if ok {
				result += specialCase
				log.Println(result)
				i += 2
				continue
			}
		}

		if len(s)-i >= 3 && s[i] == s[i+1] && s[i] == s[i+2] {
			result += symbolMap[string(s[i])] * 3
			i += 3
			continue
		}

		result += symbolMap[string(s[i])]
		i++
	}

	return result
}
