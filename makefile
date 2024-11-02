build:
	@GOARCH=wasm GOOS=js go build -o docs/web/app.wasm ./src
	@go build -o docs/sw1pr0g ./src

run: build
	@cd docs && ./sw1pr0g local

develop:
	@cd docs && air

build-github: build
	@cd docs && ./sw1pr0g github

github: build-github clean

# TODO("Implement tests")

clean:
	@go clean ./...
	@-rm docs/sw1pr0g