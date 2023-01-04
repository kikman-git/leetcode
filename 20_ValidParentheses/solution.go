package validparentheses

import "log"

func isValid(s string) bool {
	log.Println("s", s)
	if len(s) <= 1 || len(s)%2 != 0 {
		return false
	}

	queue := []byte{}
	queue = append(queue, s[0])

	for i := 1; i < len(s); i++ {
		if len(queue) != 0 && getOpenBracket(s[i]) == queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
			continue
		}
		queue = append(queue, s[i])
	}

	return len(queue) == 0
}

func getOpenBracket(s byte) byte {
	switch s {
	case 41:
		return 40
	case 125:
		return 123
	case 93:
		return 91
	default:
		return 0
	}
}
