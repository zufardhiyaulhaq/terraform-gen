build:
	go build -o /usr/local/bin/terraform-gen ./generator/...

install: build
	mkdir /usr/local/terraform-gen
	mkdir /usr/local/terraform-gen/templates
	cp templates/* /usr/local/terraform-gen/templates/