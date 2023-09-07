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

func symbolImage(symbol string, configurations ...appkit.ImageSymbolConfiguration) appkit.Image {
	image := appkit.Image_ImageWithSystemSymbolNameAccessibilityDescription(symbol, symbol)
	for _, conf := range configurations {
		image = image.ImageWithSymbolConfiguration(conf)
	}
	return image
}

func getDividerColor() appkit.Color {
	if isDark(appkit.Appearance_CurrentDrawingAppearance()) {
		return appkit.Color_ColorWithSRGBRedGreenBlueAlpha(1, 1, 1, 0.15)
	}
	return appkit.Color_ColorWithSRGBRedGreenBlueAlpha(0, 0, 0, 0.15)
}

func isDark(appearance appkit.IAppearance) bool {
	name := appearance.Name()
	return name == appkit.AppearanceNameDarkAqua ||
		name == appkit.AppearanceNameVibrantDark ||
		name == appkit.AppearanceNameAccessibilityHighContrastVibrantDark ||
		name == appkit.AppearanceNameAccessibilityHighContrastDarkAqua
}
