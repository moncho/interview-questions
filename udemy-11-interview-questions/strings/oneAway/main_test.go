package oneAway

import "testing"

func Test_isOneAway(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"abcde -> abcd",
			args{
				a: "abcde",
				b: "abcd",
			},
			true,
		},
		{
			"abde -> abcde",
			args{
				a: "abde",
				b: "abcde",
			},
			true,
		},
		{
			"abde -> abcde",
			args{
				a: "abde",
				b: "abcde",
			},
			true,
		},
		{
			"a -> a",
			args{
				a: "a",
				b: "a",
			},
			true,
		},
		{
			"abcdef -> abqdef",
			args{
				a: "abcdef",
				b: "abqdef",
			},
			true,
		},
		{
			"abcdef -> abccef",
			args{
				a: "abcdef",
				b: "abccef",
			},
			true,
		},
		{
			"abcdef -> abcce",
			args{
				a: "abcdef",
				b: "abcce",
			},
			false,
		},
		{
			"aaa -> abc",
			args{
				a: "aaa",
				b: "abc",
			},
			false,
		},
		{
			"abcde -> abc",
			args{
				a: "abcde",
				b: "abc",
			},
			false,
		},
		{
			"abc -> abcde",
			args{
				a: "abc",
				b: "abcde",
			},
			false,
		},
		{
			"abc -> bcc",
			args{
				a: "abc",
				b: "bcc",
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOneAway(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("isOneAway() = %v, want %v", got, tt.want)
			}
		})
	}
}
