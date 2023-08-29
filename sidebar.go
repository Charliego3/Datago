package main

import (
	"github.com/progrium/macdriver/helper/layout"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

func getSidebar() appkit.IView {
	const defaultWidth = 200
	contentView := appkit.NewView()
	contentView.SetFrameSize(foundation.Size{Width: defaultWidth, Height: minFrameSize.Height})

	topLine := getHorizontalLine(defaultWidth)
	contentView.AddSubview(topLine)
	layoutActive(
		topLine.TopAnchor().ConstraintEqualToAnchorConstant(contentView.TopAnchor(), 38),
		topLine.LeadingAnchor().ConstraintEqualToAnchor(contentView.LeadingAnchor()),
		topLine.TrailingAnchor().ConstraintEqualToAnchor(contentView.TrailingAnchor()),
	)

	segment := getSegmentControl()
	contentView.AddSubview(segment)
	layoutActive(
		segment.TopAnchor().ConstraintEqualToAnchorConstant(topLine.TopAnchor(), 5.5),
		segment.LeadingAnchor().ConstraintEqualToAnchorConstant(contentView.LeadingAnchor(), 20),
		segment.TrailingAnchor().ConstraintEqualToAnchorConstant(contentView.TrailingAnchor(), -20),
	)

	bottomLine := getHorizontalLine(defaultWidth)
	contentView.AddSubview(bottomLine)
	layoutActive(
		bottomLine.TopAnchor().ConstraintEqualToAnchorConstant(segment.BottomAnchor(), 4),
		bottomLine.LeadingAnchor().ConstraintEqualToAnchor(contentView.LeadingAnchor()),
		bottomLine.TrailingAnchor().ConstraintEqualToAnchor(contentView.TrailingAnchor()),
	)

	tabViewController := appkit.NewTabViewController()
	tabViewController.SetTabViewItems([]appkit.ITabViewItem{
		appkit.TabViewItem_TabViewItemWithViewController(getConnectionController()),
	})

	tabView := tabViewController.TabView()
	tabView.SetTabViewBorderType(appkit.TabViewBorderTypeNone)
	tabView.SetTabViewType(appkit.NoTabsNoBorder)
	contentView.AddSubview(tabView)
	tabView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	layoutActive(
		tabView.TopAnchor().ConstraintEqualToAnchor(bottomLine.BottomAnchor()),
		tabView.LeadingAnchor().ConstraintEqualToAnchor(contentView.LeadingAnchor()),
		tabView.TrailingAnchor().ConstraintEqualToAnchor(contentView.TrailingAnchor()),
	)

	//outline := appkit.NewOutlineView()
	//outline.SetFrameSize(foundation.Size{Width: defaultWidth, Height: minFrameSize.Height})
	//setSidebarDataSource(outline)
	//contentView.AddSubview(outline)
	return contentView
}

func getConnectionController() appkit.ViewController {
	textField := appkit.NewLabel("Connections")
	textField.SetTranslatesAutoresizingMaskIntoConstraints(false)
	view := appkit.NewView()
	view.AddSubview(textField)
	layout.PinEdgesToSuperView(textField, foundation.EdgeInsets{})

	controller := appkit.NewViewController()
	controller.SetView(view)
	return controller
}

func getSegmentControl() appkit.SegmentedControl {
	segment := appkit.NewSegmentedControl()
	segment.SetFocusRingType(appkit.FocusRingTypeNone)
	segment.SetSegmentCount(2)
	segment.SetImageForSegment(symbolImage("square.and.arrow.up.on.square"), 0)
	segment.SetImageForSegment(symbolImage("eraser.line.dashed.fill"), 1)
	segment.SetAutoresizesSubviews(true)
	segment.SetSegmentStyle(appkit.SegmentStyleRoundRect)
	segment.SetSpringLoaded(true)
	segment.SetSegmentDistribution(appkit.SegmentDistributionFillProportionally)
	segment.SetAlignment(appkit.TextAlignmentCenter)
	segment.SetSelectedSegment(0)
	segment.SetSelectedSegmentBezelColor(appkit.Color_MagentaColor())
	segment.SetUserInterfaceLayoutDirection(appkit.UserInterfaceLayoutDirectionLeftToRight)
	segment.SetIgnoresMultiClick(true)
	segment.SetUsesSingleLineMode(true)
	segment.SetTranslatesAutoresizingMaskIntoConstraints(false)
	segment.SetWantsLayer(true)
	return segment
}

func getHorizontalLine(width float64) appkit.Box {
	line := appkit.NewBoxWithFrame(rectOf(0, 0, width, 1))
	line.SetTranslatesAutoresizingMaskIntoConstraints(false)
	line.SetBoxType(appkit.BoxCustom)
	line.SetBorderColor(appkit.Color_ColorWithSRGBRedGreenBlueAlpha(0, 0, 0, 0.1))
	return line
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
