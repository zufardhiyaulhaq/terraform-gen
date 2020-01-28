# KVM Usage
Make sure you have following [setup tutorial](https://github.com/zufardhiyaulhaq/terraform-gen/blob/master/docs/kvm/setup.md) before using this.

---
After KVM & terraform is installed, you can generate terraform configuration with this tools.

### Installing
- Clone Repository
```
git clone https://github.com/zufardhiyaulhaq/terraform-gen.git
```
- Build & Install
```
cd terraform-gen
sudo make install
```
### Usage
- Create folder for storing the configuration
```
mkdir config && cd config
terraform init
```
- Create your configuration following the [example](https://github.com/zufardhiyaulhaq/terraform-gen/tree/master/docs/kvm/example)
- Run the tools
```
terraform-gen provider-example.yaml
terraform-gen vm-example.yaml
```
- Validate & Apply terraform
```
terraform validate
terraform apply
```

