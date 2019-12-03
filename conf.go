package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func readConf(r io.Reader) (targets []target, err error) {
	var errString string
	s := bufio.NewScanner(r)
	lnr := 0
	for s.Scan() {
		l := s.Text()
		lnr++
		// Check for a comment line
		if strings.HasPrefix(l, "#") {
			continue
		}
		e := strings.Split(l, "::")
		if len(e) != 2 {
			errString = fmt.Sprintf("%02d:%s%s\n", lnr, errString, l)
			continue
		}
		fpath := e[0]
		exts := e[1]
		t := target{fpath, splitExtString(exts)}
		targets = append(targets, t)
	}
	if errString != "" {
		err = fmt.Errorf("Wrong configuration lines:\n%s", errString)
	}
	return targets, err
}

func splitExtString(s string) []string {
	exts := strings.Split(s, ",")
	for i := range exts {
		exts[i] = fmt.Sprintf(".%s", exts[i])
	}
	return exts
}
