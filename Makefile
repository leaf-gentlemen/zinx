pre-commit:
	go mod tidy
	#fieldalignment -fix ./...
	golangci-lint run --issues-exit-code 1 -v "./..."
yaml-lint:
	sudo docker run --rm -v $(shell pwd):/code registry.gitlab.com/pipeline-components/yamllint yamllint .