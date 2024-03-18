.PHONY: build
build: ## build binary at /bin/
	@go mod tidy
	GOOS=linux GOARCH=amd64 go build -o bin/toolbox-linux-amd64
	GOOS=linux GOARCH=386 go build -o bin/toolbox-linux-386
	GOOS=darwin GOARCH=amd64 go build -o bin/toolbox-macos-amd64
	GOOS=windows GOARCH=amd64 go build -o bin/toolbox-windows-amd64.exe
	GOOS=windows GOARCH=386 go build -o bin/toolbox-windows-386.exe