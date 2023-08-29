package main

import (
	"github.com/progrium/macdriver/helper/action"
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
	segment := getSegmentControl()
	contentView.AddSubview(segment)
	bottomLine := getHorizontalLine(defaultWidth)
	contentView.AddSubview(bottomLine)

	// TODO: TabViewItems Start
	textField := appkit.NewLabel("Label")
	controller := appkit.NewViewController()
	controller.SetView(textField)

	label := appkit.NewLabel("Label1111")
	controller1 := appkit.NewViewController()
	controller1.SetView(label)
	// TODO: TabViewItems Ended

	tabView := appkit.NewTabView()
	tabView.SetTabViewType(appkit.NoTabsNoBorder)
	tabView.SetTabViewBorderType(appkit.TabViewBorderTypeNone)
	tabView.SetUserInterfaceLayoutDirection(appkit.UserInterfaceLayoutDirectionLeftToRight)
	tabView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	tabView.SetTabViewItems([]appkit.ITabViewItem{
		appkit.TabViewItem_TabViewItemWithViewController(controller),
		appkit.TabViewItem_TabViewItemWithViewController(controller1),
	})

	contentView.AddSubview(tabView)
	segment.trigger = tabView.SelectTabViewItemAtIndex

	layoutActive(
		topLine.TopAnchor().ConstraintEqualToAnchorConstant(contentView.TopAnchor(), 38),
		topLine.LeadingAnchor().ConstraintEqualToAnchor(contentView.LeadingAnchor()),
		topLine.TrailingAnchor().ConstraintEqualToAnchor(contentView.TrailingAnchor()),
		segment.TopAnchor().ConstraintEqualToAnchorConstant(topLine.TopAnchor(), 5.3),
		segment.LeadingAnchor().ConstraintEqualToAnchorConstant(contentView.LeadingAnchor(), 20),
		segment.TrailingAnchor().ConstraintEqualToAnchorConstant(contentView.TrailingAnchor(), -20),
		bottomLine.TopAnchor().ConstraintEqualToAnchorConstant(segment.BottomAnchor(), 4),
		bottomLine.LeadingAnchor().ConstraintEqualToAnchor(contentView.LeadingAnchor()),
		bottomLine.TrailingAnchor().ConstraintEqualToAnchor(contentView.TrailingAnchor()),
		tabView.TopAnchor().ConstraintEqualToAnchor(bottomLine.BottomAnchor()),
		tabView.LeadingAnchor().ConstraintEqualToAnchor(contentView.LeadingAnchor()),
		tabView.TrailingAnchor().ConstraintEqualToAnchor(contentView.TrailingAnchor()),
		tabView.BottomAnchor().ConstraintEqualToAnchor(contentView.BottomAnchor()),
	)
	return contentView
}

type SegmentControl struct {
	appkit.SegmentedControl
	symbols  []SFSymbol
	selected int
	trigger  func(idx int)
}

type SFSymbol struct {
	normal string
	filled string
}

func getSegmentControl() *SegmentControl {
	// TODO: Segment cell display borderless
	segment := &SegmentControl{
		SegmentedControl: appkit.NewSegmentedControl(),
		symbols: []SFSymbol{
			{"personalhotspot.circle", "personalhotspot.circle.fill"},
			{"rectangle.stack", "rectangle.stack.fill"},
		},
	}
	segment.SetFocusRingType(appkit.FocusRingTypeNone)
	segment.SetAutoresizesSubviews(true)
	segment.SetSegmentStyle(appkit.SegmentStyleTexturedSquare)
	segment.SetSpringLoaded(true)
	segment.SetSegmentDistribution(appkit.SegmentDistributionFillProportionally)
	segment.SetAlignment(appkit.TextAlignmentCenter)
	segment.SetUserInterfaceLayoutDirection(appkit.UserInterfaceLayoutDirectionLeftToRight)
	segment.SetIgnoresMultiClick(true)
	segment.SetUsesSingleLineMode(true)
	segment.SetTranslatesAutoresizingMaskIntoConstraints(false)
	segment.SetTrackingMode(appkit.SegmentSwitchTrackingSelectOne)
	segment.SetSegmentCount(len(segment.symbols))
	segment.SetSelectedSegment(segment.selected)
	target, selector := segment.Clicked()
	segment.SetAction(selector)
	segment.SetTarget(target)
	for idx, item := range segment.symbols {
		if idx == segment.selected {
			segment.SetImageForSegment(symbolImage(item.filled), idx)
			continue
		}
		segment.SetImageForSegment(symbolImage(item.normal), idx)
	}
	return segment
}

func (s *SegmentControl) Clicked() (target action.Target, selector objc.Selector) {
	return action.Wrap(func(sender objc.Object) {
		selected := s.SelectedSegment()
		if selected == s.selected {
			return
		}

		s.SetImageForSegment(symbolImage(s.symbols[selected].filled), selected)
		s.SetImageForSegment(symbolImage(s.symbols[s.selected].normal), s.selected)
		s.selected = s.SelectedSegment()
		if s.trigger != nil {
			s.trigger(s.selected)
		}
	})
}

func getHorizontalLine(width float64) appkit.Box {
	line := appkit.NewBoxWithFrame(rectOf(0, 0, width, 1))
	line.SetTranslatesAutoresizingMaskIntoConstraints(false)
	line.SetBoxType(appkit.BoxCustom)
	line.SetBorderColor(appkit.Color_ColorWithSRGBRedGreenBlueAlpha(0, 0, 0, 0.1))
	layout.SetMaxHeight(line, 1)
	return line
}

func setConnectionsDataSource(outline appkit.OutlineView) {
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
