package main

import "testing"

func Test_useTarget(t *testing.T) {
	type args struct {
		t     target
		fpath string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"matching extension",
			args{
				target{"abcd/foo", []string{".jpg", ".png"}},
				"foo/bar/abc.png",
			},
			true,
		},
		{
			"wildcard",
			args{
				target{"abcd/foo", []string{"*", ".png"}},
				"foo/bar/abc.mp4",
			},
			true,
		},
		{
			"no match",
			args{
				target{"abcd/foo", []string{".jpg", ".png"}},
				"foo/bar/abc.mp4",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := useTarget(tt.args.t, tt.args.fpath); got != tt.want {
				t.Errorf("useTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
