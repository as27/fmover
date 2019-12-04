package main

import (
	"io"
	"strings"
	"testing"
)

func Test_makeHash(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"simple Text md5-Hash",
			args{strings.NewReader("abcdefg")},
			"7ac66c0f148de9519b8bd264312c4d64",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := makeHash(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("makeHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addHashToName(t *testing.T) {
	type args struct {
		fname string
		hash  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"simple hash",
			args{
				"foo/bar/filename.txt",
				"7ac66c0f148de9519b8bd264312c4d64",
			},
			"filename.7ac66c0f148de9519b8bd264312c4d64.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addHashToName(tt.args.fname, tt.args.hash); got != tt.want {
				t.Errorf("addHashToName() = %v, want %v", got, tt.want)
			}
		})
	}
}
