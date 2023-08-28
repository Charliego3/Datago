package main

import (
	"github.com/progrium/macdriver/helper/layout"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

func getSidebar() appkit.IView {
	contentView := appkit.NewView()
	contentView.SetFrameSize(foundation.Size{Width: 200, Height: minFrameSize.Height})

	topLine := appkit.NewBoxWithFrame(rectOf(0, 0, 200, 200))
	contentView.AddSubview(topLine)
	//topLine.TopAnchor().ConstraintEqualToAnchorConstant(contentView.TopAnchor(), 38).SetActive(true)
	//topLine.LeadingAnchor().ConstraintEqualToAnchor(contentView.LeadingAnchor()).SetActive(true)
	//topLine.TrailingAnchor().ConstraintEqualToAnchor(contentView.TrailingAnchor()).SetActive(true)
	//topLine.BottomAnchor().ConstraintEqualToAnchorConstant(contentView.BottomAnchor(), 200)

	segment := appkit.NewSegmentedControl()
	segment.SetFocusRingType(appkit.FocusRingTypeNone)
	segment.SetFrameSize(foundation.Size{Width: contentView.Frame().Size.Width, Height: 20})
	segment.SetHorizontalContentSizeConstraintActive(true)
	segment.SetSegmentCount(2)
	segment.SetImageForSegment(appkit.Image_ImageWithSystemSymbolNameAccessibilityDescription("square.and.arrow.up.on.square", ""), 0)
	segment.SetImageForSegment(appkit.Image_ImageWithSystemSymbolNameAccessibilityDescription("eraser.line.dashed.fill", ""), 1)
	segment.SetShowsMenuIndicatorForSegment(false, 0)
	segment.SetShowsMenuIndicatorForSegment(false, 1)
	segment.SetAutoresizesSubviews(true)
	segment.SetSegmentStyle(appkit.SegmentStyleRoundRect)
	segment.SetSelectedSegmentBezelColor(appkit.Color_OrangeColor())
	segment.SetSpringLoaded(true)
	segment.SetSegmentDistribution(appkit.SegmentDistributionFillProportionally)
	segment.SetAlignment(appkit.TextAlignmentCenter)
	segment.DrawPageBorderWithSize(foundation.Size{})
	segment.SetWantsLayer(true)
	//segment.EnclosingScrollView().SetBorderType(appkit.NoBorder)
	segment.SetTranslatesAutoresizingMaskIntoConstraints(false)
	contentView.AddSubview(segment)
	layout.PinEdgesToSuperView(segment, foundation.EdgeInsets{
		Top:    100,
		Left:   20,
		Bottom: minFrameSize.Height - 50,
		Right:  20,
	})

	//layout.PinAnchorTo(segment.View.TopAnchor(), contentView.TopAnchor(), 50)

	outline := appkit.NewOutlineView()
	outline.SetFrameSize(foundation.Size{Width: 200, Height: minFrameSize.Height})
	setSidebarDataSource(outline)
	contentView.AddSubview(outline)
	return contentView
}

func setSidebarDataSource(outline appkit.OutlineView) {
	var dataSource IOutlineDataSource = OutlineDataSource{}
	po0 := objc.WrapAsProtocol("NSOutlineViewDataSource", dataSource)
	objc.SetAssociatedObject(outline, objc.AssociationKey("setDataSource"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](outline, objc.Sel("setDataSource:"), po0)
}

type IOutlineDataSource interface {
	HasOutlineViewNumberOfChildrenOfItem() bool
	OutlineViewNumberOfChildrenOfItem(outlineView appkit.IOutlineView, item objc.Object) int
}

type OutlineDataSource struct {
	objc.Object
}

func (d OutlineDataSource) HasOutlineViewNumberOfChildrenOfItem() bool {
	return d.RespondsToSelector(objc.Sel("outlineView:numberOfChildrenOfItem:"))
}

func (d OutlineDataSource) OutlineViewNumberOfChildrenOfItem(outlineView appkit.IOutlineView, item objc.Object) int {
	objc.Call[int](d, objc.Sel("outlineView:numberOfChildrenOfItem:"), objc.Ptr(outlineView), item)
	return 10
}
