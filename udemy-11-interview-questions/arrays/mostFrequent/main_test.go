package mostFrequent

import "testing"

func TestMostFrequent(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{
				a: []int{1, 3, 1, 3, 2, 1},
			},
			1,
		},
		{
			"",
			args{
				a: []int{3, 3, 1, 3, 2, 1},
			},
			3,
		},
		{
			"",
			args{
				a: []int{0},
			},
			0,
		},
		{
			"",
			args{
				a: []int{0, -1, 10, 10, -1, 10, -1, -1, -1, 1},
			},
			-1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MostFrequent(tt.args.a); got != tt.want {
				t.Errorf("MostFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
