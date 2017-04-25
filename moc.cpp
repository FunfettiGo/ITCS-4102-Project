

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAction>
#include <QActionEvent>
#include <QApplication>
#include <QByteArray>
#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDockWidget>
#include <QDrag>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QFocusEvent>
#include <QHideEvent>
#include <QIcon>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QLabel>
#include <QList>
#include <QMainWindow>
#include <QMenu>
#include <QMetaMethod>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QMovie>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPalette>
#include <QPicture>
#include <QPixmap>
#include <QPoint>
#include <QResizeEvent>
#include <QScreen>
#include <QSessionManager>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QTabletEvent>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QWheelEvent>
#include <QWidget>
#include <QWindow>


class App: public QApplication
{
Q_OBJECT
public:
	App(int &argc, char **argv) : QApplication(argc, argv) {};
	 ~App() { callbackApp_DestroyApp(this); };
	bool event(QEvent * e) { return callbackApp_Event(this, e) != 0; };
	void aboutQt() { callbackApp_AboutQt(this); };
	void closeAllWindows() { callbackApp_CloseAllWindows(this); };
	void Signal_FocusChanged(QWidget * old, QWidget * now) { callbackApp_FocusChanged(this, old, now); };
	void setAutoSipEnabled(const bool enabled) { callbackApp_SetAutoSipEnabled(this, enabled); };
	void setStyleSheet(const QString & sheet) { QByteArray t542ebc = sheet.toUtf8(); Moc_PackedString sheetPacked = { const_cast<char*>(t542ebc.prepend("WHITESPACE").constData()+10), t542ebc.size()-10 };callbackApp_SetStyleSheet(this, sheetPacked); };
	bool autoSipEnabled() const { return callbackApp_AutoSipEnabled(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	void Signal_ApplicationDisplayNameChanged() { callbackApp_ApplicationDisplayNameChanged(this); };
	void Signal_ApplicationStateChanged(Qt::ApplicationState state) { callbackApp_ApplicationStateChanged(this, state); };
	void Signal_CommitDataRequest(QSessionManager & manager) { callbackApp_CommitDataRequest(this, static_cast<QSessionManager*>(&manager)); };
	void Signal_FocusObjectChanged(QObject * focusObject) { callbackApp_FocusObjectChanged(this, focusObject); };
	void Signal_FocusWindowChanged(QWindow * focusWindow) { callbackApp_FocusWindowChanged(this, focusWindow); };
	void Signal_FontDatabaseChanged() { callbackApp_FontDatabaseChanged(this); };
	void Signal_LastWindowClosed() { callbackApp_LastWindowClosed(this); };
	void Signal_LayoutDirectionChanged(Qt::LayoutDirection direction) { callbackApp_LayoutDirectionChanged(this, direction); };
	void Signal_PaletteChanged(const QPalette & palette) { callbackApp_PaletteChanged(this, const_cast<QPalette*>(&palette)); };
	void Signal_PrimaryScreenChanged(QScreen * screen) { callbackApp_PrimaryScreenChanged(this, screen); };
	void Signal_SaveStateRequest(QSessionManager & manager) { callbackApp_SaveStateRequest(this, static_cast<QSessionManager*>(&manager)); };
	void Signal_ScreenAdded(QScreen * screen) { callbackApp_ScreenAdded(this, screen); };
	void Signal_ScreenRemoved(QScreen * screen) { callbackApp_ScreenRemoved(this, screen); };
	void Signal_AboutToQuit() { callbackApp_AboutToQuit(this); };
	void quit() { callbackApp_Quit(this); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackApp_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackApp_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackApp_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackApp_CustomEvent(this, event); };
	void deleteLater() { callbackApp_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackApp_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackApp_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackApp_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackApp_TimerEvent(this, event); };
	
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(App*)

class MWidget: public QWidget
{
Q_OBJECT
public:
	MWidget(QWidget *parent = Q_NULLPTR, Qt::WindowFlags f = Qt::WindowFlags()) : QWidget(parent, f) {};
	bool close() { return callbackMWidget_Close(this) != 0; };
	bool event(QEvent * event) { return callbackMWidget_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackMWidget_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackMWidget_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void actionEvent(QActionEvent * event) { callbackMWidget_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackMWidget_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackMWidget_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackMWidget_ContextMenuEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackMWidget_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackMWidget_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackMWidget_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackMWidget_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackMWidget_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackMWidget_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackMWidget_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackMWidget_FocusOutEvent(this, event); };
	void hide() { callbackMWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackMWidget_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackMWidget_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackMWidget_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackMWidget_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackMWidget_LeaveEvent(this, event); };
	void lower() { callbackMWidget_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackMWidget_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackMWidget_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackMWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackMWidget_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackMWidget_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackMWidget_PaintEvent(this, event); };
	void raise() { callbackMWidget_Raise(this); };
	void repaint() { callbackMWidget_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackMWidget_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackMWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackMWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackMWidget_SetFocus2(this); };
	void setHidden(bool hidden) { callbackMWidget_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); Moc_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackMWidget_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackMWidget_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackMWidget_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackMWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackMWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackMWidget_ShowEvent(this, event); };
	void showFullScreen() { callbackMWidget_ShowFullScreen(this); };
	void showMaximized() { callbackMWidget_ShowMaximized(this); };
	void showMinimized() { callbackMWidget_ShowMinimized(this); };
	void showNormal() { callbackMWidget_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackMWidget_TabletEvent(this, event); };
	void update() { callbackMWidget_Update(this); };
	void updateMicroFocus() { callbackMWidget_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackMWidget_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackMWidget_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackMWidget_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackMWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackMWidget_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackMWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackMWidget_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackMWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackMWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackMWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackMWidget_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackMWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackMWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackMWidget_CustomEvent(this, event); };
	void deleteLater() { callbackMWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackMWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackMWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackMWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackMWidget_TimerEvent(this, event); };
	
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(MWidget*)

class MWindow: public QMainWindow
{
Q_OBJECT
public:
	MWindow(QWidget *parent = Q_NULLPTR, Qt::WindowFlags flags = Qt::WindowFlags()) : QMainWindow(parent, flags) {};
	QMenu * createPopupMenu() { return static_cast<QMenu*>(callbackMWindow_CreatePopupMenu(this)); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackMWindow_ContextMenuEvent(this, event); };
	void Signal_IconSizeChanged(const QSize & iconSize) { callbackMWindow_IconSizeChanged(this, const_cast<QSize*>(&iconSize)); };
	void setAnimated(bool enabled) { callbackMWindow_SetAnimated(this, enabled); };
	void setDockNestingEnabled(bool enabled) { callbackMWindow_SetDockNestingEnabled(this, enabled); };
	void setUnifiedTitleAndToolBarOnMac(bool set) { callbackMWindow_SetUnifiedTitleAndToolBarOnMac(this, set); };
	void Signal_TabifiedDockWidgetActivated(QDockWidget * dockWidget) { callbackMWindow_TabifiedDockWidgetActivated(this, dockWidget); };
	void Signal_ToolButtonStyleChanged(Qt::ToolButtonStyle toolButtonStyle) { callbackMWindow_ToolButtonStyleChanged(this, toolButtonStyle); };
	bool close() { return callbackMWindow_Close(this) != 0; };
	bool event(QEvent * event) { return callbackMWindow_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackMWindow_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackMWindow_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void actionEvent(QActionEvent * event) { callbackMWindow_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackMWindow_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackMWindow_CloseEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackMWindow_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackMWindow_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackMWindow_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackMWindow_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackMWindow_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackMWindow_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackMWindow_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackMWindow_FocusOutEvent(this, event); };
	void hide() { callbackMWindow_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackMWindow_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackMWindow_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackMWindow_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackMWindow_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackMWindow_LeaveEvent(this, event); };
	void lower() { callbackMWindow_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackMWindow_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackMWindow_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackMWindow_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackMWindow_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackMWindow_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackMWindow_PaintEvent(this, event); };
	void raise() { callbackMWindow_Raise(this); };
	void repaint() { callbackMWindow_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackMWindow_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackMWindow_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackMWindow_SetEnabled(this, vbo); };
	void setFocus() { callbackMWindow_SetFocus2(this); };
	void setHidden(bool hidden) { callbackMWindow_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); Moc_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackMWindow_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackMWindow_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackMWindow_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackMWindow_SetWindowTitle(this, vqsPacked); };
	void show() { callbackMWindow_Show(this); };
	void showEvent(QShowEvent * event) { callbackMWindow_ShowEvent(this, event); };
	void showFullScreen() { callbackMWindow_ShowFullScreen(this); };
	void showMaximized() { callbackMWindow_ShowMaximized(this); };
	void showMinimized() { callbackMWindow_ShowMinimized(this); };
	void showNormal() { callbackMWindow_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackMWindow_TabletEvent(this, event); };
	void update() { callbackMWindow_Update(this); };
	void updateMicroFocus() { callbackMWindow_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackMWindow_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackMWindow_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackMWindow_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackMWindow_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackMWindow_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackMWindow_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackMWindow_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackMWindow_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackMWindow_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackMWindow_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackMWindow_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackMWindow_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackMWindow_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackMWindow_CustomEvent(this, event); };
	void deleteLater() { callbackMWindow_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackMWindow_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackMWindow_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackMWindow_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackMWindow_TimerEvent(this, event); };
	
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(MWindow*)

class MocLabel: public QLabel
{
Q_OBJECT
public:
	MocLabel(QWidget *parent = Q_NULLPTR, Qt::WindowFlags f = Qt::WindowFlags()) : QLabel(parent, f) {};
	MocLabel(const QString &text, QWidget *parent = Q_NULLPTR, Qt::WindowFlags f = Qt::WindowFlags()) : QLabel(text, parent, f) {};
	void Signal_UpdateLabel(QString v0) { QByteArray tea1dd7 = v0.toUtf8(); Moc_PackedString v0Packed = { const_cast<char*>(tea1dd7.prepend("WHITESPACE").constData()+10), tea1dd7.size()-10 };callbackMocLabel_UpdateLabel(this, v0Packed); };
	bool focusNextPrevChild(bool next) { return callbackMocLabel_FocusNextPrevChild(this, next) != 0; };
	void changeEvent(QEvent * ev) { callbackMocLabel_ChangeEvent(this, ev); };
	void clear() { callbackMocLabel_Clear(this); };
	void contextMenuEvent(QContextMenuEvent * ev) { callbackMocLabel_ContextMenuEvent(this, ev); };
	void focusInEvent(QFocusEvent * ev) { callbackMocLabel_FocusInEvent(this, ev); };
	void focusOutEvent(QFocusEvent * ev) { callbackMocLabel_FocusOutEvent(this, ev); };
	void keyPressEvent(QKeyEvent * ev) { callbackMocLabel_KeyPressEvent(this, ev); };
	void Signal_LinkActivated(const QString & link) { QByteArray t4f0aa5 = link.toUtf8(); Moc_PackedString linkPacked = { const_cast<char*>(t4f0aa5.prepend("WHITESPACE").constData()+10), t4f0aa5.size()-10 };callbackMocLabel_LinkActivated(this, linkPacked); };
	void Signal_LinkHovered(const QString & link) { QByteArray t4f0aa5 = link.toUtf8(); Moc_PackedString linkPacked = { const_cast<char*>(t4f0aa5.prepend("WHITESPACE").constData()+10), t4f0aa5.size()-10 };callbackMocLabel_LinkHovered(this, linkPacked); };
	void mouseMoveEvent(QMouseEvent * ev) { callbackMocLabel_MouseMoveEvent(this, ev); };
	void mousePressEvent(QMouseEvent * ev) { callbackMocLabel_MousePressEvent(this, ev); };
	void mouseReleaseEvent(QMouseEvent * ev) { callbackMocLabel_MouseReleaseEvent(this, ev); };
	void paintEvent(QPaintEvent * vqp) { callbackMocLabel_PaintEvent(this, vqp); };
	void setMovie(QMovie * movie) { callbackMocLabel_SetMovie(this, movie); };
	void setNum(double num) { callbackMocLabel_SetNum2(this, num); };
	void setNum(int num) { callbackMocLabel_SetNum(this, num); };
	void setPicture(const QPicture & picture) { callbackMocLabel_SetPicture(this, const_cast<QPicture*>(&picture)); };
	void setPixmap(const QPixmap & vqp) { callbackMocLabel_SetPixmap(this, const_cast<QPixmap*>(&vqp)); };
	void setText(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackMocLabel_SetText(this, vqsPacked); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackMocLabel_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackMocLabel_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	int heightForWidth(int w) const { return callbackMocLabel_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	bool close() { return callbackMocLabel_Close(this) != 0; };
	bool event(QEvent * event) { return callbackMocLabel_Event(this, event) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackMocLabel_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void actionEvent(QActionEvent * event) { callbackMocLabel_ActionEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackMocLabel_CloseEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackMocLabel_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackMocLabel_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackMocLabel_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackMocLabel_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackMocLabel_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackMocLabel_EnterEvent(this, event); };
	void hide() { callbackMocLabel_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackMocLabel_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackMocLabel_InputMethodEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackMocLabel_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackMocLabel_LeaveEvent(this, event); };
	void lower() { callbackMocLabel_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackMocLabel_MouseDoubleClickEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackMocLabel_MoveEvent(this, event); };
	void raise() { callbackMocLabel_Raise(this); };
	void repaint() { callbackMocLabel_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackMocLabel_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackMocLabel_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackMocLabel_SetEnabled(this, vbo); };
	void setFocus() { callbackMocLabel_SetFocus2(this); };
	void setHidden(bool hidden) { callbackMocLabel_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); Moc_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackMocLabel_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackMocLabel_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackMocLabel_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackMocLabel_SetWindowTitle(this, vqsPacked); };
	void show() { callbackMocLabel_Show(this); };
	void showEvent(QShowEvent * event) { callbackMocLabel_ShowEvent(this, event); };
	void showFullScreen() { callbackMocLabel_ShowFullScreen(this); };
	void showMaximized() { callbackMocLabel_ShowMaximized(this); };
	void showMinimized() { callbackMocLabel_ShowMinimized(this); };
	void showNormal() { callbackMocLabel_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackMocLabel_TabletEvent(this, event); };
	void update() { callbackMocLabel_Update(this); };
	void updateMicroFocus() { callbackMocLabel_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackMocLabel_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackMocLabel_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackMocLabel_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackMocLabel_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackMocLabel_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackMocLabel_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackMocLabel_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackMocLabel_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackMocLabel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackMocLabel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackMocLabel_CustomEvent(this, event); };
	void deleteLater() { callbackMocLabel_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackMocLabel_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackMocLabel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackMocLabel_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackMocLabel_TimerEvent(this, event); };
	
signals:
	void updateLabel(QString v0);
public slots:
private:
};

Q_DECLARE_METATYPE(MocLabel*)

int App_App_QRegisterMetaType()
{
	return qRegisterMetaType<App*>();
}

int App_App_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<App>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int App_App_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<App>();
#else
	return 0;
#endif
}

void* App___allWidgets_atList(void* ptr, int i)
{
	return const_cast<QWidget*>(static_cast<QList<QWidget *>*>(ptr)->at(i));
}

void App___allWidgets_setList(void* ptr, void* i)
{
	static_cast<QList<QWidget *>*>(ptr)->append(static_cast<QWidget*>(i));
}

void* App___allWidgets_newList(void* ptr)
{
	return new QList<QWidget *>;
}

void* App___topLevelWidgets_atList(void* ptr, int i)
{
	return const_cast<QWidget*>(static_cast<QList<QWidget *>*>(ptr)->at(i));
}

void App___topLevelWidgets_setList(void* ptr, void* i)
{
	static_cast<QList<QWidget *>*>(ptr)->append(static_cast<QWidget*>(i));
}

void* App___topLevelWidgets_newList(void* ptr)
{
	return new QList<QWidget *>;
}

void* App___screens_atList(void* ptr, int i)
{
	return const_cast<QScreen*>(static_cast<QList<QScreen *>*>(ptr)->at(i));
}

void App___screens_setList(void* ptr, void* i)
{
	static_cast<QList<QScreen *>*>(ptr)->append(static_cast<QScreen*>(i));
}

void* App___screens_newList(void* ptr)
{
	return new QList<QScreen *>;
}

void* App___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void App___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* App___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* App___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void App___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* App___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* App___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void App___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* App___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* App___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void App___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* App___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* App___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void App___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* App___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

void* App_NewApp(int argc, char* argv)
{
	static int argcs = argc;
	static char** argvs = static_cast<char**>(malloc(argcs * sizeof(char*)));

	QList<QByteArray> aList = QByteArray(argv).split('|');
	for (int i = 0; i < argcs; i++)
		argvs[i] = (new QByteArray(aList.at(i)))->data();

	return new App(argcs, argvs);
}

void App_DestroyApp(void* ptr)
{
	static_cast<App*>(ptr)->~App();
}

void App_DestroyAppDefault(void* ptr)
{

}

char App_EventDefault(void* ptr, void* e)
{
	return static_cast<App*>(ptr)->QApplication::event(static_cast<QEvent*>(e));
}

void App_AboutQtDefault(void* ptr)
{
	static_cast<App*>(ptr)->QApplication::aboutQt();
}

void App_CloseAllWindowsDefault(void* ptr)
{
	static_cast<App*>(ptr)->QApplication::closeAllWindows();
}

void App_SetAutoSipEnabledDefault(void* ptr, char enabled)
{
	static_cast<App*>(ptr)->QApplication::setAutoSipEnabled(enabled != 0);
}

void App_SetStyleSheetDefault(void* ptr, char* sheet)
{
	static_cast<App*>(ptr)->QApplication::setStyleSheet(QString(sheet));
}

char App_AutoSipEnabledDefault(void* ptr)
{
	return static_cast<App*>(ptr)->QApplication::autoSipEnabled();
}

void App_QuitDefault(void* ptr)
{
	static_cast<App*>(ptr)->QApplication::quit();
}

char App_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<App*>(ptr)->QApplication::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void App_ChildEventDefault(void* ptr, void* event)
{
	static_cast<App*>(ptr)->QApplication::childEvent(static_cast<QChildEvent*>(event));
}

void App_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<App*>(ptr)->QApplication::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void App_CustomEventDefault(void* ptr, void* event)
{
	static_cast<App*>(ptr)->QApplication::customEvent(static_cast<QEvent*>(event));
}

void App_DeleteLaterDefault(void* ptr)
{
	static_cast<App*>(ptr)->QApplication::deleteLater();
}

void App_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<App*>(ptr)->QApplication::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void App_TimerEventDefault(void* ptr, void* event)
{
	static_cast<App*>(ptr)->QApplication::timerEvent(static_cast<QTimerEvent*>(event));
}



int MWidget_MWidget_QRegisterMetaType()
{
	return qRegisterMetaType<MWidget*>();
}

int MWidget_MWidget_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<MWidget>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int MWidget_MWidget_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<MWidget>();
#else
	return 0;
#endif
}

void* MWidget___addActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void MWidget___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MWidget___addActions_actions_newList(void* ptr)
{
	return new QList<QAction *>;
}

void* MWidget___insertActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void MWidget___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MWidget___insertActions_actions_newList(void* ptr)
{
	return new QList<QAction *>;
}

void* MWidget___actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void MWidget___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MWidget___actions_newList(void* ptr)
{
	return new QList<QAction *>;
}

void* MWidget___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void MWidget___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* MWidget___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* MWidget___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void MWidget___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MWidget___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* MWidget___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void MWidget___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MWidget___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* MWidget___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void MWidget___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MWidget___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* MWidget___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void MWidget___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MWidget___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

void* MWidget_NewMWidget(void* parent, long long fo)
{
		return new MWidget(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(fo));
}

void MWidget_DestroyMWidget(void* ptr)
{
	static_cast<MWidget*>(ptr)->~MWidget();
}

char MWidget_CloseDefault(void* ptr)
{
	return static_cast<MWidget*>(ptr)->QWidget::close();
}

char MWidget_EventDefault(void* ptr, void* event)
{
	return static_cast<MWidget*>(ptr)->QWidget::event(static_cast<QEvent*>(event));
}

char MWidget_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<MWidget*>(ptr)->QWidget::focusNextPrevChild(next != 0);
}

char MWidget_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
	return static_cast<MWidget*>(ptr)->QWidget::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void MWidget_ActionEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void MWidget_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::changeEvent(static_cast<QEvent*>(event));
}

void MWidget_CloseEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void MWidget_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void MWidget_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void MWidget_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void MWidget_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void MWidget_DropEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void MWidget_EnterEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::enterEvent(static_cast<QEvent*>(event));
}

void MWidget_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void MWidget_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void MWidget_HideDefault(void* ptr)
{
	static_cast<MWidget*>(ptr)->QWidget::hide();
}

void MWidget_HideEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void MWidget_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void MWidget_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void MWidget_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void MWidget_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::leaveEvent(static_cast<QEvent*>(event));
}

void MWidget_LowerDefault(void* ptr)
{
	static_cast<MWidget*>(ptr)->QWidget::lower();
}

void MWidget_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void MWidget_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void MWidget_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void MWidget_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void MWidget_MoveEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void MWidget_PaintEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void MWidget_RaiseDefault(void* ptr)
{
	static_cast<MWidget*>(ptr)->QWidget::raise();
}

void MWidget_RepaintDefault(void* ptr)
{
	static_cast<MWidget*>(ptr)->QWidget::repaint();
}

void MWidget_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void MWidget_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<MWidget*>(ptr)->QWidget::setDisabled(disable != 0);
}

