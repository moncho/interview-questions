package click

import (
	"reflect"
	"testing"
)

func Test_click(t *testing.T) {
	type args struct {
		game gameMap
		i    int
		j    int
	}
	tests := []struct {
		name string
		args args
		want gameMap
	}{
		{
			"3x5 click 2 2",
			args{
				[][]int{{0, 0, 0, 0, 0},
					{0, 1, 1, 1, 0},
					{0, 1, -1, 1, 0}},
				2,
				2,
			},
			[][]int{{0, 0, 0, 0, 0},
				{0, 1, 1, 1, 0},
				{0, 1, -1, 1, 0}},
		},
		{
			"3x5 click 1 4",
			args{
				[][]int{{0, 0, 0, 0, 0},
					{0, 1, 1, 1, 0},
					{0, 1, -1, 1, 0}},
				1,
				4,
			},
			[][]int{{-2, -2, -2, -2, -2},
				{-2, 1, 1, 1, -2},
				{-2, 1, -1, 1, -2}},
		},
		{
			"4x4 click 0 1",
			args{
				[][]int{{-1, 1, 0, 0},
					{1, 1, 0, 0},
					{0, 0, 1, 1},
					{0, 0, 1, -1}},
				0,
				1,
			},
			[][]int{{-1, 1, 0, 0},
				{1, 1, 0, 0},
				{0, 0, 1, 1},
				{0, 0, 1, -1}},
		},
		{
			"4x4 click 1 3",
			args{
				[][]int{{-1, 1, 0, 0},
					{1, 1, 0, 0},
					{0, 0, 1, 1},
					{0, 0, 1, -1}},
				1,
				3,
			},
			[][]int{{-1, 1, -2, -2},
				{1, 1, -2, -2},
				{-2, -2, 1, 1},
				{-2, -2, 1, -1}},
		},
		{
			"3x6 click 1 2",
			args{
				[][]int{{0, 0, 1, 1, 1, 0},
					{0, 0, 1, -1, 1, 0},
					{0, 0, 1, 1, 1, 0}},
				1,
				1,
			},
			[][]int{{-2, -2, 1, 1, 1, 0},
				{-2, -2, 1, -1, 1, 0},
				{-2, -2, 1, 1, 1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := click(tt.args.game, tt.args.i, tt.args.j); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("click() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_clickBFS(t *testing.T) {
	type args struct {
		game gameMap
		i    int
		j    int
	}
	tests := []struct {
		name string
		args args
		want gameMap
	}{
		{
			"3x5 click 2 2",
			args{
				[][]int{{0, 0, 0, 0, 0},
					{0, 1, 1, 1, 0},
					{0, 1, -1, 1, 0}},
				2,
				2,
			},
			[][]int{{0, 0, 0, 0, 0},
				{0, 1, 1, 1, 0},
				{0, 1, -1, 1, 0}},
		},
		{
			"3x5 click 1 4",
			args{
				[][]int{{0, 0, 0, 0, 0},
					{0, 1, 1, 1, 0},
					{0, 1, -1, 1, 0}},
				1,
				4,
			},
			[][]int{{-2, -2, -2, -2, -2},
				{-2, 1, 1, 1, -2},
				{-2, 1, -1, 1, -2}},
		},
		{
			"4x4 click 0 1",
			args{
				[][]int{{-1, 1, 0, 0},
					{1, 1, 0, 0},
					{0, 0, 1, 1},
					{0, 0, 1, -1}},
				0,
				1,
			},
			[][]int{{-1, 1, 0, 0},
				{1, 1, 0, 0},
				{0, 0, 1, 1},
				{0, 0, 1, -1}},
		},
		{
			"4x4 click 1 3",
			args{
				[][]int{{-1, 1, 0, 0},
					{1, 1, 0, 0},
					{0, 0, 1, 1},
					{0, 0, 1, -1}},
				1,
				3,
			},
			[][]int{{-1, 1, -2, -2},
				{1, 1, -2, -2},
				{-2, -2, 1, 1},
				{-2, -2, 1, -1}},
		},
		{
			"3x6 click 1 2",
			args{
				[][]int{{0, 0, 1, 1, 1, 0},
					{0, 0, 1, -1, 1, 0},
					{0, 0, 1, 1, 1, 0}},
				1,
				1,
			},
			[][]int{{-2, -2, 1, 1, 1, 0},
				{-2, -2, 1, -1, 1, 0},
				{-2, -2, 1, 1, 1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clickBFS(tt.args.game, tt.args.i, tt.args.j); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("click() = %v, want %v", got, tt.want)
			}
		})
	}
}
