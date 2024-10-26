build:
	@GOARCH=wasm GOOS=js go build -o docs/web/app.wasm ./src/sw1pr0g
	@go build -o docs/sw1pr0g ./src/sw1pr0g

clean:
	@echo "cleaning up.."
	@go clean ./...
	@-rm -f web/sw1site
#	@PORT=$$(yq e '.http.port' config/config.yml) && echo "port: $$PORT" && \
#    	PID=$$(lsof -t -i:$$PORT) && echo "PID: $$PID\n successfully cleaned!" || echo "No process found on port $$PORT" && \
#    	if [ -n "$$PID" ]; then kill $$PID; else echo "No process to kill"; fi