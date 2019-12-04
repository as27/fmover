package main

import (
	"testing"
)

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
				target{folder: "abcd/foo", exts: []string{".jpg", ".png"}},
				"foo/bar/abc.png",
			},
			true,
		},
		{
			"wildcard",
			args{
				target{folder: "abcd/foo", exts: []string{".*", ".png"}},
				"foo/bar/abc.mp4",
			},
			true,
		},
		{
			"no match",
			args{
				target{folder: "abcd/foo", exts: []string{".jpg", ".png"}},
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

func Test_hashFromName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name     string
		args     args
		wantHash string
		wantOk   bool
	}{
		{
			"simple extract",
			args{"filename.7ac66c0f148de9519b8bd264312c4d64.jpg"},
			"7ac66c0f148de9519b8bd264312c4d64",
			true,
		},
		{
			"no valid hash",
			args{"filename.7zc66c0f148de9519b8bd264312c4d64.jpg"},
			"",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHash, gotOk := hashFromName(tt.args.name)
			if gotHash != tt.wantHash {
				t.Errorf("hashFromName() gotHash = %v, want %v", gotHash, tt.wantHash)
			}
			if gotOk != tt.wantOk {
				t.Errorf("hashFromName() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
