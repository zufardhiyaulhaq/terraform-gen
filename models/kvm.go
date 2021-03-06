package models

type VirtualMachineKVM struct {
	Spec struct { 
		Name string `yaml:"name"`
		Pool string `yaml:"pool"`
		ImagePath string `yaml:"imagePath"`
		Size string `yaml:"size"`
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

type VirtualNetworkKVM struct {
	Spec struct {
		Name string `yaml:"name"`
		Mode string `yaml:"mode"`
		Network string `yaml:"network"`
		AutoStart string `yaml:"autostart"`
		DHCPEnable string `yaml:"dhcpenable"`
	} `yaml:"spec"`
}