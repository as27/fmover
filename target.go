package main

import "path/filepath"

type target struct {
	folder string
	exts   []string
}

func useTarget(t target, fpath string) bool {
	ext := filepath.Ext(fpath)
	for _, e := range t.exts {
		if e == ext || e == "*" {
			return true
		}
	}
	return false
}
