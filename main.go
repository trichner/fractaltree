//go:build js && wasm
// +build js,wasm

package main

import (
	"honnef.co/go/js/dom/v2"
	"syscall/js"
)

func main() {
	js.Global().Set("mountTree", js.FuncOf(mountTreeFunction))

	select {
	//nop: keep wasm waiting
	}
}

func mountTreeFunction(this js.Value, p []js.Value) interface{} {

	treeElId := p[0].String()
	controlsElId := p[1].String()
	detailsElId := p[2].String()

	var document = dom.GetWindow().Document()

	treeEl := document.GetElementByID(treeElId)
	controlsEl := document.GetElementByID(controlsElId)
	detailsEl := document.GetElementByID(detailsElId).(*dom.HTMLDivElement)

	MountTree(treeEl, controlsEl, detailsEl)
	return nil
}
