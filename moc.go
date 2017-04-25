package main

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type MocLabel_ITF interface {
	widgets.QLabel_ITF
	MocLabel_PTR() *MocLabel
}

func (ptr *MocLabel) MocLabel_PTR() *MocLabel {
	return ptr
}

func (ptr *MocLabel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QLabel_PTR().Pointer()
	}
	return nil
}

func (ptr *MocLabel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QLabel_PTR().SetPointer(p)
	}
}

func PointerFromMocLabel(ptr MocLabel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.MocLabel_PTR().Pointer()
	}
	return nil
}

func NewMocLabelFromPointer(ptr unsafe.Pointer) *MocLabel {
	var n = new(MocLabel)
	n.SetPointer(ptr)
	return n
}

//export callbackMocLabel_UpdateLabel
func callbackMocLabel_UpdateLabel(ptr unsafe.Pointer, v0 C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "updateLabel"); signal != nil {
		signal.(func(string))(cGoUnpackString(v0))
	}

}

func (ptr *MocLabel) ConnectUpdateLabel(f func(v0 string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "updateLabel") {
			C.MocLabel_ConnectUpdateLabel(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "updateLabel"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "updateLabel", func(v0 string) {
				signal.(func(string))(v0)
				f(v0)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "updateLabel", f)
		}
	}
}

func (ptr *MocLabel) DisconnectUpdateLabel() {
	if ptr.Pointer() != nil {
		C.MocLabel_DisconnectUpdateLabel(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "updateLabel")
	}
}

func (ptr *MocLabel) UpdateLabel(v0 string) {
	if ptr.Pointer() != nil {
		var v0C *C.char
		if v0 != "" {
			v0C = C.CString(v0)
			defer C.free(unsafe.Pointer(v0C))
		}
		C.MocLabel_UpdateLabel(ptr.Pointer(), v0C)
	}
}

func MocLabel_QRegisterMetaType() int {
	return int(int32(C.MocLabel_MocLabel_QRegisterMetaType()))
}

func (ptr *MocLabel) QRegisterMetaType() int {
	return int(int32(C.MocLabel_MocLabel_QRegisterMetaType()))
}

func MocLabel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.MocLabel_MocLabel_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *MocLabel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.MocLabel_MocLabel_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func MocLabel_QmlRegisterType() int {
	return int(int32(C.MocLabel_MocLabel_QmlRegisterType()))
}

func (ptr *MocLabel) QmlRegisterType() int {
	return int(int32(C.MocLabel_MocLabel_QmlRegisterType()))
}

func (ptr *MocLabel) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.MocLabel___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MocLabel) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *MocLabel) __addActions_actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.MocLabel___addActions_actions_newList(ptr.Pointer()))
}

func (ptr *MocLabel) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.MocLabel___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MocLabel) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *MocLabel) __insertActions_actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.MocLabel___insertActions_actions_newList(ptr.Pointer()))
}

func (ptr *MocLabel) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.MocLabel___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MocLabel) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *MocLabel) __actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.MocLabel___actions_newList(ptr.Pointer()))
}

func (ptr *MocLabel) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.MocLabel___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *MocLabel) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *MocLabel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.MocLabel___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *MocLabel) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.MocLabel___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MocLabel) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *MocLabel) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.MocLabel___findChildren_newList2(ptr.Pointer()))
}

func (ptr *MocLabel) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.MocLabel___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MocLabel) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *MocLabel) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.MocLabel___findChildren_newList3(ptr.Pointer()))
}

func (ptr *MocLabel) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.MocLabel___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MocLabel) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *MocLabel) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.MocLabel___findChildren_newList(ptr.Pointer()))
}

func (ptr *MocLabel) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.MocLabel___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MocLabel) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *MocLabel) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.MocLabel___children_newList(ptr.Pointer()))
}

func NewMocLabel(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *MocLabel {
	var tmpValue = NewMocLabelFromPointer(C.MocLabel_NewMocLabel(widgets.PointerFromQWidget(parent), C.longlong(fo)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewMocLabel2(text string, parent widgets.QWidget_ITF, fo core.Qt__WindowType) *MocLabel {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	var tmpValue = NewMocLabelFromPointer(C.MocLabel_NewMocLabel2(textC, widgets.PointerFromQWidget(parent), C.longlong(fo)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *MocLabel) DestroyMocLabel() {
	if ptr.Pointer() != nil {
		C.MocLabel_DestroyMocLabel(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackMocLabel_Event
func callbackMocLabel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMocLabelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *MocLabel) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.MocLabel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackMocLabel_FocusNextPrevChild
func callbackMocLabel_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMocLabelFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *MocLabel) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return C.MocLabel_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackMocLabel_ChangeEvent
func callbackMocLabel_ChangeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewMocLabelFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *MocLabel) ChangeEventDefault(ev core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

//export callbackMocLabel_Clear
func callbackMocLabel_Clear(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "clear"); signal != nil {
		signal.(func())()
	} else {
		NewMocLabelFromPointer(ptr).ClearDefault()
	}
}

func (ptr *MocLabel) ClearDefault() {
	if ptr.Pointer() != nil {
		C.MocLabel_ClearDefault(ptr.Pointer())
	}
}

//export callbackMocLabel_ContextMenuEvent
func callbackMocLabel_ContextMenuEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(ev))
	} else {
		NewMocLabelFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(ev))
	}
}

func (ptr *MocLabel) ContextMenuEventDefault(ev gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(ev))
	}
}

//export callbackMocLabel_FocusInEvent
func callbackMocLabel_FocusInEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewMocLabelFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *MocLabel) FocusInEventDefault(ev gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackMocLabel_FocusOutEvent
func callbackMocLabel_FocusOutEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewMocLabelFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *MocLabel) FocusOutEventDefault(ev gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackMocLabel_KeyPressEvent
func callbackMocLabel_KeyPressEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(ev))
	} else {
		NewMocLabelFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(ev))
	}
}

func (ptr *MocLabel) KeyPressEventDefault(ev gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

//export callbackMocLabel_LinkActivated
func callbackMocLabel_LinkActivated(ptr unsafe.Pointer, link C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "linkActivated"); signal != nil {
		signal.(func(string))(cGoUnpackString(link))
	}

}

//export callbackMocLabel_LinkHovered
func callbackMocLabel_LinkHovered(ptr unsafe.Pointer, link C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "linkHovered"); signal != nil {
		signal.(func(string))(cGoUnpackString(link))
	}

}

