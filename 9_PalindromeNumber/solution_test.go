package _9

import (
	"strconv"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{
			input: 123,
			want:  false,
		},
		{
			input: 234,
			want:  false,
		},
		{
			input: 111,
			want:  true,
		},
		{
			input: 121,
			want:  true,
		},
		{
			input: 10001,
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.input), func(t *testing.T) {
			if got := isPalindrome(tt.input); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
