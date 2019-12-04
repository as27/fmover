package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"time"
)

var (
	flagConfFile    = flag.String("conf", "./conf.txt", "Configuration file")
	flagWatchFolder = flag.String("w", ".", "the folder to watch")
)

func main() {
	flag.Parse()
	cf, err := os.Open(*flagConfFile)
	if err != nil {
		fmt.Println("error open conf file")
		fmt.Println(err)
		os.Exit(100)
	}
	targets, err := readConf(cf)
	if err != nil {
		fmt.Println(err)
	}
	cf.Close()
	// Ctrl+c quits the programm
	// but moves the files
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	quit := make(chan struct{})
	go func() {
		<-c
		close(quit)
	}()
	filesToMove := watch(quit)
	moveToTargets(targets, filesToMove)
}

func watch(quit chan struct{}) []string {
	var filesToMove []string
	start := time.Now()
	fmt.Println(start)
	<-quit
	fmt.Println("Analysing folder files:")
	fs, err := ioutil.ReadDir(*flagWatchFolder)
	if err != nil {
		fmt.Printf("watch(): %s", err)
	}
	for _, fi := range fs {
		if fi.IsDir() {
			continue
		}
		if start.After(fi.ModTime()) {
			continue
		}
		filesToMove = append(filesToMove, fi.Name())
	}
	return filesToMove
}

func moveToTargets(targets []target, filesToMove []string) {
	for _, t := range targets {
		err := t.loadHashes()
		if err != nil {
			fmt.Printf("cannot load hashes: %s\n", err)
			continue
		}
	}
	for _, fname := range filesToMove {
		hash, err := hashFile(fname)
		if err != nil {
			fmt.Println("hashFile():", err)
			continue
		}
		fmt.Printf("%s --> %s\n", fname, hash)
		for _, t := range targets {
			if !useTarget(t, fname) {
				continue
			}
			fmt.Printf("target: %s\n", t.folder)
			if t.hashExists(hash) {
				fmt.Println("file already in target")
				err = os.Remove(fname)
				if err != nil {
					fmt.Println("cannot remove file:", err)
				}
				continue
			}
			tfname := addHashToName(fname, hash)
			tpath := filepath.Join(t.folder, tfname)
			err = mv(tpath, fname)
			if err != nil {
				fmt.Println("cannot move: ", err)
			}
		}
	}
}

func mv(dst, src string) error {
	dFile, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return fmt.Errorf("mv() dFile: %w", err)
	}
	sFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("mv() sFile: %w", err)
	}
	io.Copy(dFile, sFile)
	sFile.Close()
	dFile.Close()
	return os.Remove(src)
}