//export callbackMocLabel_MouseMoveEvent
func callbackMocLabel_MouseMoveEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
	} else {
		NewMocLabelFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *MocLabel) MouseMoveEventDefault(ev gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackMocLabel_MousePressEvent
func callbackMocLabel_MousePressEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
	} else {
		NewMocLabelFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *MocLabel) MousePressEventDefault(ev gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackMocLabel_MouseReleaseEvent
func callbackMocLabel_MouseReleaseEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
	} else {
		NewMocLabelFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *MocLabel) MouseReleaseEventDefault(ev gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackMocLabel_PaintEvent
func callbackMocLabel_PaintEvent(ptr unsafe.Pointer, vqp unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(vqp))
	} else {
		NewMocLabelFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(vqp))
	}
}

func (ptr *MocLabel) PaintEventDefault(vqp gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(vqp))
	}
}

//export callbackMocLabel_SetMovie
func callbackMocLabel_SetMovie(ptr unsafe.Pointer, movie unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setMovie"); signal != nil {
		signal.(func(*gui.QMovie))(gui.NewQMovieFromPointer(movie))
	} else {
		NewMocLabelFromPointer(ptr).SetMovieDefault(gui.NewQMovieFromPointer(movie))
	}
}

func (ptr *MocLabel) SetMovieDefault(movie gui.QMovie_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_SetMovieDefault(ptr.Pointer(), gui.PointerFromQMovie(movie))
	}
}

//export callbackMocLabel_SetNum2
func callbackMocLabel_SetNum2(ptr unsafe.Pointer, num C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setNum2"); signal != nil {
		signal.(func(float64))(float64(num))
	} else {
		NewMocLabelFromPointer(ptr).SetNum2Default(float64(num))
	}
}

func (ptr *MocLabel) SetNum2Default(num float64) {
	if ptr.Pointer() != nil {
		C.MocLabel_SetNum2Default(ptr.Pointer(), C.double(num))
	}
}

//export callbackMocLabel_SetNum
func callbackMocLabel_SetNum(ptr unsafe.Pointer, num C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setNum"); signal != nil {
		signal.(func(int))(int(int32(num)))
	} else {
		NewMocLabelFromPointer(ptr).SetNumDefault(int(int32(num)))
	}
}

func (ptr *MocLabel) SetNumDefault(num int) {
	if ptr.Pointer() != nil {
		C.MocLabel_SetNumDefault(ptr.Pointer(), C.int(int32(num)))
	}
}

//export callbackMocLabel_SetPicture
func callbackMocLabel_SetPicture(ptr unsafe.Pointer, picture unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setPicture"); signal != nil {
		signal.(func(*gui.QPicture))(gui.NewQPictureFromPointer(picture))
	} else {
		NewMocLabelFromPointer(ptr).SetPictureDefault(gui.NewQPictureFromPointer(picture))
	}
}

func (ptr *MocLabel) SetPictureDefault(picture gui.QPicture_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_SetPictureDefault(ptr.Pointer(), gui.PointerFromQPicture(picture))
	}
}

//export callbackMocLabel_SetPixmap
func callbackMocLabel_SetPixmap(ptr unsafe.Pointer, vqp unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setPixmap"); signal != nil {
		signal.(func(*gui.QPixmap))(gui.NewQPixmapFromPointer(vqp))
	} else {
		NewMocLabelFromPointer(ptr).SetPixmapDefault(gui.NewQPixmapFromPointer(vqp))
	}
}

func (ptr *MocLabel) SetPixmapDefault(vqp gui.QPixmap_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_SetPixmapDefault(ptr.Pointer(), gui.PointerFromQPixmap(vqp))
	}
}

//export callbackMocLabel_SetText
func callbackMocLabel_SetText(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setText"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewMocLabelFromPointer(ptr).SetTextDefault(cGoUnpackString(vqs))
	}
}

func (ptr *MocLabel) SetTextDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.MocLabel_SetTextDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackMocLabel_MinimumSizeHint
func callbackMocLabel_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewMocLabelFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *MocLabel) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.MocLabel_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackMocLabel_SizeHint
func callbackMocLabel_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewMocLabelFromPointer(ptr).SizeHintDefault())
}

func (ptr *MocLabel) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.MocLabel_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackMocLabel_HeightForWidth
func callbackMocLabel_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewMocLabelFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *MocLabel) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MocLabel_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

<<<<<<< HEAD
//export callbackMocLabel_Close
func callbackMocLabel_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewMocLabelFromPointer(ptr).CloseDefault())))
}

func (ptr *MocLabel) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return C.MocLabel_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackMocLabel_NativeEvent
func callbackMocLabel_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result C.long) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QByteArray, unsafe.Pointer, int) bool)(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMocLabelFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
}

func (ptr *MocLabel) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.MocLabel_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackMocLabel_ActionEvent
func callbackMocLabel_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewMocLabelFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *MocLabel) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackMocLabel_CloseEvent
func callbackMocLabel_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewMocLabelFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *MocLabel) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackMocLabel_CustomContextMenuRequested
func callbackMocLabel_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "customContextMenuRequested"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(pos))
	}

}

//export callbackMocLabel_DragEnterEvent
func callbackMocLabel_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewMocLabelFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *MocLabel) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackMocLabel_DragLeaveEvent
func callbackMocLabel_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewMocLabelFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *MocLabel) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackMocLabel_DragMoveEvent
func callbackMocLabel_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewMocLabelFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *MocLabel) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackMocLabel_DropEvent
func callbackMocLabel_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewMocLabelFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *MocLabel) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackMocLabel_EnterEvent
func callbackMocLabel_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewMocLabelFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *MocLabel) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackMocLabel_Hide
func callbackMocLabel_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "hide"); signal != nil {
		signal.(func())()
	} else {
		NewMocLabelFromPointer(ptr).HideDefault()
	}
}

func (ptr *MocLabel) HideDefault() {
	if ptr.Pointer() != nil {
		C.MocLabel_HideDefault(ptr.Pointer())
	}
}

//export callbackMocLabel_HideEvent
func callbackMocLabel_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewMocLabelFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *MocLabel) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackMocLabel_InputMethodEvent
func callbackMocLabel_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewMocLabelFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *MocLabel) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackMocLabel_KeyReleaseEvent
func callbackMocLabel_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewMocLabelFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *MocLabel) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackMocLabel_LeaveEvent
func callbackMocLabel_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewMocLabelFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *MocLabel) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackMocLabel_Lower
func callbackMocLabel_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "lower"); signal != nil {
		signal.(func())()
	} else {
		NewMocLabelFromPointer(ptr).LowerDefault()
	}
}

func (ptr *MocLabel) LowerDefault() {
	if ptr.Pointer() != nil {
		C.MocLabel_LowerDefault(ptr.Pointer())
	}
}

//export callbackMocLabel_MouseDoubleClickEvent
func callbackMocLabel_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewMocLabelFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MocLabel) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMocLabel_MoveEvent
func callbackMocLabel_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewMocLabelFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *MocLabel) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackMocLabel_Raise
func callbackMocLabel_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "raise"); signal != nil {
		signal.(func())()
	} else {
		NewMocLabelFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *MocLabel) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.MocLabel_RaiseDefault(ptr.Pointer())
	}
}

