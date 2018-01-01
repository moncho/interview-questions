package minesweeper

import (
	"reflect"
	"testing"
)

func Test_mineSweeper(t *testing.T) {
	type args struct {
		bombs   [][]int
		rows    int
		columns int
	}
	tests := []struct {
		name string
		args args
		want gameMap
	}{
		{
			"2 bombs 3x3 game",
			args{
				[][]int{{0, 2}, {2, 0}},
				3,
				3,
			},
			[][]int{[]int{0, 1, -1}, []int{1, 2, 1}, []int{-1, 1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mineSweeper(tt.args.bombs, tt.args.rows, tt.args.columns); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mineSweeper() = %v, want %v", got, tt.want)
			}
		})
	}
}
