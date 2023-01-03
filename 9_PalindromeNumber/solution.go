package _9

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var tmpArr []int
	for x >= 1 {
		tmpArr = append(tmpArr, x%10)
		x /= 10
	}
	length := len(tmpArr)
	if length <= 1 {
		return true
	}
	i := 0
	j := length - 1
	for i < length/2 {
		if tmpArr[i] != tmpArr[j] {
			return false
		}
		i++
		j--
	}
	return true
}
