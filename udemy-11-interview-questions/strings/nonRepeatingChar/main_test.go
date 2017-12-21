package nonRepeatingChar

import "testing"

func TestNonRepeatingChar(t *testing.T) {
	type args struct {
		chars string
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{
			"abcab",
			args{
				"abcab",
			},
			'c',
		},
		{
			"abab",
			args{
				"abab",
			},
			0,
		},
		{
			"aabbbc",
			args{
				"aabbbc",
			},
			'c',
		},
		{
			"aabbdbc",
			args{
				"aabbdbc",
			},
			'd',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NonRepeatingChar(tt.args.chars); got != tt.want {
				t.Errorf("NonRepeatingChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
