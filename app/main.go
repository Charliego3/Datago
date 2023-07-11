package main

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

func main() {
	app := cocoa.NSApp_WithDidLaunch(func(notification objc.Object) {
		window := cocoa.NSWindow_New()
		window.SetStyleMask(
			cocoa.NSClosableWindowMask |
				cocoa.NSMiniaturizableWindowMask |
				cocoa.NSFullSizeContentViewWindowMask |
				cocoa.NSResizableWindowMask |
				cocoa.NSTitledWindowMask,
		)

		frame := core.Rect(0, 0, 300, 500)

		tempView := cocoa.NSTextView_Alloc().Init_AsNSTextView()
		tempView.SetString("temp text view")
		rootView := cocoa.NSView_Init(frame)
		rootView.AddSubview(tempView)

		cocoa.NSSegmentedControl_Alloc().Init_AsNSSegmentedControl()

		controller := cocoa.NSSplitViewController_Alloc().Init_AsNSSplitViewController()
		controller.SetMinimumThicknessForInlineSidebars(300)
		sidebarViewController := cocoa.NSViewController_New()
		sidebarViewController.SetView(rootView)
		controller.AddSplitViewItem(cocoa.NSSplitViewItem_SidebarWithViewController(sidebarViewController))

		contentViewController := cocoa.NSViewController_New()
		contentViewController.SetView(cocoa.NSView_Init(frame))
		controller.AddSplitViewItem(cocoa.NSSplitViewItem_SplitViewItemWithViewController(contentViewController))

		toolbar := cocoa.NSToolbar_Alloc().Init_AsNSToolbar()
		toolbar.SetDelegate(objc.Get("ToolbarDelegate").Alloc().InitObject())

		window.SetToolbar(toolbar)
		window.SetShowsToolbarButton(true)
		window.SetTitlebarAppearsTransparent(true)
		window.Set("toolbarStyle:", core.NSUInteger(4))
		window.Set("titlebarSeparatorStyle:", core.NSUInteger(2))
		window.SetHasShadow(true)
		window.SetTitle("MacDriver Demo")
		window.Retain()
		window.SetContentViewController(controller)
		window.Center()
		window.SetMinSize(core.Size(400, 500))
		window.SetBackingType(core.NSUInteger(cocoa.NSBackingStoreRetained))
		window.MakeKeyAndOrderFront(nil)
	})

	app.SetActivationPolicy(cocoa.NSApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)
	app.Run()
}
