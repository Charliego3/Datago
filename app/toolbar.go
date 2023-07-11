package main

import (
	"log"

	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

// var id1 = core.String("NSToolbarToggleSidebarItemIdentifier")
// var items = core.NSArray_WithObjects(&id1)

// type ToolbarDelegate struct {
// 	objc.Object `objc:"ToolbarDelegate : NSObject"`
// }

// func (t *ToolbarDelegate) ItemFor(toolbar, identifier objc.Object, will bool) objc.Object {
// 	//item := cocoa.NSToolbarItem_Alloc()
// 	//item.Set("itemIdentifier:", "NSToolbarToggleSidebarItemIdentifier")
// 	//return item
// 	return nil
// }

// func (t *ToolbarDelegate) AllowedItems(_ objc.Object) objc.Object {
// 	log.Println("Allowed items called..........")
// 	return items
// }

// func (t *ToolbarDelegate) DefaultItems(_ objc.Object) objc.Object {
// 	log.Println("Default items called..........")
// 	return &items
// }

func init() {
	sidebarItem := cocoa.NSToolbarItem_Alloc()
	sidebarItem.SetTitle(core.String("Sidebar"))
	// sidebarItem.SetImage(cocoa.NSImage_SystemSymbolName("sidebar.leading"))
	sidebarItem.SetEnabled(true)
	sidebarItem.SetBordered(true)
	sidebarItem.SetLabel(core.String("Sidebar label"))

	log.Println(objc.Get("ToolbarItemIdentifier"))
	c := objc.NewClass("ToolbarDelegate", "NSObject")
	c.AddMethod("toolbar:itemForItemIdentifier:willBeInsertedIntoToolbar:", func(toolbar, identifier objc.Object, will bool) objc.Object {
		return nil
	})
	c.AddMethod("toolbarAllowedItemIdentifiers:", func(_ objc.Object) objc.Object {
		log.Println("toolbarAllowedItemIdentifiers items called..........")
		return core.NSArray_FromRef(objc.Get("NSArray").Send("arrayWithObjects:", core.String("NSToolbarToggleSidebarItemIdentifier")))
	})
	// c.AddMethod("toolbarDefaultItemIdentifiers:", func(_ objc.Object) objc.Object {
	// 	log.Println("toolbarDefaultItemIdentifiers items called..........")
	// 	// return core.NSArray_FromRef(objc.Get("NSArray").Send("arrayWithObjects:", sidebarItem))
	// 	return core.NSArray_FromRef(objc.Get("NSArray").Send("arrayWithObjects:", core.String("NSToolbarToggleSidebarItemIdentifier")))
	// })
	objc.RegisterClass(c)
}
