package utils

import (
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
)

func RectOf(x, y, width, height float64) foundation.Rect {
	return foundation.Rect{Origin: foundation.Point{X: x, Y: y}, Size: foundation.Size{Width: width, Height: height}}
}

func LayoutActive(constraints ...appkit.LayoutConstraint) {
	for _, constraint := range constraints {
		constraint.SetActive(true)
	}
}

func SymbolImage(symbol string, configurations ...appkit.ImageSymbolConfiguration) appkit.Image {
	image := appkit.Image_ImageWithSystemSymbolNameAccessibilityDescription(symbol, symbol)
	for _, conf := range configurations {
		image = image.ImageWithSymbolConfiguration(conf)
	}
	return image
}

func DividerColor() appkit.Color {
	if IsDark(appkit.Appearance_CurrentDrawingAppearance()) {
		return appkit.Color_ColorWithSRGBRedGreenBlueAlpha(1, 1, 1, 0.15)
	}
	return appkit.Color_ColorWithSRGBRedGreenBlueAlpha(0, 0, 0, 0.15)
}

func IsDark(appearance appkit.IAppearance) bool {
	name := appkit.Appearance_CurrentDrawingAppearance().Name()
	if !appearance.IsNil() {
		name = appearance.Name()
	}
	return name == appkit.AppearanceNameDarkAqua ||
		name == appkit.AppearanceNameVibrantDark ||
		name == appkit.AppearanceNameAccessibilityHighContrastVibrantDark ||
		name == appkit.AppearanceNameAccessibilityHighContrastDarkAqua
}
