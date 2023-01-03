package _14

import (
	"fmt"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs []string
		want string
	}{
		{
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			strs: []string{"car"},
			want: "car",
		},
		{
			strs: []string{},
			want: "",
		},
		{
			strs: []string{"reflower", "flow", "flight"},
			want: "",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := longestCommonPrefix(tt.strs); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
