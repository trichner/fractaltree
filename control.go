package main

import (
	"honnef.co/go/js/dom/v2"
	"math"
	"strconv"
)

type Controls struct {
	scaleEl, angleEl, reseedEl *dom.HTMLInputElement
}

func (c Controls) RegisterRedraw(f func(event dom.Event)) {

	if c.angleEl != nil {
		c.angleEl.AddEventListener("input", true, f)
	}
	if c.scaleEl != nil {
		c.scaleEl.AddEventListener("input", true, f)
	}
}

func (c Controls) RegisterReseed(f func(event dom.Event)) {
	if c.reseedEl != nil {
		c.reseedEl.AddEventListener("click", true, f)
	}
}

func (c Controls) Angle() float64 {

	maxAngle, _ := readMaxFromInput(c.angleEl)
	angle := c.angleEl.ValueAsNumber() / (float64(maxAngle)) * math.Pi
	return angle
}

func (c Controls) Scale() float64 {

	// empirically determined magic value
	const scaleFactor = 1.1

	maxScale, _ := readMaxFromInput(c.scaleEl)
	scale := c.scaleEl.ValueAsNumber() / (float64(maxScale) * scaleFactor)
	return scale
}

func readMaxFromInput(el *dom.HTMLInputElement) (int64, error) {
	return strconv.ParseInt(el.Max(), 10, 32)
}

func MountControlsContainer(element dom.Element) Controls {
	html := `
	<div class="controls-container" style="width: 30vw">
	<div>
	<label for="angle-control">Angle</label><br>
	<input type="range" min="0" max="1000" value="150" class="slider control-slider angle-control" id="angle-control">
	</div>
	<div>
	<label for="scale-control">Scale</label><br>
	<input type="range" min="0" max="999" value="850" class="slider control-slider scale-control" id="scale-control">
	</div>
	<div>
	<input type="button" class="control-button reseed-control" value="Reseed">
	</div>
	</div>`

	element.SetInnerHTML(html)

	return getControls(element)
}

func getControls(controlsEl dom.Element) Controls {
	angleControlEl := controlsEl.GetElementsByClassName("angle-control")[0].(*dom.HTMLInputElement)
	scaleControlEl := controlsEl.GetElementsByClassName("scale-control")[0].(*dom.HTMLInputElement)
	reseedControlEl := controlsEl.GetElementsByClassName("reseed-control")[0].(*dom.HTMLInputElement)

	return Controls{
		scaleEl:  scaleControlEl,
		angleEl:  angleControlEl,
		reseedEl: reseedControlEl,
	}
}
