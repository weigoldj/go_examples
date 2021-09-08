package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	dirPtr *string
)

func init() {
	dirPtr = flag.String("dir", ".", "directory to scan")
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func iterate(path string) {
	fmt.Printf("dir ", path)
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Printf("File Name: %s, ext: %s \n", info.Name(), filepath.Ext(info.Name()))
		return nil
	})
}

func toString() {
	fmt.Println("flag_me values: ")
	fmt.Println("dir:", *dirPtr)
}

func main() {
	flag.Parse()

	currentDirectory := *dirPtr
	value, _ := exists(currentDirectory)
	if value {
		iterate(currentDirectory)
		toString()
	} else {
		fmt.Printf("Error: directory does not exist! ", dirPtr)
	}

}
