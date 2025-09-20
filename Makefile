main_package_path=./cmd/api
binary_name=api

## tidy: tidy modfiles and format .go files
PHONY: tidy
tidy:
	go mod tidy -v
	go fmt './...'

## build: build the application
.PHONY: build
build:
	go build -o=./bin/${binary_name} -gcflags='all=-N -l' ${main_package_path}

## run: run the  application
.PHONY: run
run: build
	./bin/${binary_name}

.PHONY: debug
debug:
	dlv debug ${main_package_path}