//export callbackMocLabel_Repaint
func callbackMocLabel_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewMocLabelFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *MocLabel) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.MocLabel_RepaintDefault(ptr.Pointer())
	}
}

//export callbackMocLabel_ResizeEvent
func callbackMocLabel_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewMocLabelFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *MocLabel) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackMocLabel_SetDisabled
func callbackMocLabel_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewMocLabelFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *MocLabel) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.MocLabel_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackMocLabel_SetEnabled
func callbackMocLabel_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewMocLabelFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *MocLabel) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.MocLabel_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackMocLabel_SetFocus2
func callbackMocLabel_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewMocLabelFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *MocLabel) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.MocLabel_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackMocLabel_SetHidden
func callbackMocLabel_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewMocLabelFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *MocLabel) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.MocLabel_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackMocLabel_SetStyleSheet
func callbackMocLabel_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewMocLabelFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *MocLabel) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.MocLabel_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackMocLabel_SetVisible
func callbackMocLabel_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewMocLabelFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *MocLabel) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.MocLabel_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackMocLabel_SetWindowModified
func callbackMocLabel_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewMocLabelFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *MocLabel) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.MocLabel_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackMocLabel_SetWindowTitle
func callbackMocLabel_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewMocLabelFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *MocLabel) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.MocLabel_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackMocLabel_Show
func callbackMocLabel_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "show"); signal != nil {
		signal.(func())()
	} else {
		NewMocLabelFromPointer(ptr).ShowDefault()
	}
}

func (ptr *MocLabel) ShowDefault() {
	if ptr.Pointer() != nil {
		C.MocLabel_ShowDefault(ptr.Pointer())
	}
}

//export callbackMocLabel_ShowEvent
func callbackMocLabel_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewMocLabelFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *MocLabel) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackMocLabel_ShowFullScreen
func callbackMocLabel_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewMocLabelFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *MocLabel) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.MocLabel_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackMocLabel_ShowMaximized
func callbackMocLabel_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewMocLabelFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *MocLabel) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.MocLabel_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackMocLabel_ShowMinimized
func callbackMocLabel_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewMocLabelFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *MocLabel) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.MocLabel_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackMocLabel_ShowNormal
func callbackMocLabel_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewMocLabelFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *MocLabel) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.MocLabel_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackMocLabel_TabletEvent
func callbackMocLabel_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewMocLabelFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *MocLabel) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackMocLabel_Update
func callbackMocLabel_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "update"); signal != nil {
		signal.(func())()
	} else {
		NewMocLabelFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *MocLabel) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.MocLabel_UpdateDefault(ptr.Pointer())
	}
}

//export callbackMocLabel_UpdateMicroFocus
func callbackMocLabel_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewMocLabelFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *MocLabel) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.MocLabel_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackMocLabel_WheelEvent
func callbackMocLabel_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewMocLabelFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *MocLabel) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackMocLabel_WindowIconChanged
func callbackMocLabel_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "windowIconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackMocLabel_WindowTitleChanged
func callbackMocLabel_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackMocLabel_PaintEngine
func callbackMocLabel_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewMocLabelFromPointer(ptr).PaintEngineDefault())
}

func (ptr *MocLabel) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.MocLabel_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackMocLabel_InputMethodQuery
func callbackMocLabel_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewMocLabelFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *MocLabel) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.MocLabel_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackMocLabel_HasHeightForWidth
func callbackMocLabel_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewMocLabelFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *MocLabel) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.MocLabel_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackMocLabel_Metric
func callbackMocLabel_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewMocLabelFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *MocLabel) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MocLabel_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackMocLabel_EventFilter
func callbackMocLabel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMocLabelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *MocLabel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.MocLabel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackMocLabel_ChildEvent
func callbackMocLabel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewMocLabelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *MocLabel) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackMocLabel_ConnectNotify
func callbackMocLabel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewMocLabelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *MocLabel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackMocLabel_CustomEvent
func callbackMocLabel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewMocLabelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *MocLabel) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackMocLabel_DeleteLater
func callbackMocLabel_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewMocLabelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *MocLabel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.MocLabel_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackMocLabel_Destroyed
func callbackMocLabel_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackMocLabel_DisconnectNotify
func callbackMocLabel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewMocLabelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *MocLabel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackMocLabel_ObjectNameChanged
func callbackMocLabel_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackMocLabel_TimerEvent
func callbackMocLabel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewMocLabelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *MocLabel) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MocLabel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type App_ITF interface {
	widgets.QApplication_ITF
	App_PTR() *App
}

func (ptr *App) App_PTR() *App {
	return ptr
}

func (ptr *App) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QApplication_PTR().Pointer()
	}
	return nil
}

func (ptr *App) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QApplication_PTR().SetPointer(p)
	}
}

func PointerFromApp(ptr App_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.App_PTR().Pointer()
	}
	return nil
}

func NewAppFromPointer(ptr unsafe.Pointer) *App {
	var n = new(App)
	n.SetPointer(ptr)
	return n
}
func App_QRegisterMetaType() int {
	return int(int32(C.App_App_QRegisterMetaType()))
}

func (ptr *App) QRegisterMetaType() int {
	return int(int32(C.App_App_QRegisterMetaType()))
}

func App_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.App_App_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *App) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.App_App_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func App_QmlRegisterType() int {
	return int(int32(C.App_App_QmlRegisterType()))
}

func (ptr *App) QmlRegisterType() int {
	return int(int32(C.App_App_QmlRegisterType()))
}

func (ptr *App) __allWidgets_atList(i int) *widgets.QWidget {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.App___allWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *App) __allWidgets_setList(i widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.App___allWidgets_setList(ptr.Pointer(), widgets.PointerFromQWidget(i))
	}
}

func (ptr *App) __allWidgets_newList() unsafe.Pointer {
	return unsafe.Pointer(C.App___allWidgets_newList(ptr.Pointer()))
}

func (ptr *App) __topLevelWidgets_atList(i int) *widgets.QWidget {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.App___topLevelWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *App) __topLevelWidgets_setList(i widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.App___topLevelWidgets_setList(ptr.Pointer(), widgets.PointerFromQWidget(i))
	}
}

func (ptr *App) __topLevelWidgets_newList() unsafe.Pointer {
	return unsafe.Pointer(C.App___topLevelWidgets_newList(ptr.Pointer()))
}

