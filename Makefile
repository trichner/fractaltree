

TARGET_DIR=docs

all: clean
	@mkdir -p $(TARGET_DIR)
	@cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js $(TARGET_DIR)/wasm_exec.js
	@cp index.html $(TARGET_DIR)/index.html
	@GOOS=js GOARCH=wasm go build -o $(TARGET_DIR)/main.wasm ./

clean:
	-rm -rf $(TARGET_DIR)

update-wasm:
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js wasm_exec.js

.PHONY: all
