package main

import (
	"github.com/progrium/macdriver/helper/action"
	"github.com/progrium/macdriver/helper/layout"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
	"runtime"
)

func rectOf(x, y, width, height float64) foundation.Rect {
	return foundation.Rect{Origin: foundation.Point{X: x, Y: y}, Size: foundation.Size{Width: width, Height: height}}
}

func launched(app appkit.Application, delegate *appkit.ApplicationDelegate) {
	frame := foundation.Rect{Size: foundation.Size{Width: 800, Height: 600}}

	sidebarView := appkit.NewView()
	datePicker := appkit.NewDatePickerWithFrame(rectOf(450, 290, 140, 25))
	sidebarView.AddSubview(datePicker)

	// buttons
	cb := appkit.NewCheckBox("check box")
	cb.SetFrame(rectOf(10, 250, 80, 25))
	sidebarView.AddSubview(cb)

	rb := appkit.NewRadioButton("radio button")
	rb.SetFrame(rectOf(150, 250, 120, 25))
	sidebarView.AddSubview(rb)

	sw := appkit.NewSwitchWithFrame(rectOf(260, 250, 120, 25))
	sidebarView.AddSubview(sw)

	li := appkit.NewLevelIndicatorWithFrame(rectOf(370, 250, 120, 25))
	li.SetCriticalValue(4)
	li.SetDoubleValue(3)
	sidebarView.AddSubview(li)

	btn := appkit.NewButtonWithTitle("change color")
	btn.SetFrame(rectOf(10, 160, 120, 25))
	sidebarView.AddSubview(btn)

	quitBtn := appkit.NewButtonWithTitle("Quit")
	quitBtn.SetFrame(rectOf(10, 130, 80, 25))
	action.Set(quitBtn, func(sender objc.Object) {
		app.Terminate(nil)
	})
	sidebarView.AddSubview(quitBtn)

	tf := appkit.NewTextField()
	sidebarView.AddSubview(tf)
	tf.SetFrame(rectOf(10, 100, 150, 25))

	layout.SetMinWidth(sidebarView, 300)
	sidebarViewController := appkit.NewViewController()
	sidebarViewController.SetView(sidebarView)
	contentViewController := appkit.NewViewController()
	contentViewController.SetView(appkit.NewView())
	splitViewController := appkit.NewSplitViewController()
	splitViewController.AddSplitViewItem(appkit.SplitViewItem_SidebarWithViewController(sidebarViewController))
	splitViewController.AddSplitViewItem(appkit.SplitViewItem_SplitViewItemWithViewController(contentViewController))
	splitViewController.View().SetFrameSize(frame.Size)

	w := appkit.NewWindowWithContentRectStyleMaskBackingDefer(frame,
		appkit.ClosableWindowMask|
			appkit.TitledWindowMask|
			appkit.WindowStyleMaskResizable|
			appkit.MiniaturizableWindowMask|
			appkit.WindowStyleMaskFullSizeContentView|
			appkit.WindowStyleMaskUnifiedTitleAndToolbar,
		appkit.BackingStoreBuffered, false)
	objc.Retain(&w)

	w.SetToolbar(getToolbar())
	w.SetToolbarStyle(appkit.WindowToolbarStyleUnifiedCompact)
	w.SetTitlebarAppearsTransparent(true)
	w.SetContentViewController(splitViewController)
	w.Center()
	w.MakeKeyAndOrderFront(w)
	configureToolbar(w.Toolbar())

	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)
}

func setMainMenu(app appkit.Application) {
	menu := appkit.NewMenuWithTitle("main")
	app.SetMainMenu(menu)

	mainMenuItem := appkit.NewMenuItemWithSelector("", "", objc.Selector{})
	mainMenuMenu := appkit.NewMenuWithTitle("App")
	mainMenuMenu.AddItem(appkit.NewMenuItemWithAction("Hide", "h", func(sender objc.Object) { app.Hide(nil) }))
	mainMenuMenu.AddItem(appkit.NewMenuItemWithAction("Quit", "q", func(sender objc.Object) { app.Terminate(nil) }))
	mainMenuItem.SetSubmenu(mainMenuMenu)
	menu.AddItem(mainMenuItem)

	testMenuItem := appkit.NewMenuItemWithSelector("", "", objc.Selector{})
	testMenu := appkit.NewMenuWithTitle("Edit")
	testMenu.AddItem(appkit.NewMenuItemWithSelector("Select All", "a", objc.Sel("selectAll:")))
	testMenu.AddItem(appkit.MenuItem_SeparatorItem())
	testMenu.AddItem(appkit.NewMenuItemWithSelector("Copy", "c", objc.Sel("copy:")))
	testMenu.AddItem(appkit.NewMenuItemWithSelector("Paste", "v", objc.Sel("paste:")))
	testMenu.AddItem(appkit.NewMenuItemWithSelector("Cut", "x", objc.Sel("cut:")))
	testMenu.AddItem(appkit.NewMenuItemWithSelector("Undo", "z", objc.Sel("undo:")))
	testMenu.AddItem(appkit.NewMenuItemWithSelector("Redo", "Z", objc.Sel("redo:")))
	testMenuItem.SetSubmenu(testMenu)
	menu.AddItem(testMenuItem)
}

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	app := appkit.Application_SharedApplication()
	delegate := &appkit.ApplicationDelegate{}
	delegate.SetApplicationDidFinishLaunching(func(notification foundation.Notification) {
		launched(app, delegate)
	})
	delegate.SetApplicationWillFinishLaunching(func(foundation.Notification) {
		setMainMenu(app)
	})
	delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
		return true
	})
	app.SetDelegate(delegate)
	app.Run()
}
