# go mod tidy
.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: run.demo
run.demo:
	@go run cmd/demo/demo.go --config="configs/demo.yaml"

.PHONY: help.demo
help.demo:
	@go run cmd/demo/demo.go --help