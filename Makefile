
all:
	GOOS=js GOARCH=wasm go build -o main.wasm ./

update-wasm:
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js wasm_exec.js

.PHONY: all
