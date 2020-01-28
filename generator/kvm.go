package main

import (
	"bytes"
	"io/ioutil"
	"text/template"
	"gopkg.in/yaml.v2"
	"github.com/zufardhiyaulhaq/terraform-config-generator/models"
)

func ProviderKVM(data []byte) {
	// unmarshal data into struct
	var spec models.ProviderKVM
	var output bytes.Buffer
	yaml.Unmarshal(data, &spec)

	// Virtual Machine
	tmplFile := "/usr/local/terraform-gen/templates/kvm/provider.tf.tmpl"
	tpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	
	err = tpl.Execute(&output, spec)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("provider.tf", output.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func VirtualMachineKVM(data []byte) {

	var spec models.VirtualMachineKVM
	var output bytes.Buffer
	yaml.Unmarshal(data, &spec)

	// Virtual Machine
	tmplFile := "/usr/local/terraform-gen/templates/kvm/virtual_machine.tf.tmpl"
	tpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	
	err = tpl.Execute(&output, spec)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(spec.Spec.Name+"_virtual_machine.tf", output.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

	//Cloud Init
	output.Reset()

	tmplFile = "/usr/local/terraform-gen/templates/kvm/cloud_init.tmpl"
	tpl, err = template.ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	
	err = tpl.Execute(&output, spec)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(spec.Spec.Name+"_cloud_init.cfg", output.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

	//Network Config
	output.Reset()

	tmplFile = "/usr/local/terraform-gen/templates/kvm/network_config.tmpl"
	tpl, err = template.ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	
	err = tpl.Execute(&output, spec)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(spec.Spec.Name+"_network_config.cfg", output.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func VirtualNetworkKVM(data []byte) {
	// unmarshal data into struct
	var spec models.VirtualNetworkKVM
	var output bytes.Buffer
	yaml.Unmarshal(data, &spec)

	// Virtual Machine
	tmplFile := "/usr/local/terraform-gen/templates/kvm/network.tf.tmpl"
	tpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	
	err = tpl.Execute(&output, spec)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("network.tf", output.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}