void MWidget_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<MWidget*>(ptr)->QWidget::setEnabled(vbo != 0);
}

void MWidget_SetFocus2Default(void* ptr)
{
	static_cast<MWidget*>(ptr)->QWidget::setFocus();
}

void MWidget_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<MWidget*>(ptr)->QWidget::setHidden(hidden != 0);
}

void MWidget_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<MWidget*>(ptr)->QWidget::setStyleSheet(QString(styleSheet));
}

void MWidget_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<MWidget*>(ptr)->QWidget::setVisible(visible != 0);
}

void MWidget_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<MWidget*>(ptr)->QWidget::setWindowModified(vbo != 0);
}

void MWidget_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<MWidget*>(ptr)->QWidget::setWindowTitle(QString(vqs));
}

void MWidget_ShowDefault(void* ptr)
{
	static_cast<MWidget*>(ptr)->QWidget::show();
}

void MWidget_ShowEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::showEvent(static_cast<QShowEvent*>(event));
}

void MWidget_ShowFullScreenDefault(void* ptr)
{
	static_cast<MWidget*>(ptr)->QWidget::showFullScreen();
}

void MWidget_ShowMaximizedDefault(void* ptr)
{
	static_cast<MWidget*>(ptr)->QWidget::showMaximized();
}

void MWidget_ShowMinimizedDefault(void* ptr)
{
	static_cast<MWidget*>(ptr)->QWidget::showMinimized();
}

