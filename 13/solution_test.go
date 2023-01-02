package _13

import "testing"

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "II",
			want:  2,
		},
		{
			input: "XXVII",
			want:  27,
		},
		{
			input: "LVIII",
			want:  58,
		},
		{
			input: "MCMXCIV",
			want:  1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := romanToInt(tt.input); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
