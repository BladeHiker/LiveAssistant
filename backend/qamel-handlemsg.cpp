#include <QObject>
#include <QQuickItem>
#include <QString>
#include <QByteArray>
#include <QQmlEngine>
#include <QMetaObject>

#include "_cgo_export.h"
#include "qamel-handlemsg.h"

class HandleMsg : public QQuickItem {
	Q_OBJECT

private:

public:
	HandleMsg(QQuickItem* parent=Q_NULLPTR) : QQuickItem(parent) {
		qamelHandleMsgConstructor(this);
	}

	~HandleMsg() {
		qamelDestroyHandleMsg(this);
	}

signals:
	void sendDanMu(QString p0);

	void sendGift(QString p0);

	void sendWelCome(QString p0);

	void sendWelComeGuard(QString p0);

	void sendGreatSailing(QString p0);

	void sendOnlineChanged(int p0);

	void sendFansChanged(int p0);

	void sendMusicURI(QString p0, QString p1, QString p2);

public slots:
	void musicControl(bool p0, QString p1) {
		qamelHandleMsgMusicControl(this, p0, p1.toLocal8Bit().data());
	}
};

void HandleMsg_SendDanMu(void* ptr, char* p0) {
	HandleMsg *obj = static_cast<HandleMsg*>(ptr);
	obj->sendDanMu(QString(p0));
}

void HandleMsg_SendGift(void* ptr, char* p0) {
	HandleMsg *obj = static_cast<HandleMsg*>(ptr);
	obj->sendGift(QString(p0));
}

void HandleMsg_SendWelCome(void* ptr, char* p0) {
	HandleMsg *obj = static_cast<HandleMsg*>(ptr);
	obj->sendWelCome(QString(p0));
}

void HandleMsg_SendWelComeGuard(void* ptr, char* p0) {
	HandleMsg *obj = static_cast<HandleMsg*>(ptr);
	obj->sendWelComeGuard(QString(p0));
}

void HandleMsg_SendGreatSailing(void* ptr, char* p0) {
	HandleMsg *obj = static_cast<HandleMsg*>(ptr);
	obj->sendGreatSailing(QString(p0));
}

void HandleMsg_SendOnlineChanged(void* ptr, int p0) {
	HandleMsg *obj = static_cast<HandleMsg*>(ptr);
	obj->sendOnlineChanged(int(p0));
}

void HandleMsg_SendFansChanged(void* ptr, int p0) {
	HandleMsg *obj = static_cast<HandleMsg*>(ptr);
	obj->sendFansChanged(int(p0));
}

void HandleMsg_SendMusicURI(void* ptr, char* p0, char* p1, char* p2) {
	HandleMsg *obj = static_cast<HandleMsg*>(ptr);
	obj->sendMusicURI(QString(p0), QString(p1), QString(p2));
}

void HandleMsg_RegisterQML(char* uri, int versionMajor, int versionMinor, char* qmlName) {
	qmlRegisterType<HandleMsg>(uri, versionMajor, versionMinor, qmlName);
}

#include "moc-qamel-handlemsg.h"
