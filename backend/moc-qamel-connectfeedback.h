/****************************************************************************
** Meta object code from reading C++ file 'qamel-connectfeedback.cpp'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.12.7)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'qamel-connectfeedback.cpp' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.12.7. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_ConnectFeedBack_t {
    QByteArrayData data[7];
    char stringdata0[68];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_ConnectFeedBack_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_ConnectFeedBack_t qt_meta_stringdata_ConnectFeedBack = {
    {
QT_MOC_LITERAL(0, 0, 15), // "ConnectFeedBack"
QT_MOC_LITERAL(1, 16, 12), // "sendFansNums"
QT_MOC_LITERAL(2, 29, 0), // ""
QT_MOC_LITERAL(3, 30, 2), // "p0"
QT_MOC_LITERAL(4, 33, 12), // "sendCompInfo"
QT_MOC_LITERAL(5, 46, 7), // "sendErr"
QT_MOC_LITERAL(6, 54, 13) // "receiveRoomID"

    },
    "ConnectFeedBack\0sendFansNums\0\0p0\0"
    "sendCompInfo\0sendErr\0receiveRoomID"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_ConnectFeedBack[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       4,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       3,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   34,    2, 0x06 /* Public */,
       4,    1,   37,    2, 0x06 /* Public */,
       5,    1,   40,    2, 0x06 /* Public */,

 // slots: name, argc, parameters, tag, flags
       6,    1,   43,    2, 0x0a /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::Int,    3,
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::Int,    3,

 // slots: parameters
    QMetaType::Void, QMetaType::Int,    3,

       0        // eod
};

void ConnectFeedBack::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<ConnectFeedBack *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->sendFansNums((*reinterpret_cast< int(*)>(_a[1]))); break;
        case 1: _t->sendCompInfo((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 2: _t->sendErr((*reinterpret_cast< int(*)>(_a[1]))); break;
        case 3: _t->receiveRoomID((*reinterpret_cast< int(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (ConnectFeedBack::*)(int );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ConnectFeedBack::sendFansNums)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (ConnectFeedBack::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ConnectFeedBack::sendCompInfo)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (ConnectFeedBack::*)(int );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ConnectFeedBack::sendErr)) {
                *result = 2;
                return;
            }
        }
    }
}

QT_INIT_METAOBJECT const QMetaObject ConnectFeedBack::staticMetaObject = { {
    &QQuickItem::staticMetaObject,
    qt_meta_stringdata_ConnectFeedBack.data,
    qt_meta_data_ConnectFeedBack,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *ConnectFeedBack::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *ConnectFeedBack::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_ConnectFeedBack.stringdata0))
        return static_cast<void*>(this);
    return QQuickItem::qt_metacast(_clname);
}

int ConnectFeedBack::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QQuickItem::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 4)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 4;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 4)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 4;
    }
    return _id;
}

// SIGNAL 0
void ConnectFeedBack::sendFansNums(int _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void ConnectFeedBack::sendCompInfo(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void ConnectFeedBack::sendErr(int _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
