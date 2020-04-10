package backend

/*
#cgo CFLAGS: -fno-keep-inline-dllexport -O2 -Wall -W -Wextra -DUNICODE -D_UNICODE -DWIN32 -DMINGW_HAS_SECURE_API=1 -DQT_NEEDS_QMAIN -DQT_NO_DEBUG -DQT_QUICKCONTROLS2_LIB -DQT_QUICK_LIB -DQT_SVG_LIB -DQT_WIDGETS_LIB -DQT_GUI_LIB -DQT_QML_LIB -DQT_NETWORK_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -fno-keep-inline-dllexport -O2 -Wall -W -Wextra -fexceptions -mthreads -DUNICODE -D_UNICODE -DWIN32 -DMINGW_HAS_SECURE_API=1 -DQT_NEEDS_QMAIN -DQT_NO_DEBUG -DQT_QUICKCONTROLS2_LIB -DQT_QUICK_LIB -DQT_SVG_LIB -DQT_WIDGETS_LIB -DQT_GUI_LIB -DQT_QML_LIB -DQT_NETWORK_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I. -IC:/Qt/Qt5.12.7/5.12.7/mingw73_64/include -IC:/Qt/Qt5.12.7/5.12.7/mingw73_64/include/QtQuickControls2 -IC:/Qt/Qt5.12.7/5.12.7/mingw73_64/include/QtQuick -IC:/Qt/Qt5.12.7/5.12.7/mingw73_64/include/QtSvg -IC:/Qt/Qt5.12.7/5.12.7/mingw73_64/include/QtWidgets -IC:/Qt/Qt5.12.7/5.12.7/mingw73_64/include/QtGui -IC:/Qt/Qt5.12.7/5.12.7/mingw73_64/include/QtANGLE -IC:/Qt/Qt5.12.7/5.12.7/mingw73_64/include/QtQml -IC:/Qt/Qt5.12.7/5.12.7/mingw73_64/include/QtNetwork -IC:/Qt/Qt5.12.7/5.12.7/mingw73_64/include/QtCore -Irelease -I/include -IC:/Qt/Qt5.12.7/5.12.7/mingw73_64/mkspecs/win32-g++
#cgo LDFLAGS: -Wl,-s -Wl,-subsystem,windows -mthreads
#cgo LDFLAGS: C:/Qt/Qt5.12.7/5.12.7/mingw73_64/lib/libQt5QuickControls2.a C:/Qt/Qt5.12.7/5.12.7/mingw73_64/lib/libQt5Quick.a C:/Qt/Qt5.12.7/5.12.7/mingw73_64/lib/libQt5Svg.a C:/Qt/Qt5.12.7/5.12.7/mingw73_64/lib/libQt5Widgets.a C:/Qt/Qt5.12.7/5.12.7/mingw73_64/lib/libQt5Gui.a C:/Qt/Qt5.12.7/5.12.7/mingw73_64/lib/libQt5Qml.a C:/Qt/Qt5.12.7/5.12.7/mingw73_64/lib/libQt5Network.a C:/Qt/Qt5.12.7/5.12.7/mingw73_64/lib/libQt5Core.a  -lmingw32 C:/Qt/Qt5.12.7/5.12.7/mingw73_64/lib/libqtmain.a -LC:/openssl/lib -LC:/Utils/my_sql/mysql-5.6.11-winx64/lib -LC:/Utils/postgresql/pgsql/lib -lshell32
#cgo CFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
#cgo CXXFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
*/
import "C"