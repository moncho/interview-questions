package commonElements

import (
	"reflect"
	"testing"
)

func TestCommonElements(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"",
			args{
				a: []int{1, 3, 4, 6, 7, 9},
				b: []int{1, 2, 4, 5, 9, 10},
			},
			[]int{1, 4, 9},
		},
		{
			"",
			args{
				a: []int{1, 2, 9, 10, 11, 12},
				b: []int{0, 1, 2, 3, 4, 5, 8, 9, 10, 12, 14, 15},
			},
			[]int{1, 2, 9, 10, 12},
		},
		{
			"",
			args{
				a: []int{0, 1, 2, 3, 4, 5},
				b: []int{6, 7, 8, 9, 10, 11},
			},
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommonElements(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommonElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
