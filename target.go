package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
)

type target struct {
	folder string
	exts   []string
	hashes []string // existing file hashes
}

func (t *target) loadHashes() error {
	fs, err := ioutil.ReadDir(t.folder)
	if err != nil {
		return fmt.Errorf("loadHashes(): %w", err)
	}
	for _, fi := range fs {
		if fi.IsDir() {
			continue
		}
		h, ok := hashFromName(fi.Name())
		if !ok {
			continue
		}
		t.hashes = append(t.hashes, h)
	}
	return nil
}

func (t *target) hashExists(hash string) bool {
	for _, h := range t.hashes {
		if h == hash {
			return true
		}
	}
	return false
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
