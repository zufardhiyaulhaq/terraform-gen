package main

import (
	"os"
	"log"
	"fmt"
	"io/ioutil"
	// "yaml"
)

func checkArgs() {
	length := len(os.Args)

	if length < 2 {
		log.Fatal("Expected 1 input, given ",length-1);
        os.Exit(1);
	}

	if length > 2 {
		log.Fatal("Expected 1 input, given ",length-1);
        os.Exit(1);
	}
}

func isFileExist(filename string) {
    _, err := os.Stat(filename)
    if os.IsNotExist(err) {
        log.Fatal("File not found");
        os.Exit(1);
	}
}


func main() {
	checkArgs()

	file := os.Args[1]
	isFileExist(file)

	dat, _ := ioutil.ReadFile(file)
	dat.asdas()
	
	// var kind models.Kind
	// yaml.Unmarshal(yamlTmpl, &kind)
}