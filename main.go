package main

import (
	"github.com/progrium/macdriver/helper/layout"
	"github.com/progrium/macdriver/macos"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

func main() {
	macos.RunApp(func(app appkit.Application, delegate *appkit.ApplicationDelegate) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)

		frame := foundation.Rect{Size: foundation.Size{Width: 800, Height: 600}}
		sidebarView := appkit.NewView()
		label := appkit.NewLabel("text")
		label.SetTranslatesAutoresizingMaskIntoConstraints(false)
		sidebarView.AddSubview(label)
		label.TopAnchor().ConstraintEqualToAnchorConstant(sidebarView.TopAnchor(), 38).SetActive(true)

		layout.SetMinWidth(sidebarView, 300)
		sidebarViewController := appkit.NewViewController()
		sidebarViewController.SetView(sidebarView)
		contentViewController := appkit.NewViewController()
		contentViewController.SetView(appkit.NewView())
		splitViewController := appkit.NewSplitViewController()
		splitViewController.AddSplitViewItem(appkit.SplitViewItem_SidebarWithViewController(sidebarViewController))
		splitViewController.AddSplitViewItem(appkit.SplitViewItem_SplitViewItemWithViewController(contentViewController))
		splitViewController.View().SetFrameSize(frame.Size)

		w := appkit.Window_InitWithContentRectStyleMaskBackingDefer(frame,
			appkit.ClosableWindowMask|
				appkit.TitledWindowMask|
				appkit.WindowStyleMaskResizable|
				appkit.MiniaturizableWindowMask|
				appkit.WindowStyleMaskFullSizeContentView|
				appkit.WindowStyleMaskUnifiedTitleAndToolbar,
			appkit.BackingStoreBuffered, false)

		toolbarDelegate := &appkit.ToolbarDelegate{}
		toolbarDelegate.SetToolbarAllowedItemIdentifiers(func(toolbar appkit.Toolbar) []appkit.ToolbarItemIdentifier {
			return []appkit.ToolbarItemIdentifier{
				appkit.ToolbarFlexibleSpaceItemIdentifier,
				appkit.ToolbarToggleSidebarItemIdentifier,
				appkit.ToolbarSidebarTrackingSeparatorItemIdentifier,
				appkit.ToolbarShowColorsItemIdentifier,
			}
		})
		toolbarDelegate.SetToolbarDefaultItemIdentifiers(func(toolbar appkit.Toolbar) []appkit.ToolbarItemIdentifier {
			return []appkit.ToolbarItemIdentifier{
				appkit.ToolbarToggleSidebarItemIdentifier,
				appkit.ToolbarSidebarTrackingSeparatorItemIdentifier,
				appkit.ToolbarShowColorsItemIdentifier,
			}
		})
		toolbarDelegate.SetToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(func(toolbar appkit.Toolbar, identifier appkit.ToolbarItemIdentifier, flag bool) appkit.IToolbarItem {
			if identifier == appkit.ToolbarToggleSidebarItemIdentifier {
				item := appkit.ToolbarItem_InitWithItemIdentifier(appkit.ToolbarToggleSidebarItemIdentifier)
				icon := appkit.Image_ImageWithSystemSymbolNameAccessibilityDescription("sidebar.leading", "")
				icon = icon.ImageWithSymbolConfiguration(appkit.ImageSymbolConfiguration_ConfigurationWithScale(appkit.ImageSymbolScaleLarge))
				item.SetImage(icon)
				item.SetAction(objc.Sel("toggleSidebar:"))
				item.SetTarget(splitViewController)
				return item
			}
			if identifier == appkit.ToolbarSidebarTrackingSeparatorItemIdentifier {
				item := appkit.TrackingSeparatorToolbarItem_InitWithItemIdentifier(identifier)
				item.SetSplitView(splitViewController.SplitView())
				item.SetDividerIndex(1)
				return item
			}
			if identifier == appkit.ToolbarShowColorsItemIdentifier {
				return appkit.SearchToolbarItem_InitWithItemIdentifier(identifier)
			}
			return nil
		})

		toolbar := appkit.NewToolbar()
		toolbar.SetDisplayMode(appkit.ToolbarDisplayModeIconOnly)
		toolbar.SetShowsBaselineSeparator(true)
		toolbar.SetDelegate(toolbarDelegate)
		w.SetToolbar(toolbar)
		w.SetToolbarStyle(appkit.WindowToolbarStyleUnifiedCompact)
		w.SetTitlebarAppearsTransparent(true)
		w.SetContentViewController(splitViewController)
		w.MakeKeyAndOrderFront(w)
		w.Center()

		delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
			return true
		})
	})
}
