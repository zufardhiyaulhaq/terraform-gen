build:
	go build -o /usr/local/bin/terraform-gen ./generator/...

install: build
	mkdir -p /usr/local/terraform-gen
	mkdir -p /usr/local/terraform-gen/templates
	cp -R templates/* /usr/local/terraform-gen/templates/

remove:
	rm -rf /usr/local/bin/terraform-gen
	rm -rf /usr/local/terraform-gen