func (ptr *App) __screens_atList(i int) *gui.QScreen {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQScreenFromPointer(C.App___screens_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *App) __screens_setList(i gui.QScreen_ITF) {
	if ptr.Pointer() != nil {
		C.App___screens_setList(ptr.Pointer(), gui.PointerFromQScreen(i))
	}
}

func (ptr *App) __screens_newList() unsafe.Pointer {
	return unsafe.Pointer(C.App___screens_newList(ptr.Pointer()))
}

func (ptr *App) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.App___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *App) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.App___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *App) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.App___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *App) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.App___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *App) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.App___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *App) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.App___findChildren_newList2(ptr.Pointer()))
}

func (ptr *App) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.App___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *App) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.App___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *App) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.App___findChildren_newList3(ptr.Pointer()))
}

func (ptr *App) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.App___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *App) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.App___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *App) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.App___findChildren_newList(ptr.Pointer()))
}

func (ptr *App) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.App___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *App) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.App___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *App) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.App___children_newList(ptr.Pointer()))
}

func NewApp(argc int, argv []string) *App {
	var argvC = C.CString(strings.Join(argv, "|"))
	defer C.free(unsafe.Pointer(argvC))
	var tmpValue = NewAppFromPointer(C.App_NewApp(C.int(int32(argc)), argvC))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackApp_DestroyApp
func callbackApp_DestroyApp(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~App"); signal != nil {
		signal.(func())()
	} else {
		NewAppFromPointer(ptr).DestroyAppDefault()
	}
}

func (ptr *App) ConnectDestroyApp(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~App"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~App", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~App", f)
		}
	}
}

func (ptr *App) DisconnectDestroyApp() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~App")
	}
}

func (ptr *App) DestroyApp() {
	if ptr.Pointer() != nil {
		C.App_DestroyApp(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *App) DestroyAppDefault() {
	if ptr.Pointer() != nil {
		C.App_DestroyAppDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackApp_Event
func callbackApp_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAppFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *App) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.App_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackApp_AboutQt
func callbackApp_AboutQt(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "aboutQt"); signal != nil {
		signal.(func())()
	} else {
		NewAppFromPointer(ptr).AboutQtDefault()
	}
}

func (ptr *App) AboutQtDefault() {
	if ptr.Pointer() != nil {
		C.App_AboutQtDefault(ptr.Pointer())
	}
}

//export callbackApp_CloseAllWindows
func callbackApp_CloseAllWindows(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "closeAllWindows"); signal != nil {
		signal.(func())()
	} else {
		NewAppFromPointer(ptr).CloseAllWindowsDefault()
	}
}

func (ptr *App) CloseAllWindowsDefault() {
	if ptr.Pointer() != nil {
		C.App_CloseAllWindowsDefault(ptr.Pointer())
	}
}

//export callbackApp_FocusChanged
func callbackApp_FocusChanged(ptr unsafe.Pointer, old unsafe.Pointer, now unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "focusChanged"); signal != nil {
		signal.(func(*widgets.QWidget, *widgets.QWidget))(widgets.NewQWidgetFromPointer(old), widgets.NewQWidgetFromPointer(now))
	}

}

//export callbackApp_SetAutoSipEnabled
func callbackApp_SetAutoSipEnabled(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setAutoSipEnabled"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	} else {
		NewAppFromPointer(ptr).SetAutoSipEnabledDefault(int8(enabled) != 0)
	}
}

func (ptr *App) SetAutoSipEnabledDefault(enabled bool) {
	if ptr.Pointer() != nil {
		C.App_SetAutoSipEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackApp_SetStyleSheet
func callbackApp_SetStyleSheet(ptr unsafe.Pointer, sheet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(sheet))
	} else {
		NewAppFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(sheet))
	}
}

func (ptr *App) SetStyleSheetDefault(sheet string) {
	if ptr.Pointer() != nil {
		var sheetC *C.char
		if sheet != "" {
			sheetC = C.CString(sheet)
			defer C.free(unsafe.Pointer(sheetC))
		}
		C.App_SetStyleSheetDefault(ptr.Pointer(), sheetC)
	}
}

//export callbackApp_AutoSipEnabled
func callbackApp_AutoSipEnabled(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "autoSipEnabled"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewAppFromPointer(ptr).AutoSipEnabledDefault())))
}

func (ptr *App) AutoSipEnabledDefault() bool {
	if ptr.Pointer() != nil {
		return C.App_AutoSipEnabledDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackApp_ApplicationDisplayNameChanged
func callbackApp_ApplicationDisplayNameChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "applicationDisplayNameChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackApp_ApplicationStateChanged
func callbackApp_ApplicationStateChanged(ptr unsafe.Pointer, state C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "applicationStateChanged"); signal != nil {
		signal.(func(core.Qt__ApplicationState))(core.Qt__ApplicationState(state))
	}

}

//export callbackApp_CommitDataRequest
func callbackApp_CommitDataRequest(ptr unsafe.Pointer, manager unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "commitDataRequest"); signal != nil {
		signal.(func(*gui.QSessionManager))(gui.NewQSessionManagerFromPointer(manager))
	}

}

//export callbackApp_FocusObjectChanged
func callbackApp_FocusObjectChanged(ptr unsafe.Pointer, focusObject unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "focusObjectChanged"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(focusObject))
	}

}

//export callbackApp_FocusWindowChanged
func callbackApp_FocusWindowChanged(ptr unsafe.Pointer, focusWindow unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "focusWindowChanged"); signal != nil {
		signal.(func(*gui.QWindow))(gui.NewQWindowFromPointer(focusWindow))
	}

}

//export callbackApp_FontDatabaseChanged
func callbackApp_FontDatabaseChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "fontDatabaseChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackApp_LastWindowClosed
func callbackApp_LastWindowClosed(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "lastWindowClosed"); signal != nil {
		signal.(func())()
	}

}

//export callbackApp_LayoutDirectionChanged
func callbackApp_LayoutDirectionChanged(ptr unsafe.Pointer, direction C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "layoutDirectionChanged"); signal != nil {
		signal.(func(core.Qt__LayoutDirection))(core.Qt__LayoutDirection(direction))
	}

}

//export callbackApp_PaletteChanged
func callbackApp_PaletteChanged(ptr unsafe.Pointer, palette unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "paletteChanged"); signal != nil {
		signal.(func(*gui.QPalette))(gui.NewQPaletteFromPointer(palette))
	}

}

//export callbackApp_PrimaryScreenChanged
func callbackApp_PrimaryScreenChanged(ptr unsafe.Pointer, screen unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "primaryScreenChanged"); signal != nil {
		signal.(func(*gui.QScreen))(gui.NewQScreenFromPointer(screen))
	}

}

//export callbackApp_SaveStateRequest
func callbackApp_SaveStateRequest(ptr unsafe.Pointer, manager unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "saveStateRequest"); signal != nil {
		signal.(func(*gui.QSessionManager))(gui.NewQSessionManagerFromPointer(manager))
	}

}