void MWidget_ShowNormalDefault(void* ptr)
{
	static_cast<MWidget*>(ptr)->QWidget::showNormal();
}

void MWidget_TabletEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void MWidget_UpdateDefault(void* ptr)
{
	static_cast<MWidget*>(ptr)->QWidget::update();
}

void MWidget_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<MWidget*>(ptr)->QWidget::updateMicroFocus();
}

void MWidget_WheelEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* MWidget_PaintEngineDefault(void* ptr)
{
	return static_cast<MWidget*>(ptr)->QWidget::paintEngine();
}

void* MWidget_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<MWidget*>(ptr)->QWidget::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* MWidget_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<MWidget*>(ptr)->QWidget::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* MWidget_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<MWidget*>(ptr)->QWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char MWidget_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<MWidget*>(ptr)->QWidget::hasHeightForWidth();
}

int MWidget_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<MWidget*>(ptr)->QWidget::heightForWidth(w);
}

int MWidget_MetricDefault(void* ptr, long long m)
{
	return static_cast<MWidget*>(ptr)->QWidget::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

char MWidget_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<MWidget*>(ptr)->QWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void MWidget_ChildEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::childEvent(static_cast<QChildEvent*>(event));
}

void MWidget_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<MWidget*>(ptr)->QWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void MWidget_CustomEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::customEvent(static_cast<QEvent*>(event));
}

