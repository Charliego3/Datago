package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type NSSegmentedControl struct {
	gen_NSSegmentedControl
}

func NSSegmentedControlWithIamges(fn interface{}, images ...NSImage) NSSegmentedControl {
	var objs []objc.Object
	for _, image := range images {
		objs = append(objs, objc.Object_FromRef(image))
	}
	var target objc.Object
	var selector objc.Selector
	if fn != nil {
		target, selector = core.Callback(fn)
	}
	return NSSegmentedControl_FromRef(NSSegmentedControl_Alloc().
		Send("segmentedControlWithImages:trackingMode:target:action:",
			core.NSArray_WithObjects(objs...), core.NSUInteger(0), target, selector,
		),
	)
}
