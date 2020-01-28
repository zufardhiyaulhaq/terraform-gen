package models

type VirtualMachineKVM struct {
	Spec struct { 
		Name string `yaml:"name"`
		Pool string `yaml:"pool"`
		ImagePath string `yaml:"imagePath"`
		User struct {
			Password string `yaml:"password"`
		} `yaml:"user"`
		Spec struct {
			RAM int `yaml:"ram"`
			VCPU int `yaml:"vcpu"`
		} `yaml:"spec"`
		Network struct {
			Name string `yaml:"name"`
			Interface string `yaml:"interface"`
			Addresses string `yaml:"addresses"`
			Gateway string `yaml:"gateway"`
			DNS string `yaml:"dns"`
		} `yaml:"network"`
	} `yaml:"spec"`
}

type ProviderKVM struct {
	Spec struct { 
		URL string `yaml:"url"`
	} `yaml:"spec"`
}