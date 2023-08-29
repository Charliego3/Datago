package main

import (
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
	"runtime"
)

var minFrameSize = foundation.Size{Width: 700, Height: 600}

func launched(app appkit.Application, delegate *appkit.ApplicationDelegate) {
	frame := foundation.Rect{Size: minFrameSize}
	sidebarViewController := appkit.NewViewController()
	sidebarViewController.SetView(getSidebar())
	contentViewController := appkit.NewViewController()
	contentViewController.SetView(appkit.NewView())
	splitViewController := appkit.NewSplitViewController()
	splitViewController.AddSplitViewItem(appkit.SplitViewItem_SidebarWithViewController(sidebarViewController))
	splitViewController.AddSplitViewItem(appkit.SplitViewItem_SplitViewItemWithViewController(contentViewController))
	splitViewController.View().SetTranslatesAutoresizingMaskIntoConstraints(false)
	splitViewController.View().SetFrame(frame)
	screenFrame := appkit.Screen_MainScreen().Frame()
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
	configureToolbar(w.Toolbar())
	w.SetContentMinSize(minFrameSize)
	w.SetContentMaxSize(screenFrame.Size)
	w.SetToolbarStyle(appkit.WindowToolbarStyleUnifiedCompact)
	w.SetTitlebarAppearsTransparent(true)
	w.SetContentViewController(splitViewController)
	w.Center()
	w.MakeKeyAndOrderFront(w)

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
