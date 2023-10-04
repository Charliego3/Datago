package components

import (
	"github.com/charliego3/datago/utils"
	"github.com/progrium/macdriver/helper/action"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/objc"
)

func Toolbar(app appkit.Application) appkit.IToolbar {
	toolbar := appkit.NewToolbar()
	toolbar.SetDisplayMode(appkit.ToolbarDisplayModeIconOnly)
	toolbar.SetShowsBaselineSeparator(true)
	toolbar.SetDelegate(getToolbarDelegate(app))
	toolbar.SetAllowsExtensionItems(true)
	return toolbar
}

var itemIdentifiers = []appkit.ToolbarItemIdentifier{
	appkit.ToolbarToggleSidebarItemIdentifier,
	appkit.ToolbarSidebarTrackingSeparatorItemIdentifier,
	toggleThemeIdentifier,
	appkit.ToolbarCloudSharingItemIdentifier,
	appkit.ToolbarShowColorsItemIdentifier,
}

var toggleThemeIdentifier = appkit.ToolbarItemIdentifier("toggleTheme")

func toolbarItemIdentifiers(appkit.Toolbar) []appkit.ToolbarItemIdentifier {
	return itemIdentifiers
}

func ConfigureToolbar(toolbar appkit.Toolbar) {
	for _, item := range toolbar.Items() {
		if item.ItemIdentifier() == appkit.ToolbarToggleSidebarItemIdentifier {
			item.View().SetFocusRingType(appkit.FocusRingTypeNone)
			item.SetNavigational(false)
			item.SetImage(item.Image().ImageWithSymbolConfiguration(
				appkit.ImageSymbolConfiguration_ConfigurationWithScale(appkit.ImageSymbolScaleLarge),
			))
			break
		}
	}
}

func getToolbarDelegate(app appkit.Application) *appkit.ToolbarDelegate {
	toolbarDelegate := &appkit.ToolbarDelegate{}
	toolbarDelegate.SetToolbarAllowedItemIdentifiers(toolbarItemIdentifiers)
	toolbarDelegate.SetToolbarDefaultItemIdentifiers(toolbarItemIdentifiers)
	toolbarDelegate.SetToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(func(
		toolbar appkit.Toolbar,
		identifier appkit.ToolbarItemIdentifier,
		flag bool,
	) appkit.ToolbarItem {
		if identifier == toggleThemeIdentifier {
			btn := appkit.NewButton()
			target, selector := action.Wrap(func(_ objc.Object) {
				var theme appkit.IAppearance
				if utils.IsDark(app.Appearance()) {
					theme = appkit.Appearance_AppearanceNamed(appkit.AppearanceNameAqua)
				} else {
					theme = appkit.Appearance_AppearanceNamed(appkit.AppearanceNameDarkAqua)
				}
				app.SetAppearance(theme)
			})
			btn.SetButtonType(appkit.ToggleButton)
			btn.SetAction(selector)
			btn.SetTarget(target)
			btn.SetTitle("Toggle Theme")
			item := appkit.NewToolbarItemWithItemIdentifier(identifier)
			item.SetView(btn)
			item.SetNavigational(true)
			return item
		}
		return appkit.ToolbarItem{}
	})
	return toolbarDelegate
}