package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_readConf(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []target
		wantErr bool
	}{
		{
			"simple conf",
			args{strings.NewReader("/abc/def/ghi::jpg,png\n/foo/bar::*")},
			[]target{
				{folder: "/abc/def/ghi", exts: []string{".jpg", ".png"}},
				{folder: "/foo/bar", exts: []string{".*"}},
			},
			false,
		},
		{
			"wrong conf",
			args{strings.NewReader("/abc/def/:jpg,png\n/foo/bar::*")},
			[]target{
				{folder: "/foo/bar", exts: []string{".*"}},
			},
			true,
		},
		{
			"comment",
			args{strings.NewReader("#/abc/def/ghi::jpg,png\n/foo/bar::*")},
			[]target{
				{folder: "/foo/bar", exts: []string{".*"}},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readConf(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("readConf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readConf() = %v, want %v", got, tt.want)
			}
		})
	}
}