//export callbackApp_ScreenAdded
func callbackApp_ScreenAdded(ptr unsafe.Pointer, screen unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "screenAdded"); signal != nil {
		signal.(func(*gui.QScreen))(gui.NewQScreenFromPointer(screen))
	}

}

//export callbackApp_ScreenRemoved
func callbackApp_ScreenRemoved(ptr unsafe.Pointer, screen unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "screenRemoved"); signal != nil {
		signal.(func(*gui.QScreen))(gui.NewQScreenFromPointer(screen))
	}

}

//export callbackApp_AboutToQuit
func callbackApp_AboutToQuit(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "aboutToQuit"); signal != nil {
		signal.(func())()
	}

}

//export callbackApp_Quit
func callbackApp_Quit(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "quit"); signal != nil {
		signal.(func())()
	} else {
		NewAppFromPointer(ptr).QuitDefault()
	}
}

func (ptr *App) QuitDefault() {
	if ptr.Pointer() != nil {
		C.App_QuitDefault(ptr.Pointer())
	}
}

//export callbackApp_EventFilter
func callbackApp_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAppFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *App) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.App_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackApp_ChildEvent
func callbackApp_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewAppFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *App) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.App_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackApp_ConnectNotify
func callbackApp_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewAppFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *App) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.App_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackApp_CustomEvent
func callbackApp_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewAppFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *App) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.App_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackApp_DeleteLater
func callbackApp_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewAppFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *App) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.App_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackApp_Destroyed
func callbackApp_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackApp_DisconnectNotify
func callbackApp_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewAppFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *App) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.App_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackApp_ObjectNameChanged
func callbackApp_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackApp_TimerEvent
func callbackApp_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewAppFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *App) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.App_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type MWidget_ITF interface {
	widgets.QWidget_ITF
	MWidget_PTR() *MWidget
}

func (ptr *MWidget) MWidget_PTR() *MWidget {
	return ptr
}

func (ptr *MWidget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *MWidget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromMWidget(ptr MWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.MWidget_PTR().Pointer()
	}
	return nil
}

func NewMWidgetFromPointer(ptr unsafe.Pointer) *MWidget {
	var n = new(MWidget)
	n.SetPointer(ptr)
	return n
}
func MWidget_QRegisterMetaType() int {
	return int(int32(C.MWidget_MWidget_QRegisterMetaType()))
}

func (ptr *MWidget) QRegisterMetaType() int {
	return int(int32(C.MWidget_MWidget_QRegisterMetaType()))
}

func MWidget_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.MWidget_MWidget_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *MWidget) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.MWidget_MWidget_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func MWidget_QmlRegisterType() int {
	return int(int32(C.MWidget_MWidget_QmlRegisterType()))
}

func (ptr *MWidget) QmlRegisterType() int {
	return int(int32(C.MWidget_MWidget_QmlRegisterType()))
}

func (ptr *MWidget) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.MWidget___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MWidget) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *MWidget) __addActions_actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.MWidget___addActions_actions_newList(ptr.Pointer()))
}

func (ptr *MWidget) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.MWidget___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MWidget) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *MWidget) __insertActions_actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.MWidget___insertActions_actions_newList(ptr.Pointer()))
}

func (ptr *MWidget) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.MWidget___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MWidget) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *MWidget) __actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.MWidget___actions_newList(ptr.Pointer()))
}

func (ptr *MWidget) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.MWidget___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *MWidget) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *MWidget) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.MWidget___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *MWidget) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.MWidget___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MWidget) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *MWidget) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.MWidget___findChildren_newList2(ptr.Pointer()))
}

func (ptr *MWidget) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.MWidget___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MWidget) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *MWidget) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.MWidget___findChildren_newList3(ptr.Pointer()))
}

func (ptr *MWidget) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.MWidget___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MWidget) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *MWidget) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.MWidget___findChildren_newList(ptr.Pointer()))
}

func (ptr *MWidget) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.MWidget___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MWidget) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *MWidget) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.MWidget___children_newList(ptr.Pointer()))
}

func NewMWidget(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *MWidget {
	var tmpValue = NewMWidgetFromPointer(C.MWidget_NewMWidget(widgets.PointerFromQWidget(parent), C.longlong(fo)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *MWidget) DestroyMWidget() {
	if ptr.Pointer() != nil {
		C.MWidget_DestroyMWidget(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackMWidget_Close
func callbackMWidget_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewMWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *MWidget) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return C.MWidget_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackMWidget_Event
func callbackMWidget_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMWidgetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *MWidget) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.MWidget_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackMWidget_FocusNextPrevChild
func callbackMWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *MWidget) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return C.MWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackMWidget_NativeEvent
func callbackMWidget_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result C.long) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QByteArray, unsafe.Pointer, int) bool)(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMWidgetFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
}

func (ptr *MWidget) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.MWidget_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackMWidget_ActionEvent
func callbackMWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *MWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackMWidget_ChangeEvent
func callbackMWidget_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *MWidget) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackMWidget_CloseEvent
func callbackMWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *MWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackMWidget_ContextMenuEvent
func callbackMWidget_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *MWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackMWidget_CustomContextMenuRequested
func callbackMWidget_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "customContextMenuRequested"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(pos))
	}

}

//export callbackMWidget_DragEnterEvent
func callbackMWidget_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *MWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackMWidget_DragLeaveEvent
func callbackMWidget_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *MWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackMWidget_DragMoveEvent
func callbackMWidget_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *MWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackMWidget_DropEvent
func callbackMWidget_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *MWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackMWidget_EnterEvent
func callbackMWidget_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *MWidget) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackMWidget_FocusInEvent
func callbackMWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *MWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackMWidget_FocusOutEvent
func callbackMWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *MWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackMWidget_Hide
func callbackMWidget_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "hide"); signal != nil {
		signal.(func())()
	} else {
		NewMWidgetFromPointer(ptr).HideDefault()
	}
}

func (ptr *MWidget) HideDefault() {
	if ptr.Pointer() != nil {
		C.MWidget_HideDefault(ptr.Pointer())
	}
}

