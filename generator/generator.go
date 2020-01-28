package main

import (
	"os"
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/zufardhiyaulhaq/terraform-config-generator/models"
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

	yamlTmpl, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	  }
	
	var k models.Kind
	yaml.Unmarshal(yamlTmpl, &k)
	
	switch k.Kind {
		case "VirtualMachine.KVM":
			VirtualMachineKVM(yamlTmpl)
		case "Provider.KVM":
			ProviderKVM(yamlTmpl)
	}
}