/****************************************************************************
** Meta object code from reading C++ file 'qamel-handlemsg.cpp'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.12.7)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'qamel-handlemsg.cpp' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.12.7. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_HandleMsg_t {
    QByteArrayData data[14];
    char stringdata0[145];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_HandleMsg_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_HandleMsg_t qt_meta_stringdata_HandleMsg = {
    {
QT_MOC_LITERAL(0, 0, 9), // "HandleMsg"
QT_MOC_LITERAL(1, 10, 9), // "sendDanMu"
QT_MOC_LITERAL(2, 20, 0), // ""
QT_MOC_LITERAL(3, 21, 2), // "p0"
QT_MOC_LITERAL(4, 24, 8), // "sendGift"
QT_MOC_LITERAL(5, 33, 11), // "sendWelCome"
QT_MOC_LITERAL(6, 45, 16), // "sendWelComeGuard"
QT_MOC_LITERAL(7, 62, 16), // "sendGreatSailing"
QT_MOC_LITERAL(8, 79, 17), // "sendOnlineChanged"
QT_MOC_LITERAL(9, 97, 15), // "sendFansChanged"
QT_MOC_LITERAL(10, 113, 12), // "sendMusicURI"
QT_MOC_LITERAL(11, 126, 2), // "p1"
QT_MOC_LITERAL(12, 129, 2), // "p2"
QT_MOC_LITERAL(13, 132, 12) // "musicControl"

    },
    "HandleMsg\0sendDanMu\0\0p0\0sendGift\0"
    "sendWelCome\0sendWelComeGuard\0"
    "sendGreatSailing\0sendOnlineChanged\0"
    "sendFansChanged\0sendMusicURI\0p1\0p2\0"
    "musicControl"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_HandleMsg[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       9,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       8,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   59,    2, 0x06 /* Public */,
       4,    1,   62,    2, 0x06 /* Public */,
       5,    1,   65,    2, 0x06 /* Public */,
       6,    1,   68,    2, 0x06 /* Public */,
       7,    1,   71,    2, 0x06 /* Public */,
       8,    1,   74,    2, 0x06 /* Public */,
       9,    1,   77,    2, 0x06 /* Public */,
      10,    3,   80,    2, 0x06 /* Public */,

 // slots: name, argc, parameters, tag, flags
      13,    2,   87,    2, 0x0a /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::Int,    3,
    QMetaType::Void, QMetaType::Int,    3,
    QMetaType::Void, QMetaType::QString, QMetaType::QString, QMetaType::QString,    3,   11,   12,

 // slots: parameters
    QMetaType::Void, QMetaType::Bool, QMetaType::QString,    3,   11,

       0        // eod
};

void HandleMsg::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<HandleMsg *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->sendDanMu((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 1: _t->sendGift((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 2: _t->sendWelCome((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 3: _t->sendWelComeGuard((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 4: _t->sendGreatSailing((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 5: _t->sendOnlineChanged((*reinterpret_cast< int(*)>(_a[1]))); break;
        case 6: _t->sendFansChanged((*reinterpret_cast< int(*)>(_a[1]))); break;
        case 7: _t->sendMusicURI((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< QString(*)>(_a[3]))); break;
        case 8: _t->musicControl((*reinterpret_cast< bool(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (HandleMsg::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&HandleMsg::sendDanMu)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (HandleMsg::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&HandleMsg::sendGift)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (HandleMsg::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&HandleMsg::sendWelCome)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (HandleMsg::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&HandleMsg::sendWelComeGuard)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (HandleMsg::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&HandleMsg::sendGreatSailing)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (HandleMsg::*)(int );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&HandleMsg::sendOnlineChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (HandleMsg::*)(int );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&HandleMsg::sendFansChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (HandleMsg::*)(QString , QString , QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&HandleMsg::sendMusicURI)) {
                *result = 7;
                return;
            }
        }
    }
}

QT_INIT_METAOBJECT const QMetaObject HandleMsg::staticMetaObject = { {
    &QQuickItem::staticMetaObject,
    qt_meta_stringdata_HandleMsg.data,
    qt_meta_data_HandleMsg,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *HandleMsg::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *HandleMsg::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_HandleMsg.stringdata0))
        return static_cast<void*>(this);
    return QQuickItem::qt_metacast(_clname);
}

int HandleMsg::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QQuickItem::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 9)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 9;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 9)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 9;
    }
    return _id;
}

// SIGNAL 0
void HandleMsg::sendDanMu(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void HandleMsg::sendGift(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void HandleMsg::sendWelCome(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void HandleMsg::sendWelComeGuard(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void HandleMsg::sendGreatSailing(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void HandleMsg::sendOnlineChanged(int _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void HandleMsg::sendFansChanged(int _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void HandleMsg::sendMusicURI(QString _t1, QString _t2, QString _t3)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)), const_cast<void*>(reinterpret_cast<const void*>(&_t2)), const_cast<void*>(reinterpret_cast<const void*>(&_t3)) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