void MWidget_DeleteLaterDefault(void* ptr)
{
	static_cast<MWidget*>(ptr)->QWidget::deleteLater();
}

void MWidget_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<MWidget*>(ptr)->QWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void MWidget_TimerEventDefault(void* ptr, void* event)
{
	static_cast<MWidget*>(ptr)->QWidget::timerEvent(static_cast<QTimerEvent*>(event));
}



int MWindow_MWindow_QRegisterMetaType()
{
	return qRegisterMetaType<MWindow*>();
}

int MWindow_MWindow_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<MWindow>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int MWindow_MWindow_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<MWindow>();
#else
	return 0;
#endif
}

void* MWindow___resizeDocks_docks_atList(void* ptr, int i)
{
	return const_cast<QDockWidget*>(static_cast<QList<QDockWidget *>*>(ptr)->at(i));
}

void MWindow___resizeDocks_docks_setList(void* ptr, void* i)
{
	static_cast<QList<QDockWidget *>*>(ptr)->append(static_cast<QDockWidget*>(i));
}

void* MWindow___resizeDocks_docks_newList(void* ptr)
{
	return new QList<QDockWidget *>;
}

int MWindow___resizeDocks_sizes_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void MWindow___resizeDocks_sizes_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* MWindow___resizeDocks_sizes_newList(void* ptr)
{
	return new QList<int>;
}

