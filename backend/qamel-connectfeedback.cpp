#include <QObject>
#include <QQuickItem>
#include <QString>
#include <QByteArray>
#include <QQmlEngine>
#include <QMetaObject>

#include "_cgo_export.h"
#include "qamel-connectfeedback.h"

class ConnectFeedBack : public QQuickItem {
	Q_OBJECT

private:

public:
	ConnectFeedBack(QQuickItem* parent=Q_NULLPTR) : QQuickItem(parent) {
		qamelConnectFeedBackConstructor(this);
	}

	~ConnectFeedBack() {
		qamelDestroyConnectFeedBack(this);
	}

signals:
	void sendFansNums(int p0);

	void sendCompInfo(QString p0);

	void sendErr(int p0);

public slots:
	void receiveRoomID(int p0) {
		qamelConnectFeedBackReceiveRoomID(this, p0);
	}
};

void ConnectFeedBack_SendFansNums(void* ptr, int p0) {
	ConnectFeedBack *obj = static_cast<ConnectFeedBack*>(ptr);
	obj->sendFansNums(int(p0));
}

void ConnectFeedBack_SendCompInfo(void* ptr, char* p0) {
	ConnectFeedBack *obj = static_cast<ConnectFeedBack*>(ptr);
	obj->sendCompInfo(QString(p0));
}

void ConnectFeedBack_SendErr(void* ptr, int p0) {
	ConnectFeedBack *obj = static_cast<ConnectFeedBack*>(ptr);
	obj->sendErr(int(p0));
}

void ConnectFeedBack_RegisterQML(char* uri, int versionMajor, int versionMinor, char* qmlName) {
	qmlRegisterType<ConnectFeedBack>(uri, versionMajor, versionMinor, qmlName);
}

#include "moc-qamel-connectfeedback.h"
