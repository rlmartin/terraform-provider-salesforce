default: build

clean:
	-rm .terraform.lock.hcl
	terraform init

plan:
	terraform plan -out=state.tfstate

build: clean plan apply

apply:
	terraform apply state.tfstate