void* MWindow___tabifiedDockWidgets_atList(void* ptr, int i)
{
	return const_cast<QDockWidget*>(static_cast<QList<QDockWidget *>*>(ptr)->at(i));
}

void MWindow___tabifiedDockWidgets_setList(void* ptr, void* i)
{
	static_cast<QList<QDockWidget *>*>(ptr)->append(static_cast<QDockWidget*>(i));
}

void* MWindow___tabifiedDockWidgets_newList(void* ptr)
{
	return new QList<QDockWidget *>;
}

void* MWindow___addActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void MWindow___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MWindow___addActions_actions_newList(void* ptr)
{
	return new QList<QAction *>;
}

void* MWindow___insertActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void MWindow___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MWindow___insertActions_actions_newList(void* ptr)
{
	return new QList<QAction *>;
}

void* MWindow___actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void MWindow___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MWindow___actions_newList(void* ptr)
{
	return new QList<QAction *>;
}

void* MWindow___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void MWindow___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* MWindow___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* MWindow___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void MWindow___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MWindow___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* MWindow___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void MWindow___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MWindow___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* MWindow___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void MWindow___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MWindow___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* MWindow___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void MWindow___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MWindow___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

void* MWindow_NewMWindow(void* parent, long long flags)
{
		return new MWindow(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void MWindow_DestroyMWindow(void* ptr)
{
	static_cast<MWindow*>(ptr)->~MWindow();
}

void* MWindow_CreatePopupMenuDefault(void* ptr)
{
	return static_cast<MWindow*>(ptr)->QMainWindow::createPopupMenu();
}

char MWindow_EventDefault(void* ptr, void* event)
{
	return static_cast<MWindow*>(ptr)->QMainWindow::event(static_cast<QEvent*>(event));
}

void MWindow_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void MWindow_SetAnimatedDefault(void* ptr, char enabled)
{
	static_cast<MWindow*>(ptr)->QMainWindow::setAnimated(enabled != 0);
}

void MWindow_SetDockNestingEnabledDefault(void* ptr, char enabled)
{
	static_cast<MWindow*>(ptr)->QMainWindow::setDockNestingEnabled(enabled != 0);
}

void MWindow_SetUnifiedTitleAndToolBarOnMacDefault(void* ptr, char set)
{
	static_cast<MWindow*>(ptr)->QMainWindow::setUnifiedTitleAndToolBarOnMac(set != 0);
}

char MWindow_CloseDefault(void* ptr)
{
	return static_cast<MWindow*>(ptr)->QMainWindow::close();
}

char MWindow_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<MWindow*>(ptr)->QMainWindow::focusNextPrevChild(next != 0);
}

char MWindow_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
	return static_cast<MWindow*>(ptr)->QMainWindow::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void MWindow_ActionEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::actionEvent(static_cast<QActionEvent*>(event));
}

