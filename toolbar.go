package main

import (
	"github.com/progrium/macdriver/macos/appkit"
)

func getToolbar() appkit.IToolbar {
	toolbar := appkit.NewToolbar()
	toolbar.SetDisplayMode(appkit.ToolbarDisplayModeIconOnly)
	toolbar.SetShowsBaselineSeparator(true)
	toolbar.SetDelegate(getToolbarDelegate())
	toolbar.SetAllowsExtensionItems(true)
	return toolbar
}

var itemIdentifiers = []appkit.ToolbarItemIdentifier{
	appkit.ToolbarToggleSidebarItemIdentifier,
	appkit.ToolbarSidebarTrackingSeparatorItemIdentifier,
	appkit.ToolbarCloudSharingItemIdentifier,
	appkit.ToolbarShowColorsItemIdentifier,
}

func toolbarItemIdentifiers(appkit.Toolbar) []appkit.ToolbarItemIdentifier {
	return itemIdentifiers
}

func configureToolbar(toolbar appkit.Toolbar) {
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

func getToolbarDelegate() *appkit.ToolbarDelegate {
	toolbarDelegate := &appkit.ToolbarDelegate{}
	toolbarDelegate.SetToolbarAllowedItemIdentifiers(toolbarItemIdentifiers)
	toolbarDelegate.SetToolbarDefaultItemIdentifiers(toolbarItemIdentifiers)
	toolbarDelegate.SetToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(func(
		toolbar appkit.Toolbar,
		identifier appkit.ToolbarItemIdentifier,
		flag bool,
	) appkit.IToolbarItem {
		return nil
	})
	return toolbarDelegate
}
