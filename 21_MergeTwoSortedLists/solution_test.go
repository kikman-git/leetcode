package mergetwosortedlists

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		args args
		want *ListNode
	}{
		{
			args: args{
				list1: convertToList([]int{1, 2, 4}),
				list2: convertToList([]int{1, 3, 4}),
			},
			want: convertToList([]int{1, 1, 2, 3, 4, 4}),
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