void MWindow_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::changeEvent(static_cast<QEvent*>(event));
}

void MWindow_CloseEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::closeEvent(static_cast<QCloseEvent*>(event));
}

void MWindow_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void MWindow_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void MWindow_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void MWindow_DropEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::dropEvent(static_cast<QDropEvent*>(event));
}

void MWindow_EnterEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::enterEvent(static_cast<QEvent*>(event));
}

void MWindow_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::focusInEvent(static_cast<QFocusEvent*>(event));
}

void MWindow_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void MWindow_HideDefault(void* ptr)
{
	static_cast<MWindow*>(ptr)->QMainWindow::hide();
}

void MWindow_HideEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::hideEvent(static_cast<QHideEvent*>(event));
}

void MWindow_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void MWindow_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void MWindow_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void MWindow_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::leaveEvent(static_cast<QEvent*>(event));
}

void MWindow_LowerDefault(void* ptr)
{
	static_cast<MWindow*>(ptr)->QMainWindow::lower();
}

void MWindow_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void MWindow_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void MWindow_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void MWindow_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void MWindow_MoveEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::moveEvent(static_cast<QMoveEvent*>(event));
}

void MWindow_PaintEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::paintEvent(static_cast<QPaintEvent*>(event));
}

void MWindow_RaiseDefault(void* ptr)
{
	static_cast<MWindow*>(ptr)->QMainWindow::raise();
}

void MWindow_RepaintDefault(void* ptr)
{
	static_cast<MWindow*>(ptr)->QMainWindow::repaint();
}

void MWindow_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::resizeEvent(static_cast<QResizeEvent*>(event));
}

void MWindow_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<MWindow*>(ptr)->QMainWindow::setDisabled(disable != 0);
}

void MWindow_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<MWindow*>(ptr)->QMainWindow::setEnabled(vbo != 0);
}

void MWindow_SetFocus2Default(void* ptr)
{
	static_cast<MWindow*>(ptr)->QMainWindow::setFocus();
}

void MWindow_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<MWindow*>(ptr)->QMainWindow::setHidden(hidden != 0);
}

void MWindow_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<MWindow*>(ptr)->QMainWindow::setStyleSheet(QString(styleSheet));
}

void MWindow_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<MWindow*>(ptr)->QMainWindow::setVisible(visible != 0);
}

void MWindow_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<MWindow*>(ptr)->QMainWindow::setWindowModified(vbo != 0);
}

void MWindow_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<MWindow*>(ptr)->QMainWindow::setWindowTitle(QString(vqs));
}

void MWindow_ShowDefault(void* ptr)
{
	static_cast<MWindow*>(ptr)->QMainWindow::show();
}

void MWindow_ShowEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::showEvent(static_cast<QShowEvent*>(event));
}

void MWindow_ShowFullScreenDefault(void* ptr)
{
	static_cast<MWindow*>(ptr)->QMainWindow::showFullScreen();
}

void MWindow_ShowMaximizedDefault(void* ptr)
{
	static_cast<MWindow*>(ptr)->QMainWindow::showMaximized();
}

void MWindow_ShowMinimizedDefault(void* ptr)
{
	static_cast<MWindow*>(ptr)->QMainWindow::showMinimized();
}

void MWindow_ShowNormalDefault(void* ptr)
{
	static_cast<MWindow*>(ptr)->QMainWindow::showNormal();
}

void MWindow_TabletEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::tabletEvent(static_cast<QTabletEvent*>(event));
}

void MWindow_UpdateDefault(void* ptr)
{
	static_cast<MWindow*>(ptr)->QMainWindow::update();
}

void MWindow_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<MWindow*>(ptr)->QMainWindow::updateMicroFocus();
}

void MWindow_WheelEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* MWindow_PaintEngineDefault(void* ptr)
{
	return static_cast<MWindow*>(ptr)->QMainWindow::paintEngine();
}

void* MWindow_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<MWindow*>(ptr)->QMainWindow::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* MWindow_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<MWindow*>(ptr)->QMainWindow::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* MWindow_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<MWindow*>(ptr)->QMainWindow::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char MWindow_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<MWindow*>(ptr)->QMainWindow::hasHeightForWidth();
}

int MWindow_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<MWindow*>(ptr)->QMainWindow::heightForWidth(w);
}

int MWindow_MetricDefault(void* ptr, long long m)
{
	return static_cast<MWindow*>(ptr)->QMainWindow::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

char MWindow_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<MWindow*>(ptr)->QMainWindow::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void MWindow_ChildEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::childEvent(static_cast<QChildEvent*>(event));
}

void MWindow_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<MWindow*>(ptr)->QMainWindow::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void MWindow_CustomEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::customEvent(static_cast<QEvent*>(event));
}

void MWindow_DeleteLaterDefault(void* ptr)
{
	static_cast<MWindow*>(ptr)->QMainWindow::deleteLater();
}

void MWindow_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<MWindow*>(ptr)->QMainWindow::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void MWindow_TimerEventDefault(void* ptr, void* event)
{
	static_cast<MWindow*>(ptr)->QMainWindow::timerEvent(static_cast<QTimerEvent*>(event));
}



void MocLabel_ConnectUpdateLabel(void* ptr)
{
	QObject::connect(static_cast<MocLabel*>(ptr), static_cast<void (MocLabel::*)(QString)>(&MocLabel::updateLabel), static_cast<MocLabel*>(ptr), static_cast<void (MocLabel::*)(QString)>(&MocLabel::Signal_UpdateLabel));
}

void MocLabel_DisconnectUpdateLabel(void* ptr)
{
	QObject::disconnect(static_cast<MocLabel*>(ptr), static_cast<void (MocLabel::*)(QString)>(&MocLabel::updateLabel), static_cast<MocLabel*>(ptr), static_cast<void (MocLabel::*)(QString)>(&MocLabel::Signal_UpdateLabel));
}

