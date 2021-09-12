package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"honnef.co/go/js/dom/v2"
	"strings"
	"time"
)

type Tree struct {
	treeEl    dom.Element
	detailsEl dom.Element
	controls  Controls
	generator *TreeGenerator
}

func MountTree(treeContainer, controlsContainer dom.Element, detailsContainer *dom.HTMLDivElement) *Tree {

	controls := MountControlsContainer(controlsContainer)

	generator := NewDefaultTreeGenerator()

	const name = "Margaret"
	generator.SetSeed(generateSeed(name))

	t := &Tree{
		treeEl:    treeContainer,
		detailsEl: detailsContainer,
		controls:  controls,
		generator: generator,
	}

	controls.RegisterRedraw(func(event dom.Event) {
		t.Redraw()
	})

	controls.RegisterReseed(func(event dom.Event) {
		t.SetSeed(pseudoRandomSeed())
		t.Redraw()
	})

	t.Redraw()

	return t
}

func (t *Tree) Redraw() {
	t.SetScale(t.controls.Scale())
	t.SetAngle(t.controls.Angle())

	t.treeEl.SetInnerHTML(renderSvg(t.generator.Generate()))

	if t.detailsEl != nil {
		detailsHtml := renderGeneratorDetails(t.generator)
		t.detailsEl.SetInnerHTML(detailsHtml)
	}
}

func (t *Tree) SetSeed(seed int64) {
	t.generator.SetSeed(seed)
}

func (t *Tree) SetAngle(angle float64) {
	t.generator.SetAngle(angle)
}

func (t *Tree) SetScale(scale float64) {
	t.generator.SetScale(scale)
}

func renderGeneratorDetails(generator *TreeGenerator) string {
	var sb strings.Builder
	var bb bytes.Buffer
	binary.Write(&bb, binary.BigEndian, generator.Seed())
	sb.WriteString(fmt.Sprintf("seed:&nbsp;&nbsp;%s", hex.EncodeToString(bb.Bytes())))
	sb.WriteString("<br>")
	sb.WriteString(fmt.Sprintf("angle:&nbsp;%.3f (%.3f)", generator.Angle(), generator.AngleVariation()))
	sb.WriteString("<br>")
	sb.WriteString(fmt.Sprintf("scale:&nbsp;%.3f (%.3f)", generator.Scale(), generator.ScaleVariation()))
	return sb.String()
}

func pseudoRandomSeed() int64 {

	var bb bytes.Buffer
	err := binary.Write(&bb, binary.BigEndian, time.Now().Unix())
	if err != nil {
		panic(err)
	}

	digest := sha256.Sum256(bb.Bytes())
	var seed int64
	err = binary.Read(bytes.NewReader(digest[:8]), binary.BigEndian, &seed)
	if err != nil {
		panic(err)
	}
	return seed
}

func generateSeed(s string) int64 {
	digest := sha256.Sum256([]byte(s))
	var seed int64
	binary.Read(bytes.NewReader(digest[:8]), binary.BigEndian, &seed)
	return seed
}
