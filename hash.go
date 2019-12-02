package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func makeHash(r io.Reader) (string, error) {
	h := md5.New()
	_, err := io.Copy(h, r)
	return fmt.Sprintf("%x", h.Sum(nil)), err
}
