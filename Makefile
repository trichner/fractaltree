
all:
	@mkdir -p site
	@cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js site/wasm_exec.js
	@cp index.html site/index.html
	@GOOS=js GOARCH=wasm go build -o site/main.wasm ./

clean:
	-rm -rf site

update-wasm:
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js wasm_exec.js

.PHONY: all
