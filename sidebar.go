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
	label := appkit.NewLabel("Label1111")
	controller1 := appkit.NewViewController()
	controller1.SetView(label)
	// TODO: TabViewItems Ended

	tabView := appkit.NewTabView()
	tabView.SetTabViewType(appkit.NoTabsNoBorder)
	tabView.SetTabViewBorderType(appkit.TabViewBorderTypeNone)
	tabView.SetUserInterfaceLayoutDirection(appkit.UserInterfaceLayoutDirectionLeftToRight)
	tabView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	connectionsViewController := getConnectionsViewController()
	tabView.SetTabViewItems([]appkit.ITabViewItem{
		appkit.TabViewItem_TabViewItemWithViewController(connectionsViewController),
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

func getConnectionsViewController() appkit.ViewController {
	scrollView := appkit.NewScrollView()
	scrollView.SetBorderType(appkit.NoBorder)
	scrollView.SetScrollerKnobStyle(appkit.ScrollerKnobStyleDefault)
	scrollView.SetScrollerStyle(appkit.ScrollerStyleOverlay)
	scrollView.SetFindBarPosition(appkit.ScrollViewFindBarPositionAboveContent)
	scrollView.SetDrawsBackground(false)

	clipView := appkit.NewClipView()
	outlineView := appkit.NewOutlineView()
	outlineView.SetFloatsGroupRows(true)
	outlineView.SetAllowsColumnResizing(true)
	outlineView.SetStyle(appkit.TableViewStyleSourceList)
	outlineView.SetSelectionHighlightStyle(appkit.TableViewSelectionHighlightStyleSourceList)
	outlineView.SetGridStyleMask(appkit.TableViewGridNone)
	outlineView.SetUsesSingleLineMode(true)
	outlineView.SetColumnAutoresizingStyle(appkit.TableViewLastColumnOnlyAutoresizingStyle)
	outlineView.SetUsesAlternatingRowBackgroundColors(false)
	outlineView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	outlineView.SetAllowsTypeSelect(true)
	outlineView.SetAllowsEmptySelection(true)
	outlineView.SetIndentationMarkerFollowsCell(true)
	outlineView.SetIntercellSpacing(foundation.Size{Width: 3})

	column := appkit.NewTableColumnWithIdentifier("ConnectionsColumn")
	column.SetEditable(true)
	column.SetHeaderToolTip("tool")

	cellView := appkit.NewTableCellView()
	imageView := appkit.NewImageView()
	imageView.SetImage(appkit.Image_ImageWithSystemSymbolNameAccessibilityDescription("sidebar.leading", ""))
	cellView.SetImageView(imageView)
	cellView.SetTextField(appkit.NewTextField())

	outlineView.ColumnForView(cellView)
	outlineView.SetOutlineTableColumn(column)

	clipView.SetContentInsets(foundation.EdgeInsets{Top: 10})
	clipView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	scrollView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	scrollView.AddSubview(clipView)
	controller := appkit.NewViewController()
	controller.SetView(scrollView)
	//layout.PinEdgesToSuperView(scrollView, foundation.EdgeInsets{})
	//layoutActive(
	//	clipView.TopAnchor().ConstraintEqualToAnchor(scrollView.TopAnchor()),
	//	clipView.TrailingAnchor().ConstraintEqualToAnchor(scrollView.TrailingAnchor()),
	//	clipView.BottomAnchor().ConstraintEqualToAnchor(scrollView.BottomAnchor()),
	//	clipView.LeadingAnchor().ConstraintEqualToAnchor(scrollView.LeadingAnchor()),
	//	outlineView.TopAnchor().ConstraintEqualToAnchor(clipView.TopAnchor()),
	//	outlineView.TrailingAnchor().ConstraintEqualToAnchor(clipView.TrailingAnchor()),
	//	outlineView.BottomAnchor().ConstraintEqualToAnchor(clipView.BottomAnchor()),
	//	outlineView.LeadingAnchor().ConstraintEqualToAnchor(clipView.LeadingAnchor()),
	//)
	return controller
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
		symbols: []SFSymbol{
			{"rectangle.stack", "rectangle.stack.fill"},
			{"personalhotspot.circle", "personalhotspot.circle.fill"},
		},
	}

	appkit.Color_ControlAccentColor()
	target, selector := segment.Clicked()
	control := appkit.SegmentedControl_SegmentedControlWithImagesTrackingModeTargetAction(
		[]appkit.IImage{
			symbolImage(
				segment.symbols[0].filled,
				appkit.ImageSymbolConfiguration_ConfigurationWithHierarchicalColor(appkit.Color_ControlAccentColor()),
			),
			symbolImage(segment.symbols[1].normal),
		},
		appkit.SegmentSwitchTrackingSelectOne,
		target, selector,
	)

	segment.SegmentedControl = control
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
	segment.SetSelectedSegment(segment.selected)
	return segment
}

func (s *SegmentControl) Clicked() (target action.Target, selector objc.Selector) {
	return action.Wrap(func(sender objc.Object) {
		selected := s.SelectedSegment()
		if selected == s.selected {
			return
		}

		configuration := appkit.ImageSymbolConfiguration_ConfigurationWithHierarchicalColor(appkit.Color_ControlAccentColor())
		s.SetImageForSegment(symbolImage(s.symbols[selected].filled, configuration), selected)
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
	line.SetBorderColor(getDividerColor())
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
