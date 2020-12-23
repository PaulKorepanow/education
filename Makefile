.PHONY: build
build:
		go build -v ./boolLibrary/api

.Default_Goal: build
