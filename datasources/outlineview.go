package datasources

import (
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type OutlineViewDatasource struct {
	_OutlineViewSortDescriptorsDidChange                   func(outlineView appkit.OutlineView, oldDescriptors []foundation.SortDescriptor)
	_OutlineViewChildOfItem                                func(outline appkit.OutlineView, index int, item objc.Object) objc.Object
	_OutlineViewPersistentObjectForItem                    func(outlineView appkit.OutlineView, item objc.Object) objc.Object
	_OutlineViewValidateDropProposedItemProposedChildIndex func(outlineView appkit.OutlineView, info appkit.DraggingInfoObject, item objc.Object, index int) appkit.DragOperation
	_OutlineViewDraggingSessionWillBeginAtPointForItems    func(outlineView appkit.OutlineView, session appkit.DraggingSession, screenPoint foundation.Point, draggedItems []objc.Object)
	_OutlineViewAcceptDropItemChildIndex                   func(outlineView appkit.OutlineView, info appkit.DraggingInfoObject, item objc.Object, index int) bool
	_OutlineViewDraggingSessionEndedAtPointOperation       func(outlineView appkit.OutlineView, session appkit.DraggingSession, screenPoint foundation.Point, operation appkit.DragOperation)
	_OutlineViewSetObjectValueForTableColumnByItem         func(outlineView appkit.OutlineView, object objc.Object, tableColumn appkit.TableColumn, item objc.Object)
	_OutlineViewPasteboardWriterForItem                    func(outlineView appkit.OutlineView, item objc.Object) appkit.PasteboardWritingObject
	_OutlineViewNumberOfChildrenOfItem                     func(outlineView appkit.OutlineView, item objc.Object) int
	_OutlineViewObjectValueForTableColumnByItem            func(outlineView appkit.OutlineView, tableColumn appkit.TableColumn, item objc.Object) objc.Object
	_OutlineViewUpdateDraggingItemsForDrag                 func(outlineView appkit.OutlineView, draggingInfo appkit.DraggingInfoObject)
	_OutlineViewItemForPersistentObject                    func(outlineView appkit.OutlineView, object objc.Object) objc.Object
	_OutlineViewIsItemExpandable                           func(outlineView appkit.OutlineView, item objc.Object) bool
}

var _ appkit.POutlineViewDataSource = (*OutlineViewDatasource)(nil)

func (ds *OutlineViewDatasource) Wrap() appkit.POutlineViewDataSource {
	return appkit.POutlineViewDataSource(&OutlineViewDatasource{})
}

// optional
func (ds *OutlineViewDatasource) SetOutlineViewSortDescriptorsDidChange(f func(outlineView appkit.OutlineView, oldDescriptors []foundation.SortDescriptor)) {
	ds._OutlineViewSortDescriptorsDidChange = f
}

func (ds *OutlineViewDatasource) OutlineViewSortDescriptorsDidChange(outlineView appkit.OutlineView, oldDescriptors []foundation.SortDescriptor) {
	ds._OutlineViewSortDescriptorsDidChange(outlineView, oldDescriptors)
}

func (ds *OutlineViewDatasource) HasOutlineViewSortDescriptorsDidChange() bool {
	return ds._OutlineViewSortDescriptorsDidChange != nil
}

// optional
func (ds *OutlineViewDatasource) SetOutlineViewChildOfItem(f func(outline appkit.OutlineView, index int, item objc.Object) objc.Object) {
	ds._OutlineViewChildOfItem = f
}

func (ds *OutlineViewDatasource) OutlineViewChildOfItem(outlineView appkit.OutlineView, index int, item objc.Object) objc.Object {
	return ds._OutlineViewChildOfItem(outlineView, index, item)
}

func (ds *OutlineViewDatasource) HasOutlineViewChildOfItem() bool {
	return ds._OutlineViewChildOfItem != nil
}

// optional
func (ds *OutlineViewDatasource) SetOutlineViewPersistentObjectForItem(f func(outlineView appkit.OutlineView, item objc.Object) objc.Object) {
	ds._OutlineViewPersistentObjectForItem = f
}

func (ds *OutlineViewDatasource) OutlineViewPersistentObjectForItem(outlineView appkit.OutlineView, item objc.Object) objc.Object {
	return ds._OutlineViewPersistentObjectForItem(outlineView, item)
}

func (ds *OutlineViewDatasource) HasOutlineViewPersistentObjectForItem() bool {
	return ds._OutlineViewPersistentObjectForItem != nil
}

// optional
func (ds *OutlineViewDatasource) SetOutlineViewValidateDropProposedItemProposedChildIndex(f func(outlineView appkit.OutlineView, info appkit.DraggingInfoObject, item objc.Object, index int) appkit.DragOperation) {
	ds._OutlineViewValidateDropProposedItemProposedChildIndex = f
}

func (ds *OutlineViewDatasource) OutlineViewValidateDropProposedItemProposedChildIndex(outlineView appkit.OutlineView, info appkit.DraggingInfoObject, item objc.Object, index int) appkit.DragOperation {
	return ds._OutlineViewValidateDropProposedItemProposedChildIndex(outlineView, info, item, index)
}

func (ds *OutlineViewDatasource) HasOutlineViewValidateDropProposedItemProposedChildIndex() bool {
	return ds._OutlineViewValidateDropProposedItemProposedChildIndex != nil
}

// optional
func (ds *OutlineViewDatasource) SetOutlineViewDraggingSessionWillBeginAtPointForItems(f func(outlineView appkit.OutlineView, session appkit.DraggingSession, screenPoint foundation.Point, draggedItems []objc.Object)) {
	ds._OutlineViewDraggingSessionWillBeginAtPointForItems = f
}

func (ds *OutlineViewDatasource) OutlineViewDraggingSessionWillBeginAtPointForItems(outlineView appkit.OutlineView, session appkit.DraggingSession, screenPoint foundation.Point, draggedItems []objc.Object) {
	ds._OutlineViewDraggingSessionWillBeginAtPointForItems(outlineView, session, screenPoint, draggedItems)
}

func (ds *OutlineViewDatasource) HasOutlineViewDraggingSessionWillBeginAtPointForItems() bool {
	return ds._OutlineViewDraggingSessionWillBeginAtPointForItems != nil
}

// optional
func (ds *OutlineViewDatasource) SetOutlineViewAcceptDropItemChildIndex(f func(outlineView appkit.OutlineView, info appkit.DraggingInfoObject, item objc.Object, index int) bool) {
	ds._OutlineViewAcceptDropItemChildIndex = f
}

func (ds *OutlineViewDatasource) OutlineViewAcceptDropItemChildIndex(outlineView appkit.OutlineView, info appkit.DraggingInfoObject, item objc.Object, index int) bool {
	return ds._OutlineViewAcceptDropItemChildIndex(outlineView, info, item, index)
}

func (ds *OutlineViewDatasource) HasOutlineViewAcceptDropItemChildIndex() bool {
	return ds._OutlineViewAcceptDropItemChildIndex != nil
}

// optional
func (ds *OutlineViewDatasource) SetOutlineViewDraggingSessionEndedAtPointOperation(f func(outlineView appkit.OutlineView, session appkit.DraggingSession, screenPoint foundation.Point, operation appkit.DragOperation)) {
	ds._OutlineViewDraggingSessionEndedAtPointOperation = f
}

func (ds *OutlineViewDatasource) OutlineViewDraggingSessionEndedAtPointOperation(outlineView appkit.OutlineView, session appkit.DraggingSession, screenPoint foundation.Point, operation appkit.DragOperation) {
	ds._OutlineViewDraggingSessionEndedAtPointOperation(outlineView, session, screenPoint, operation)
}

func (ds *OutlineViewDatasource) HasOutlineViewDraggingSessionEndedAtPointOperation() bool {
	return ds._OutlineViewDraggingSessionEndedAtPointOperation != nil
}

// optional
func (ds *OutlineViewDatasource) SetOutlineViewSetObjectValueForTableColumnByItem(f func(outlineView appkit.OutlineView, object objc.Object, tableColumn appkit.TableColumn, item objc.Object)) {
	ds._OutlineViewSetObjectValueForTableColumnByItem = f
}

func (ds *OutlineViewDatasource) OutlineViewSetObjectValueForTableColumnByItem(outlineView appkit.OutlineView, object objc.Object, tableColumn appkit.TableColumn, item objc.Object) {
	ds._OutlineViewSetObjectValueForTableColumnByItem(outlineView, object, tableColumn, item)
}

func (ds *OutlineViewDatasource) HasOutlineViewSetObjectValueForTableColumnByItem() bool {
	return ds._OutlineViewSetObjectValueForTableColumnByItem != nil
}

// optional
func (ds *OutlineViewDatasource) SetOutlineViewPasteboardWriterForItem(f func(outlineView appkit.OutlineView, item objc.Object) appkit.PasteboardWritingObject) {
	ds._OutlineViewPasteboardWriterForItem = f
}

func (ds *OutlineViewDatasource) OutlineViewPasteboardWriterForItem(outlineView appkit.OutlineView, item objc.Object) appkit.PasteboardWritingObject {
	return ds._OutlineViewPasteboardWriterForItem(outlineView, item)
}

func (ds *OutlineViewDatasource) HasOutlineViewPasteboardWriterForItem() bool {
	return ds._OutlineViewPasteboardWriterForItem != nil
}

// optional
func (ds *OutlineViewDatasource) SetOutlineViewNumberOfChildrenOfItem(f func(outlineView appkit.OutlineView, item objc.Object) int) {
	ds._OutlineViewNumberOfChildrenOfItem = f
}

func (ds *OutlineViewDatasource) OutlineViewNumberOfChildrenOfItem(outlineView appkit.OutlineView, item objc.Object) int {
	return ds._OutlineViewNumberOfChildrenOfItem(outlineView, item)
}

func (ds *OutlineViewDatasource) HasOutlineViewNumberOfChildrenOfItem() bool {
	return ds._OutlineViewNumberOfChildrenOfItem != nil
}

// optional
func (ds *OutlineViewDatasource) SetOutlineViewObjectValueForTableColumnByItem(f func(outlineView appkit.OutlineView, tableColumn appkit.TableColumn, item objc.Object) objc.Object) {
	ds._OutlineViewObjectValueForTableColumnByItem = f
}

func (ds *OutlineViewDatasource) OutlineViewObjectValueForTableColumnByItem(outlineView appkit.OutlineView, tableColumn appkit.TableColumn, item objc.Object) objc.Object {
	return ds._OutlineViewObjectValueForTableColumnByItem(outlineView, tableColumn, item)
}

func (ds *OutlineViewDatasource) HasOutlineViewObjectValueForTableColumnByItem() bool {
	return ds._OutlineViewObjectValueForTableColumnByItem != nil
}

// optional
func (ds *OutlineViewDatasource) SetOutlineViewUpdateDraggingItemsForDrag(f func(outlineView appkit.OutlineView, draggingInfo appkit.DraggingInfoObject)) {
	ds._OutlineViewUpdateDraggingItemsForDrag = f
}

func (ds *OutlineViewDatasource) OutlineViewUpdateDraggingItemsForDrag(outlineView appkit.OutlineView, draggingInfo appkit.DraggingInfoObject) {
	ds._OutlineViewUpdateDraggingItemsForDrag(outlineView, draggingInfo)
}

func (ds *OutlineViewDatasource) HasOutlineViewUpdateDraggingItemsForDrag() bool {
	return ds._OutlineViewUpdateDraggingItemsForDrag != nil
}

// optional
func (ds *OutlineViewDatasource) SetOutlineViewItemForPersistentObject(f func(outlineView appkit.OutlineView, object objc.Object) objc.Object) {
	ds._OutlineViewItemForPersistentObject = f
}

func (ds *OutlineViewDatasource) OutlineViewItemForPersistentObject(outlineView appkit.OutlineView, object objc.Object) objc.Object {
	return ds._OutlineViewItemForPersistentObject(outlineView, object)
}

func (ds *OutlineViewDatasource) HasOutlineViewItemForPersistentObject() bool {
	return ds._OutlineViewItemForPersistentObject != nil
}

// optional
func (ds *OutlineViewDatasource) SetOutlineViewIsItemExpandable(f func(outlineView appkit.OutlineView, item objc.Object) bool) {
	ds._OutlineViewIsItemExpandable = f
}

func (ds *OutlineViewDatasource) OutlineViewIsItemExpandable(outlineView appkit.OutlineView, item objc.Object) bool {
	return ds._OutlineViewIsItemExpandable(outlineView, item)
}

func (ds *OutlineViewDatasource) HasOutlineViewIsItemExpandable() bool {
	return ds._OutlineViewIsItemExpandable != nil
}