//export callbackMWidget_HideEvent
func callbackMWidget_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *MWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackMWidget_InputMethodEvent
func callbackMWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *MWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackMWidget_KeyPressEvent
func callbackMWidget_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *MWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackMWidget_KeyReleaseEvent
func callbackMWidget_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *MWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackMWidget_LeaveEvent
func callbackMWidget_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *MWidget) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackMWidget_Lower
func callbackMWidget_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "lower"); signal != nil {
		signal.(func())()
	} else {
		NewMWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *MWidget) LowerDefault() {
	if ptr.Pointer() != nil {
		C.MWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackMWidget_MouseDoubleClickEvent
func callbackMWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMWidget_MouseMoveEvent
func callbackMWidget_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMWidget_MousePressEvent
func callbackMWidget_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMWidget_MouseReleaseEvent
func callbackMWidget_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMWidget_MoveEvent
func callbackMWidget_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *MWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackMWidget_PaintEvent
func callbackMWidget_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *MWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackMWidget_Raise
func callbackMWidget_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "raise"); signal != nil {
		signal.(func())()
	} else {
		NewMWidgetFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *MWidget) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.MWidget_RaiseDefault(ptr.Pointer())
	}
}

//export callbackMWidget_Repaint
func callbackMWidget_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewMWidgetFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *MWidget) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.MWidget_RepaintDefault(ptr.Pointer())
	}
}

//export callbackMWidget_ResizeEvent
func callbackMWidget_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *MWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackMWidget_SetDisabled
func callbackMWidget_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewMWidgetFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *MWidget) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.MWidget_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackMWidget_SetEnabled
func callbackMWidget_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewMWidgetFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *MWidget) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.MWidget_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackMWidget_SetFocus2
func callbackMWidget_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewMWidgetFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *MWidget) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.MWidget_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackMWidget_SetHidden
func callbackMWidget_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewMWidgetFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *MWidget) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.MWidget_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackMWidget_SetStyleSheet
func callbackMWidget_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewMWidgetFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *MWidget) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.MWidget_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackMWidget_SetVisible
func callbackMWidget_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewMWidgetFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *MWidget) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.MWidget_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackMWidget_SetWindowModified
func callbackMWidget_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewMWidgetFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *MWidget) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.MWidget_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackMWidget_SetWindowTitle
func callbackMWidget_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewMWidgetFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *MWidget) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.MWidget_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackMWidget_Show
func callbackMWidget_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "show"); signal != nil {
		signal.(func())()
	} else {
		NewMWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *MWidget) ShowDefault() {
	if ptr.Pointer() != nil {
		C.MWidget_ShowDefault(ptr.Pointer())
	}
}

//export callbackMWidget_ShowEvent
func callbackMWidget_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *MWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackMWidget_ShowFullScreen
func callbackMWidget_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewMWidgetFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *MWidget) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.MWidget_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackMWidget_ShowMaximized
func callbackMWidget_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewMWidgetFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *MWidget) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.MWidget_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackMWidget_ShowMinimized
func callbackMWidget_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewMWidgetFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *MWidget) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.MWidget_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackMWidget_ShowNormal
func callbackMWidget_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewMWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *MWidget) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.MWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackMWidget_TabletEvent
func callbackMWidget_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *MWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackMWidget_Update
func callbackMWidget_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "update"); signal != nil {
		signal.(func())()
	} else {
		NewMWidgetFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *MWidget) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.MWidget_UpdateDefault(ptr.Pointer())
	}
}

//export callbackMWidget_UpdateMicroFocus
func callbackMWidget_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewMWidgetFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *MWidget) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.MWidget_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackMWidget_WheelEvent
func callbackMWidget_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *MWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackMWidget_WindowIconChanged
func callbackMWidget_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "windowIconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackMWidget_WindowTitleChanged
func callbackMWidget_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackMWidget_PaintEngine
func callbackMWidget_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewMWidgetFromPointer(ptr).PaintEngineDefault())
}

func (ptr *MWidget) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.MWidget_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackMWidget_MinimumSizeHint
func callbackMWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewMWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *MWidget) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.MWidget_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackMWidget_SizeHint
func callbackMWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewMWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *MWidget) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.MWidget_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackMWidget_InputMethodQuery
func callbackMWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewMWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *MWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.MWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackMWidget_HasHeightForWidth
func callbackMWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewMWidgetFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *MWidget) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.MWidget_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackMWidget_HeightForWidth
func callbackMWidget_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewMWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *MWidget) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackMWidget_Metric
func callbackMWidget_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewMWidgetFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *MWidget) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MWidget_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackMWidget_EventFilter
func callbackMWidget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *MWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.MWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackMWidget_ChildEvent
func callbackMWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *MWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackMWidget_ConnectNotify
func callbackMWidget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewMWidgetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *MWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackMWidget_CustomEvent
func callbackMWidget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *MWidget) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackMWidget_DeleteLater
func callbackMWidget_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewMWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *MWidget) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.MWidget_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackMWidget_Destroyed
func callbackMWidget_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackMWidget_DisconnectNotify
func callbackMWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewMWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *MWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackMWidget_ObjectNameChanged
func callbackMWidget_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackMWidget_TimerEvent
func callbackMWidget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewMWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *MWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type MWindow_ITF interface {
	widgets.QMainWindow_ITF
	MWindow_PTR() *MWindow
=======
type MocLabel_ITF interface {
	widgets.QLabel_ITF
	MocLabel_PTR() *MocLabel
>>>>>>> 93aba97015d6fbc692a2b977f3ee5e96a88a2577
}

func (ptr *MWindow) MWindow_PTR() *MWindow {
	return ptr
}

func (ptr *MWindow) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QMainWindow_PTR().Pointer()
	}
	return nil
}

func (ptr *MWindow) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QMainWindow_PTR().SetPointer(p)
	}
}

func PointerFromMWindow(ptr MWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.MWindow_PTR().Pointer()
	}
	return nil
}

func NewMWindowFromPointer(ptr unsafe.Pointer) *MWindow {
	var n = new(MWindow)
	n.SetPointer(ptr)
	return n
}
func MWindow_QRegisterMetaType() int {
	return int(int32(C.MWindow_MWindow_QRegisterMetaType()))
}

func (ptr *MWindow) QRegisterMetaType() int {
	return int(int32(C.MWindow_MWindow_QRegisterMetaType()))
}

func MWindow_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.MWindow_MWindow_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *MWindow) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.MWindow_MWindow_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func MWindow_QmlRegisterType() int {
	return int(int32(C.MWindow_MWindow_QmlRegisterType()))
}

func (ptr *MWindow) QmlRegisterType() int {
	return int(int32(C.MWindow_MWindow_QmlRegisterType()))
}

func (ptr *MWindow) __resizeDocks_docks_atList(i int) *widgets.QDockWidget {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQDockWidgetFromPointer(C.MWindow___resizeDocks_docks_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MWindow) __resizeDocks_docks_setList(i widgets.QDockWidget_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow___resizeDocks_docks_setList(ptr.Pointer(), widgets.PointerFromQDockWidget(i))
	}
}

func (ptr *MWindow) __resizeDocks_docks_newList() unsafe.Pointer {
	return unsafe.Pointer(C.MWindow___resizeDocks_docks_newList(ptr.Pointer()))
}

