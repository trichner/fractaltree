package main

import (
	"fmt"
	svg "github.com/ajstarks/svgo"
	"strings"
)

func round(x float64) int {
	return int(x + 0.5)
}

func hsl(hue, saturation, light float64) string {

	return fmt.Sprintf("hsl(%f,%f%%,%f%%)", hue, saturation, light)
}

const zoom = 100

func getBoundingBox(lines []Line) AABB {

	var aabb = lines[0].BoundingBox()

	lines = lines[1:]
	for _, l := range lines {
		aabb = aabb.Union(l.BoundingBox())
	}

	return aabb
}

func calculateViewBox(lines []Line) string {
	bb := getBoundingBox(lines)
	return fmt.Sprintf("viewBox=\"%d %d %d %d\"", int(bb.Origin.X*zoom), int(bb.Origin.Y*zoom), int(bb.Size.X*zoom+0.5), int(bb.Size.Y*zoom+0.5))
}

func renderSvg(lines []Line) string {

	stringWriter := &strings.Builder{}
	canvas := svg.New(stringWriter)
	canvas.Startraw(calculateViewBox(lines))
	for _, l := range lines {
		//stroke := l.Len() * zoom * 1.0 / 20.0
		stroke := l.Len() * zoom * 1.0 / 5.0
		strokeWidthStyle := fmt.Sprintf("stroke-width:%d", int(stroke+0.5))
		strokeWidthColor := fmt.Sprintf("stroke:%s", hsl(40+(2-l.Len())*50, 50, 50))
		canvas.Line(round(l.from.X*zoom), round(l.from.Y*zoom), round(l.to.X*zoom), round(l.to.Y*zoom), strokeWidthColor+";"+strokeWidthStyle)
	}
	canvas.End()
	return stringWriter.String()
}
