package components

import (
	"github.com/charliego3/datago/datasources"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type ConnectionsViewController struct {
	appkit.ViewController
}

func NewConnectionsViewController() ConnectionsViewController {
	controller := ConnectionsViewController{appkit.NewViewController()}
	controller.setContentView()
	return controller
}

func (c ConnectionsViewController) setContentView() {
	outlineView := c.getOutlineView()
	scrollView := c.getScrollView()
	clipView := appkit.NewClipView()

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
	// clipView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	clipView.AddSubview(outlineView)
	// scrollView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	scrollView.AddSubview(clipView)
	c.SetView(scrollView)
}

func (c ConnectionsViewController) getOutlineView() appkit.OutlineView {
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
	c.setDataSource(outlineView)
	return outlineView
}

func (c ConnectionsViewController) setDataSource(outlineView appkit.OutlineView) {
	datasource := &datasources.OutlineViewDatasource{}
	datasource.SetOutlineViewChildOfItem(func(outline appkit.OutlineView, index int, item objc.Object) objc.Object {
		return objc.ObjectFrom(foundation.String_StringWithString("Hahahahahahah").Ptr())
	})
	datasource.SetOutlineViewIsItemExpandable(func(outlineView appkit.OutlineView, item objc.Object) bool {
		return false
	})
	datasource.SetOutlineViewNumberOfChildrenOfItem(func(outlineView appkit.OutlineView, item objc.Object) int {
		return 1
	})
	datasource.SetOutlineViewObjectValueForTableColumnByItem(func(outlineView appkit.OutlineView, tableColumn appkit.TableColumn, item objc.Object) objc.Object {
		return objc.ObjectFrom(foundation.String_StringWithString("Hahahahahahah").Ptr())
	})
	po0 := objc.WrapAsProtocol("NSOutlineViewDataSource", datasource.Wrap())
	objc.SetAssociatedObject(outlineView, objc.AssociationKey("setDataSource"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](outlineView, objc.Sel("setDataSource:"), po0)
}

func (c ConnectionsViewController) getScrollView() appkit.ScrollView {
	scrollView := appkit.NewScrollView()
	scrollView.SetBorderType(appkit.NoBorder)
	scrollView.SetScrollerKnobStyle(appkit.ScrollerKnobStyleDefault)
	scrollView.SetScrollerStyle(appkit.ScrollerStyleOverlay)
	scrollView.SetFindBarPosition(appkit.ScrollViewFindBarPositionAboveContent)
	scrollView.SetDrawsBackground(false)
	return scrollView
}
