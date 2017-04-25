

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



#include "moc_moc.h"
