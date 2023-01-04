package validparentheses

import "testing"

func Test_isValid(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "()",
			want:  true,
		},
		{
			input: "()",
			want:  true,
		},
		{
			input: "()[]{}",
			want:  true,
		},
		{
			input: "(]",
			want:  false,
		},
		{
			input: "([])",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := isValid(tt.input); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
