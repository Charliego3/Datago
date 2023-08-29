package main

import (
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
)

func rectOf(x, y, width, height float64) foundation.Rect {
	return foundation.Rect{Origin: foundation.Point{X: x, Y: y}, Size: foundation.Size{Width: width, Height: height}}
}

func layoutActive(constraints ...appkit.LayoutConstraint) {
	for _, constraint := range constraints {
		constraint.SetActive(true)
	}
}

func symbolImage(symbol string) appkit.Image {
	return appkit.Image_ImageWithSystemSymbolNameAccessibilityDescription(symbol, "")
}
