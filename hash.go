package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func makeHash(r io.Reader) (string, error) {
	h := md5.New()
	_, err := io.Copy(h, r)
	return fmt.Sprintf("%x", h.Sum(nil)), err
}

func hashFile(fname string) (string, error) {
	fd, err := os.Open(fname)
	if err != nil {
		return "", fmt.Errorf("hashFile(%s):\n%w", fname, err)
	}
	defer fd.Close()
	return makeHash(fd)
}

func addHashToName(fname string, hash string) string {
	fname = filepath.Base(fname)
	ext := filepath.Ext(fname)
	return fmt.Sprintf("%s.%s%s", fname[:len(fname)-len(ext)], hash, ext)
}

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
