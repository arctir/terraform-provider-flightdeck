VERSION = 1.0.0
PROVIDER_NAME = terraform-provider-flightdeck_v${VERSION}
PLUGIN_NAMESPACE = registry.terraform.io/arctir/flightdeck
API_SPEC ?= "https://${GITHUB_TOKEN}@raw.githubusercontent.com/arctir/flightdeck-api/main/generated/v1/api.gen.yaml"

.PHONY: all
all: provider

.PHONY: gen
gen:
	redocly bundle -d -o apispec.bundled.json --ext json ${API_SPEC}
	python scripts/generate.py apispec.bundled.json > spec.json
	tfplugingen-framework generate all --input spec.json --output pkg/provider
	go run github.com/jmattheis/goverter/cmd/goverter gen ./pkg/conversion 

.PHONY: docs
docs:
	cd tools; go generate ./...

.PHONY: deps
deps:
#	go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@v0.19.4
#	go install github.com/hashicorp/terraform-plugin-codegen-openapi/cmd/tfplugingen-openapi@latest
#	go get github.com/jmattheis/goverter/cmd/goverter@v1.5.1
	go install github.com/hashicorp/terraform-plugin-codegen-framework/cmd/tfplugingen-framework@v0.4.1

.PHONY: provider
provider: gen docs
	go build -o ${PROVIDER_NAME} .

.PHONY: install-dev
install-dev: provider
	mkdir -p ~/.terraform.d/plugins/${PLUGIN_NAMESPACE}/${VERSION}/linux_amd64
	cp ${PROVIDER_NAME} ~/.terraform.d/plugins/${PLUGIN_NAMESPACE}/${VERSION}/linux_amd64/