void MocLabel_UpdateLabel(void* ptr, char* v0)
{
	static_cast<MocLabel*>(ptr)->updateLabel(QString(v0));
}

int MocLabel_MocLabel_QRegisterMetaType()
{
	return qRegisterMetaType<MocLabel*>();
}

int MocLabel_MocLabel_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<MocLabel>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int MocLabel_MocLabel_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<MocLabel>();
#else
	return 0;
#endif
}

void* MocLabel___addActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void MocLabel___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MocLabel___addActions_actions_newList(void* ptr)
{
	return new QList<QAction *>;
}

void* MocLabel___insertActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void MocLabel___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MocLabel___insertActions_actions_newList(void* ptr)
{
	return new QList<QAction *>;
}

void* MocLabel___actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void MocLabel___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MocLabel___actions_newList(void* ptr)
{
	return new QList<QAction *>;
}

void* MocLabel___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void MocLabel___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* MocLabel___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* MocLabel___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void MocLabel___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MocLabel___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* MocLabel___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void MocLabel___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MocLabel___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* MocLabel___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void MocLabel___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MocLabel___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* MocLabel___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void MocLabel___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MocLabel___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

void* MocLabel_NewMocLabel(void* parent, long long fo)
{
		return new MocLabel(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(fo));
}

void* MocLabel_NewMocLabel2(char* text, void* parent, long long fo)
{
		return new MocLabel(QString(text), static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(fo));
}

void MocLabel_DestroyMocLabel(void* ptr)
{
	static_cast<MocLabel*>(ptr)->~MocLabel();
}

char MocLabel_EventDefault(void* ptr, void* e)
{
	return static_cast<MocLabel*>(ptr)->QLabel::event(static_cast<QEvent*>(e));
}

char MocLabel_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<MocLabel*>(ptr)->QLabel::focusNextPrevChild(next != 0);
}

void MocLabel_ChangeEventDefault(void* ptr, void* ev)
{
	static_cast<MocLabel*>(ptr)->QLabel::changeEvent(static_cast<QEvent*>(ev));
}

void MocLabel_ClearDefault(void* ptr)
{
	static_cast<MocLabel*>(ptr)->QLabel::clear();
}

void MocLabel_ContextMenuEventDefault(void* ptr, void* ev)
{
	static_cast<MocLabel*>(ptr)->QLabel::contextMenuEvent(static_cast<QContextMenuEvent*>(ev));
}

void MocLabel_FocusInEventDefault(void* ptr, void* ev)
{
	static_cast<MocLabel*>(ptr)->QLabel::focusInEvent(static_cast<QFocusEvent*>(ev));
}

void MocLabel_FocusOutEventDefault(void* ptr, void* ev)
{
	static_cast<MocLabel*>(ptr)->QLabel::focusOutEvent(static_cast<QFocusEvent*>(ev));
}

void MocLabel_KeyPressEventDefault(void* ptr, void* ev)
{
	static_cast<MocLabel*>(ptr)->QLabel::keyPressEvent(static_cast<QKeyEvent*>(ev));
}

void MocLabel_MouseMoveEventDefault(void* ptr, void* ev)
{
	static_cast<MocLabel*>(ptr)->QLabel::mouseMoveEvent(static_cast<QMouseEvent*>(ev));
}

void MocLabel_MousePressEventDefault(void* ptr, void* ev)
{
	static_cast<MocLabel*>(ptr)->QLabel::mousePressEvent(static_cast<QMouseEvent*>(ev));
}

void MocLabel_MouseReleaseEventDefault(void* ptr, void* ev)
{
	static_cast<MocLabel*>(ptr)->QLabel::mouseReleaseEvent(static_cast<QMouseEvent*>(ev));
}

void MocLabel_PaintEventDefault(void* ptr, void* vqp)
{
	static_cast<MocLabel*>(ptr)->QLabel::paintEvent(static_cast<QPaintEvent*>(vqp));
}

void MocLabel_SetMovieDefault(void* ptr, void* movie)
{
	static_cast<MocLabel*>(ptr)->QLabel::setMovie(static_cast<QMovie*>(movie));
}

void MocLabel_SetNum2Default(void* ptr, double num)
{
	static_cast<MocLabel*>(ptr)->QLabel::setNum(num);
}

void MocLabel_SetNumDefault(void* ptr, int num)
{
	static_cast<MocLabel*>(ptr)->QLabel::setNum(num);
}

void MocLabel_SetPictureDefault(void* ptr, void* picture)
{
	static_cast<MocLabel*>(ptr)->QLabel::setPicture(*static_cast<QPicture*>(picture));
}

void MocLabel_SetPixmapDefault(void* ptr, void* vqp)
{
	static_cast<MocLabel*>(ptr)->QLabel::setPixmap(*static_cast<QPixmap*>(vqp));
}

void MocLabel_SetTextDefault(void* ptr, char* vqs)
{
	static_cast<MocLabel*>(ptr)->QLabel::setText(QString(vqs));
}

void* MocLabel_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<MocLabel*>(ptr)->QLabel::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* MocLabel_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<MocLabel*>(ptr)->QLabel::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

int MocLabel_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<MocLabel*>(ptr)->QLabel::heightForWidth(w);
}

char MocLabel_CloseDefault(void* ptr)
{
	return static_cast<MocLabel*>(ptr)->QLabel::close();
}

char MocLabel_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
	return static_cast<MocLabel*>(ptr)->QLabel::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void MocLabel_ActionEventDefault(void* ptr, void* event)
{
	static_cast<MocLabel*>(ptr)->QLabel::actionEvent(static_cast<QActionEvent*>(event));
}

void MocLabel_CloseEventDefault(void* ptr, void* event)
{
	static_cast<MocLabel*>(ptr)->QLabel::closeEvent(static_cast<QCloseEvent*>(event));
}

