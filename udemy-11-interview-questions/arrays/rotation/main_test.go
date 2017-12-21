package rotation

import "testing"

func TestIsRotation(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"",
			args{
				a: []int{1, 2, 3, 4, 5, 6, 7},
				b: []int{4, 5, 6, 7, 8, 1, 2, 3},
			},
			false,
		},
		{
			"",
			args{
				a: []int{1, 2, 3, 4, 5, 6, 7},
				b: []int{4, 5, 6, 7, 1, 2, 3},
			},
			true,
		},
		{
			"",
			args{
				a: []int{1, 2, 3, 4, 5, 6, 7},
				b: []int{4, 5, 6, 9, 1, 2, 3},
			},
			false,
		},
		{
			"",
			args{
				a: []int{1, 2, 3, 4, 5, 6, 7},
				b: []int{4, 6, 5, 7, 1, 2, 3},
			},
			false,
		},
		{
			"",
			args{
				a: []int{1, 2, 3, 4, 5, 6, 7},
				b: []int{4, 5, 6, 7, 0, 2, 3},
			},
			false,
		},
		{
			"",
			args{
				a: []int{1, 2, 3, 4, 5, 6, 7},
				b: []int{1, 2, 3, 4, 5, 6, 7},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRotation(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("IsRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