func (ptr *MWindow) __resizeDocks_sizes_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MWindow___resizeDocks_sizes_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *MWindow) __resizeDocks_sizes_setList(i int) {
	if ptr.Pointer() != nil {
		C.MWindow___resizeDocks_sizes_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *MWindow) __resizeDocks_sizes_newList() unsafe.Pointer {
	return unsafe.Pointer(C.MWindow___resizeDocks_sizes_newList(ptr.Pointer()))
}

func (ptr *MWindow) __tabifiedDockWidgets_atList(i int) *widgets.QDockWidget {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQDockWidgetFromPointer(C.MWindow___tabifiedDockWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MWindow) __tabifiedDockWidgets_setList(i widgets.QDockWidget_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow___tabifiedDockWidgets_setList(ptr.Pointer(), widgets.PointerFromQDockWidget(i))
	}
}

func (ptr *MWindow) __tabifiedDockWidgets_newList() unsafe.Pointer {
	return unsafe.Pointer(C.MWindow___tabifiedDockWidgets_newList(ptr.Pointer()))
}

func (ptr *MWindow) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.MWindow___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MWindow) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *MWindow) __addActions_actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.MWindow___addActions_actions_newList(ptr.Pointer()))
}

func (ptr *MWindow) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.MWindow___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MWindow) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *MWindow) __insertActions_actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.MWindow___insertActions_actions_newList(ptr.Pointer()))
}

func (ptr *MWindow) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.MWindow___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MWindow) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *MWindow) __actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.MWindow___actions_newList(ptr.Pointer()))
}

func (ptr *MWindow) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.MWindow___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *MWindow) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *MWindow) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.MWindow___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *MWindow) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.MWindow___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MWindow) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *MWindow) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.MWindow___findChildren_newList2(ptr.Pointer()))
}

func (ptr *MWindow) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.MWindow___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MWindow) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *MWindow) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.MWindow___findChildren_newList3(ptr.Pointer()))
}

func (ptr *MWindow) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.MWindow___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MWindow) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *MWindow) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.MWindow___findChildren_newList(ptr.Pointer()))
}

func (ptr *MWindow) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.MWindow___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MWindow) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *MWindow) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.MWindow___children_newList(ptr.Pointer()))
}

func NewMWindow(parent widgets.QWidget_ITF, flags core.Qt__WindowType) *MWindow {
	var tmpValue = NewMWindowFromPointer(C.MWindow_NewMWindow(widgets.PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *MWindow) DestroyMWindow() {
	if ptr.Pointer() != nil {
		C.MWindow_DestroyMWindow(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackMWindow_CreatePopupMenu
func callbackMWindow_CreatePopupMenu(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "createPopupMenu"); signal != nil {
		return widgets.PointerFromQMenu(signal.(func() *widgets.QMenu)())
	}

	return widgets.PointerFromQMenu(NewMWindowFromPointer(ptr).CreatePopupMenuDefault())
}

func (ptr *MWindow) CreatePopupMenuDefault() *widgets.QMenu {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQMenuFromPointer(C.MWindow_CreatePopupMenuDefault(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackMWindow_Event
func callbackMWindow_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMWindowFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *MWindow) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.MWindow_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackMWindow_ContextMenuEvent
func callbackMWindow_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *MWindow) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackMWindow_IconSizeChanged
func callbackMWindow_IconSizeChanged(ptr unsafe.Pointer, iconSize unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "iconSizeChanged"); signal != nil {
		signal.(func(*core.QSize))(core.NewQSizeFromPointer(iconSize))
	}

}

//export callbackMWindow_SetAnimated
func callbackMWindow_SetAnimated(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setAnimated"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	} else {
		NewMWindowFromPointer(ptr).SetAnimatedDefault(int8(enabled) != 0)
	}
}

func (ptr *MWindow) SetAnimatedDefault(enabled bool) {
	if ptr.Pointer() != nil {
		C.MWindow_SetAnimatedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackMWindow_SetDockNestingEnabled
func callbackMWindow_SetDockNestingEnabled(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setDockNestingEnabled"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	} else {
		NewMWindowFromPointer(ptr).SetDockNestingEnabledDefault(int8(enabled) != 0)
	}
}

func (ptr *MWindow) SetDockNestingEnabledDefault(enabled bool) {
	if ptr.Pointer() != nil {
		C.MWindow_SetDockNestingEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackMWindow_SetUnifiedTitleAndToolBarOnMac
func callbackMWindow_SetUnifiedTitleAndToolBarOnMac(ptr unsafe.Pointer, set C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setUnifiedTitleAndToolBarOnMac"); signal != nil {
		signal.(func(bool))(int8(set) != 0)
	} else {
		NewMWindowFromPointer(ptr).SetUnifiedTitleAndToolBarOnMacDefault(int8(set) != 0)
	}
}

func (ptr *MWindow) SetUnifiedTitleAndToolBarOnMacDefault(set bool) {
	if ptr.Pointer() != nil {
		C.MWindow_SetUnifiedTitleAndToolBarOnMacDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(set))))
	}
}

//export callbackMWindow_TabifiedDockWidgetActivated
func callbackMWindow_TabifiedDockWidgetActivated(ptr unsafe.Pointer, dockWidget unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "tabifiedDockWidgetActivated"); signal != nil {
		signal.(func(*widgets.QDockWidget))(widgets.NewQDockWidgetFromPointer(dockWidget))
	}

}

//export callbackMWindow_ToolButtonStyleChanged
func callbackMWindow_ToolButtonStyleChanged(ptr unsafe.Pointer, toolButtonStyle C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "toolButtonStyleChanged"); signal != nil {
		signal.(func(core.Qt__ToolButtonStyle))(core.Qt__ToolButtonStyle(toolButtonStyle))
	}

}

//export callbackMWindow_Close
func callbackMWindow_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewMWindowFromPointer(ptr).CloseDefault())))
}

func (ptr *MWindow) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return C.MWindow_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackMWindow_FocusNextPrevChild
func callbackMWindow_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMWindowFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *MWindow) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return C.MWindow_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackMWindow_NativeEvent
func callbackMWindow_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result C.long) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QByteArray, unsafe.Pointer, int) bool)(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMWindowFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
}

func (ptr *MWindow) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.MWindow_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackMWindow_ActionEvent
func callbackMWindow_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *MWindow) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackMWindow_ChangeEvent
func callbackMWindow_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *MWindow) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackMWindow_CloseEvent
func callbackMWindow_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *MWindow) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackMWindow_CustomContextMenuRequested
func callbackMWindow_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "customContextMenuRequested"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(pos))
	}

}

