package main

import (
	"github.com/progrium/macdriver/macos"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

func main() {
	macos.RunApp(func(app appkit.Application, delegate *appkit.ApplicationDelegate) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)

		frame := foundation.Rect{Size: foundation.Size{800, 600}}
		sidebarViewController := appkit.NewViewController()
		contentViewController := appkit.ViewController_Alloc()
		splitViewController := appkit.SplitViewController_Alloc()
		splitViewController.AddSplitViewItem(appkit.SplitViewItem_SidebarWithViewController(sidebarViewController))
		splitViewController.AddSplitViewItem(appkit.SplitViewItem_SplitViewItemWithViewController(contentViewController))

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
				appkit.ToolbarToggleSidebarItemIdentifier,
			}
		})
		toolbarDelegate.SetToolbarDefaultItemIdentifiers(func(toolbar appkit.Toolbar) []appkit.ToolbarItemIdentifier {
			return []appkit.ToolbarItemIdentifier{
				appkit.ToolbarToggleSidebarItemIdentifier,
			}
		})
		toolbarDelegate.SetToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(func(toolbar appkit.Toolbar, identifier appkit.ToolbarItemIdentifier, flag bool) appkit.IToolbarItem {
			if identifier == appkit.ToolbarToggleSidebarItemIdentifier {
				item := appkit.ToolbarItem_InitWithItemIdentifier(appkit.ToolbarToggleSidebarItemIdentifier)
				item.SetAction(objc.Sel("toggleSidebar"))
				item.SetTarget(splitViewController)
				return item
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
