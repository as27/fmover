package main

import (
	"path/filepath"
	"regexp"
	"strings"
)

type target struct {
	folder string
	exts   []string
}

func useTarget(t target, fpath string) bool {
	ext := filepath.Ext(fpath)
	for _, e := range t.exts {
		if e == ext || e == ".*" {
			return true
		}
	}
	return false
}

var validHash = regexp.MustCompile("[a-fA-F0-9]{32}")

func hashFromName(name string) (hash string, ok bool) {
	e := strings.Split(name, ".")
	hash = e[len(e)-2]
	ok = true
	if !validHash.MatchString(hash) {
		ok = false
		hash = ""
	}
	return hash, ok
}
