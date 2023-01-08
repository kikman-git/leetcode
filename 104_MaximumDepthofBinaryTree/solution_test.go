package maximumdepthofbinarytree

import (
	"fmt"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		// TODO: Add test cases.
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := maxDepth(tt.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