//export callbackMWindow_DragEnterEvent
func callbackMWindow_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *MWindow) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackMWindow_DragLeaveEvent
func callbackMWindow_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *MWindow) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackMWindow_DragMoveEvent
func callbackMWindow_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *MWindow) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackMWindow_DropEvent
func callbackMWindow_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *MWindow) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackMWindow_EnterEvent
func callbackMWindow_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *MWindow) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackMWindow_FocusInEvent
func callbackMWindow_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *MWindow) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackMWindow_FocusOutEvent
func callbackMWindow_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *MWindow) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackMWindow_Hide
func callbackMWindow_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "hide"); signal != nil {
		signal.(func())()
	} else {
		NewMWindowFromPointer(ptr).HideDefault()
	}
}

func (ptr *MWindow) HideDefault() {
	if ptr.Pointer() != nil {
		C.MWindow_HideDefault(ptr.Pointer())
	}
}

//export callbackMWindow_HideEvent
func callbackMWindow_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *MWindow) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackMWindow_InputMethodEvent
func callbackMWindow_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *MWindow) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackMWindow_KeyPressEvent
func callbackMWindow_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *MWindow) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackMWindow_KeyReleaseEvent
func callbackMWindow_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *MWindow) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackMWindow_LeaveEvent
func callbackMWindow_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *MWindow) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackMWindow_Lower
func callbackMWindow_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "lower"); signal != nil {
		signal.(func())()
	} else {
		NewMWindowFromPointer(ptr).LowerDefault()
	}
}

func (ptr *MWindow) LowerDefault() {
	if ptr.Pointer() != nil {
		C.MWindow_LowerDefault(ptr.Pointer())
	}
}

//export callbackMWindow_MouseDoubleClickEvent
func callbackMWindow_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MWindow) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMWindow_MouseMoveEvent
func callbackMWindow_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MWindow) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMWindow_MousePressEvent
func callbackMWindow_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MWindow) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMWindow_MouseReleaseEvent
func callbackMWindow_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MWindow) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMWindow_MoveEvent
func callbackMWindow_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *MWindow) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackMWindow_PaintEvent
func callbackMWindow_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *MWindow) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackMWindow_Raise
func callbackMWindow_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "raise"); signal != nil {
		signal.(func())()
	} else {
		NewMWindowFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *MWindow) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.MWindow_RaiseDefault(ptr.Pointer())
	}
}

//export callbackMWindow_Repaint
func callbackMWindow_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewMWindowFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *MWindow) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.MWindow_RepaintDefault(ptr.Pointer())
	}
}

//export callbackMWindow_ResizeEvent
func callbackMWindow_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *MWindow) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackMWindow_SetDisabled
func callbackMWindow_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewMWindowFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *MWindow) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.MWindow_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackMWindow_SetEnabled
func callbackMWindow_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewMWindowFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *MWindow) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.MWindow_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackMWindow_SetFocus2
func callbackMWindow_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewMWindowFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *MWindow) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.MWindow_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackMWindow_SetHidden
func callbackMWindow_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewMWindowFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *MWindow) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.MWindow_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackMWindow_SetStyleSheet
func callbackMWindow_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewMWindowFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *MWindow) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.MWindow_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackMWindow_SetVisible
func callbackMWindow_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewMWindowFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *MWindow) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.MWindow_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackMWindow_SetWindowModified
func callbackMWindow_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewMWindowFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *MWindow) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.MWindow_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackMWindow_SetWindowTitle
func callbackMWindow_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewMWindowFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *MWindow) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.MWindow_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackMWindow_Show
func callbackMWindow_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "show"); signal != nil {
		signal.(func())()
	} else {
		NewMWindowFromPointer(ptr).ShowDefault()
	}
}

func (ptr *MWindow) ShowDefault() {
	if ptr.Pointer() != nil {
		C.MWindow_ShowDefault(ptr.Pointer())
	}
}

//export callbackMWindow_ShowEvent
func callbackMWindow_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *MWindow) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackMWindow_ShowFullScreen
func callbackMWindow_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewMWindowFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *MWindow) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.MWindow_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackMWindow_ShowMaximized
func callbackMWindow_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewMWindowFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *MWindow) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.MWindow_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackMWindow_ShowMinimized
func callbackMWindow_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewMWindowFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *MWindow) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.MWindow_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackMWindow_ShowNormal
func callbackMWindow_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewMWindowFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *MWindow) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.MWindow_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackMWindow_TabletEvent
func callbackMWindow_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *MWindow) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackMWindow_Update
func callbackMWindow_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "update"); signal != nil {
		signal.(func())()
	} else {
		NewMWindowFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *MWindow) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.MWindow_UpdateDefault(ptr.Pointer())
	}
}

//export callbackMWindow_UpdateMicroFocus
func callbackMWindow_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewMWindowFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *MWindow) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.MWindow_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackMWindow_WheelEvent
func callbackMWindow_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *MWindow) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackMWindow_WindowIconChanged
func callbackMWindow_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "windowIconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackMWindow_WindowTitleChanged
func callbackMWindow_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackMWindow_PaintEngine
func callbackMWindow_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewMWindowFromPointer(ptr).PaintEngineDefault())
}

func (ptr *MWindow) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.MWindow_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackMWindow_MinimumSizeHint
func callbackMWindow_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewMWindowFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *MWindow) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.MWindow_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackMWindow_SizeHint
func callbackMWindow_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewMWindowFromPointer(ptr).SizeHintDefault())
}

func (ptr *MWindow) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.MWindow_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackMWindow_InputMethodQuery
func callbackMWindow_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewMWindowFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *MWindow) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.MWindow_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackMWindow_HasHeightForWidth
func callbackMWindow_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewMWindowFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *MWindow) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.MWindow_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackMWindow_HeightForWidth
func callbackMWindow_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewMWindowFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *MWindow) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MWindow_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackMWindow_Metric
func callbackMWindow_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewMWindowFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *MWindow) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MWindow_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackMWindow_EventFilter
func callbackMWindow_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMWindowFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *MWindow) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.MWindow_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackMWindow_ChildEvent
func callbackMWindow_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *MWindow) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackMWindow_ConnectNotify
func callbackMWindow_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewMWindowFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *MWindow) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackMWindow_CustomEvent
func callbackMWindow_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *MWindow) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackMWindow_DeleteLater
func callbackMWindow_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewMWindowFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *MWindow) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.MWindow_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackMWindow_Destroyed
func callbackMWindow_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackMWindow_DisconnectNotify
func callbackMWindow_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewMWindowFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *MWindow) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackMWindow_ObjectNameChanged
func callbackMWindow_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackMWindow_TimerEvent
func callbackMWindow_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewMWindowFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *MWindow) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MWindow_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}
