# Terraform Setup
Install terraform inside your KVM server.
- Install Development tools, libvirt-devel, & mkisofs
```
yum groupinstall "Development Tools"
yum install libvirt-devel mkisofs
```
- Install Golang
Follow the [official documentation](https://golang.org/doc/install) to install Golang.
```
wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.13.4.linux-amd64.tar.gz
```
```
nano /etc/profile
export PATH=$PATH:/usr/local/go/bin
```
exit the terminal and login again to reload the env.
- Build libvirt provider
```
go get github.com/dmacvicar/terraform-provider-libvirt
go install github.com/dmacvicar/terraform-provider-libvirt
```
```
ls $HOME/go/bin/
```
- Download Terraform
```
wget https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_linux_amd64.zip
unzip terraform_0.12.20_linux_amd64.zip
chmod +x terraform
mv terraform /usr/local/bin/
```
- Init Terraform
this will create folder `.terraform.d` in your home directory
```
terraform init
```
- Copy Program
```
mkdir $HOME/.terraform.d/plugins
cp $HOME/go/bin/terraform-provider-libvirt $HOME/.terraform.d/plugins
ls $HOME/.terraform.d/plugins
```
