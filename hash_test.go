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
