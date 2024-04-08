package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type node struct {
	name    string
	size    int
	entries []*node
}

func main() {

	dir := &node{name: ".", size: 2, entries: nil}
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		if file.IsDir() {
			dir.addFile(file.Name(), file.Size())
		} else {
			dir.addFile(file.Name(), file.Size())
		}
	}
}

func (n *node) addFile() {

}

func (n *node) addDir() {

}