void MocLabel_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<MocLabel*>(ptr)->QLabel::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void MocLabel_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<MocLabel*>(ptr)->QLabel::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void MocLabel_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<MocLabel*>(ptr)->QLabel::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void MocLabel_DropEventDefault(void* ptr, void* event)
{
	static_cast<MocLabel*>(ptr)->QLabel::dropEvent(static_cast<QDropEvent*>(event));
}

void MocLabel_EnterEventDefault(void* ptr, void* event)
{
	static_cast<MocLabel*>(ptr)->QLabel::enterEvent(static_cast<QEvent*>(event));
}

void MocLabel_HideDefault(void* ptr)
{
	static_cast<MocLabel*>(ptr)->QLabel::hide();
}

void MocLabel_HideEventDefault(void* ptr, void* event)
{
	static_cast<MocLabel*>(ptr)->QLabel::hideEvent(static_cast<QHideEvent*>(event));
}

void MocLabel_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<MocLabel*>(ptr)->QLabel::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void MocLabel_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<MocLabel*>(ptr)->QLabel::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void MocLabel_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<MocLabel*>(ptr)->QLabel::leaveEvent(static_cast<QEvent*>(event));
}

void MocLabel_LowerDefault(void* ptr)
{
	static_cast<MocLabel*>(ptr)->QLabel::lower();
}

void MocLabel_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<MocLabel*>(ptr)->QLabel::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void MocLabel_MoveEventDefault(void* ptr, void* event)
{
	static_cast<MocLabel*>(ptr)->QLabel::moveEvent(static_cast<QMoveEvent*>(event));
}

void MocLabel_RaiseDefault(void* ptr)
{
	static_cast<MocLabel*>(ptr)->QLabel::raise();
}

void MocLabel_RepaintDefault(void* ptr)
{
	static_cast<MocLabel*>(ptr)->QLabel::repaint();
}

void MocLabel_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<MocLabel*>(ptr)->QLabel::resizeEvent(static_cast<QResizeEvent*>(event));
}

void MocLabel_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<MocLabel*>(ptr)->QLabel::setDisabled(disable != 0);
}

void MocLabel_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<MocLabel*>(ptr)->QLabel::setEnabled(vbo != 0);
}

void MocLabel_SetFocus2Default(void* ptr)
{
	static_cast<MocLabel*>(ptr)->QLabel::setFocus();
}

void MocLabel_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<MocLabel*>(ptr)->QLabel::setHidden(hidden != 0);
}

void MocLabel_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<MocLabel*>(ptr)->QLabel::setStyleSheet(QString(styleSheet));
}

void MocLabel_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<MocLabel*>(ptr)->QLabel::setVisible(visible != 0);
}

void MocLabel_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<MocLabel*>(ptr)->QLabel::setWindowModified(vbo != 0);
}

void MocLabel_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<MocLabel*>(ptr)->QLabel::setWindowTitle(QString(vqs));
}

void MocLabel_ShowDefault(void* ptr)
{
	static_cast<MocLabel*>(ptr)->QLabel::show();
}

void MocLabel_ShowEventDefault(void* ptr, void* event)
{
	static_cast<MocLabel*>(ptr)->QLabel::showEvent(static_cast<QShowEvent*>(event));
}

void MocLabel_ShowFullScreenDefault(void* ptr)
{
	static_cast<MocLabel*>(ptr)->QLabel::showFullScreen();
}

void MocLabel_ShowMaximizedDefault(void* ptr)
{
	static_cast<MocLabel*>(ptr)->QLabel::showMaximized();
}

void MocLabel_ShowMinimizedDefault(void* ptr)
{
	static_cast<MocLabel*>(ptr)->QLabel::showMinimized();
}

void MocLabel_ShowNormalDefault(void* ptr)
{
	static_cast<MocLabel*>(ptr)->QLabel::showNormal();
}

void MocLabel_TabletEventDefault(void* ptr, void* event)
{
	static_cast<MocLabel*>(ptr)->QLabel::tabletEvent(static_cast<QTabletEvent*>(event));
}

void MocLabel_UpdateDefault(void* ptr)
{
	static_cast<MocLabel*>(ptr)->QLabel::update();
}

void MocLabel_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<MocLabel*>(ptr)->QLabel::updateMicroFocus();
}

void MocLabel_WheelEventDefault(void* ptr, void* event)
{
	static_cast<MocLabel*>(ptr)->QLabel::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* MocLabel_PaintEngineDefault(void* ptr)
{
	return static_cast<MocLabel*>(ptr)->QLabel::paintEngine();
}

void* MocLabel_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<MocLabel*>(ptr)->QLabel::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char MocLabel_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<MocLabel*>(ptr)->QLabel::hasHeightForWidth();
}

int MocLabel_MetricDefault(void* ptr, long long m)
{
	return static_cast<MocLabel*>(ptr)->QLabel::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

char MocLabel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<MocLabel*>(ptr)->QLabel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void MocLabel_ChildEventDefault(void* ptr, void* event)
{
	static_cast<MocLabel*>(ptr)->QLabel::childEvent(static_cast<QChildEvent*>(event));
}

void MocLabel_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<MocLabel*>(ptr)->QLabel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void MocLabel_CustomEventDefault(void* ptr, void* event)
{
	static_cast<MocLabel*>(ptr)->QLabel::customEvent(static_cast<QEvent*>(event));
}

void MocLabel_DeleteLaterDefault(void* ptr)
{
	static_cast<MocLabel*>(ptr)->QLabel::deleteLater();
}

void MocLabel_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<MocLabel*>(ptr)->QLabel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void MocLabel_TimerEventDefault(void* ptr, void* event)
{
	static_cast<MocLabel*>(ptr)->QLabel::timerEvent(static_cast<QTimerEvent*>(event));
}



#include "moc_moc.h"
