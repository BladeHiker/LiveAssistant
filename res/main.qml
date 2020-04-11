import QtQuick 2.12
import QtQuick.Controls 2.5
import QtMultimedia 5.12
import QtGraphicalEffects 1.0
import ConnectFeedBack 1.0
import HandleMsg 1.0
import Qt.labs.platform 1.1

import "fonts/FontAwesome" as FA
import "components"

ApplicationWindow {
    id: mainwindow
    visible: true
    width: 250
    height: 300
    property int roomID
    property int linked: 1
    property int cnt: 0
    property bool stayTop: true
    opacity: 1
    flags: Qt.Window | Qt.FramelessWindowHint | Qt.WindowStaysOnTopHint

    color: "#00000000"
    background: Rectangle {
        color: "#BB5B6B63"
        radius: 5
        height: mainwindow.height
        width: mainwindow.width
        anchors.fill: parent
    }

    property var online: "-"
    property var fans: "-"

    SystemTrayIcon {
        id: trayIcon
        visible: true
        icon.mask: true
        iconSource: "qrc:/res/icon.png"
        icon.name: "LiveAssistant"
        //        Component.onCompleted: showMessage("LiveAssistant", "LiveAssistant已启动")
        menu: Menu {
            MenuItem {
                visible: roomID != 0
                text: "直播间" + roomID
            }

            MenuItem {
                text: "退出"
                onTriggered: Qt.quit()
            }
        }
    }

    ListModel {
        id: dm
        onCountChanged: {
            if (count > 100) {
                remove(0, 50)
            }
        }
    }
    ListModel {
        id: gf
        onCountChanged: {
            if (count > 60) {
                remove(0, 30)
            }
        }
    }
    ConnectFeedBack {
        id: connectFeedBack
        onSendFansNums: function (n) {
            fans = n
        }
        //        onSendCompInfo: function (m) {//TODO
        //        }
        signal startState(var s)
        onSendErr: function (e) {
            if (e === 0) {
                mainwindow.height = 500
                mainwindow.width = 300
                mainloader.sourceComponent = main
            } else if (e === -1) {
                startState("inp")
                //                roomidArea.state = "inp"
            }
        }
    }

    Behavior on height {
        NumberAnimation {
            duration: 500
        }
    }
    Behavior on width {
        NumberAnimation {
            duration: 500
        }
    }

    Loader {
        id: mainloader
        sourceComponent: start
    }
    Button {
        id: closeBtn
        background: Rectangle {
            anchors.fill: parent
            radius: width
            Text {
                id: xbtn
                text: FA.Icons.faTimes
                font.family: FA.Fonts.solid
                anchors.horizontalCenter: parent.horizontalCenter
                anchors.verticalCenter: parent.verticalCenter
                font.bold: true
                font.pixelSize: 10
                visible: false
                anchors.centerIn: parent.Center
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
            }
            color: "#fc625d"
            MouseArea {
                anchors.fill: parent
                hoverEnabled: true
                onEntered: {
                    xbtn.visible = true
                    xbtnon.start()
                }
                onClicked: {
                    trayIcon.hide()
                    Qt.quit()
                }
                onExited: {

                    xbtn.visible = false
                }

                NumberAnimation {
                    id: xbtnon
                    target: xbtn
                    property: "rotation"
                    from: -30
                    to: 0
                    duration: 200
                    easing.type: Easing.InOutQuad
                }
            }
        }
        anchors.top: parent.top
        anchors.right: parent.right
        width: 15
        height: 15
        z: 100
    }
    Button {
        id: topBtn
        background: Rectangle {
            anchors.fill: parent
            radius: width
            Text {
                id: topSym
                text: FA.Icons.faThumbtack
                font.family: FA.Fonts.solid
                anchors.horizontalCenter: parent.horizontalCenter
                anchors.verticalCenter: parent.verticalCenter
                //                font.bold: true
                font.pixelSize: 8
                //                visible: false
                anchors.centerIn: parent.Center
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
                rotation: 0
            }
            color: "#fdbc40"
            MouseArea {
                anchors.fill: parent
                hoverEnabled: true
                onClicked: {
                    if (stayTop) {
                        mainwindow.flags = Qt.Window | Qt.FramelessWindowHint

                        stayTop = false
                        topSym.rotation = 30
                    } else {
                        mainwindow.flags = Qt.Window | Qt.FramelessWindowHint
                                | Qt.WindowStaysOnTopHint
                        topSym.rotation = 0
                        stayTop = true
                    }
                }
            }
        }
        anchors.top: parent.top
        anchors.right: parent.right
        anchors.rightMargin: 20
        width: 15
        height: 15
        z: 100
    }
    Component {
        id: start
        Rectangle {
            visible: true
            width: 250
            height: 300
            radius: 10
            Rectangle {
                id: logoarea
                width: parent.width
                height: 100
                anchors.top: parent.top
                color: "#32c7ff"
                clip: true
                Image {
                    id: panel
                    width: parent.width
                    height: parent.height
                    anchors.centerIn: parent.Center
                    source: "qrc:/res/panel.png"
                }
                Text {
                    text: "V1.0 Alpha"
                    verticalAlignment: Text.AlignVCenter
                    font.family: "Arial"
                    color: "#aaaaaa"
                    font.bold: false
                    horizontalAlignment: Text.AlignHCenter
                    font.pointSize: 8
                    anchors.bottom: parent.bottom
                    anchors.bottomMargin: 2
                    anchors.horizontalCenter: parent.horizontalCenter
                }

                MouseArea {
                    z: 1
                    width: parent.width
                    height: parent.height
                    anchors.top: parent.top
                    property point prePoint: Qt.point(0, 0)
                    property int datX: 0
                    property int datY: 0

                    onPressed: {
                        prePoint = Qt.point(mouseX, mouseY)
                    }

                    onPositionChanged: {
                        if (pressed
                                && !(datX == -1 * (mouseX - prePoint.x)
                                     && datY == -1 * (mouseY - prePoint.y))) {
                            datX = mouseX - prePoint.x
                            datY = mouseY - prePoint.y
                            mainwindow.x += datX
                            mainwindow.y += datY
                        }
                    }
                }
            }
            Rectangle {
                id: roomidArea
                anchors.top: logoarea.bottom
                height: 200
                width: parent.width
                color: "#eee"
                Connections {
                    target: connectFeedBack
                    onStartState: function (s) {
                        roomidArea.state = s
                    }
                }
                Keys.enabled: true
                Keys.onPressed: {
                    if (!connBtn.enabled)
                        return
                    if (event.key === Qt.Key_Enter
                            || event.key === Qt.Key_Return) {
                        roomidArea.state = "conn"
                        roomID = Number(roomidInp.text)
                        connectFeedBack.receiveRoomID(roomID)
                    }
                    event.accepted = true
                }

                states: [
                    State {
                        name: "inp"
                        PropertyChanges {
                            target: connBtn
                            width: 200
                        }
                        PropertyChanges {
                            target: roomidTip
                            visible: true
                            text: "连接失败！请再试一次"
                        }
                        PropertyChanges {
                            target: connBtnText
                            text: "连接直播间"
                            font.family: "黑体"
                            rotation: 0
                        }
                        PropertyChanges {
                            target: connLoadingAni
                            running: false
                        }
                        PropertyChanges {
                            target: roomidInp
                            enabled: true
                            background.visible: true
                            height: 35
                            width: 200
                            font.pixelSize: 15
                            anchors.topMargin: 40
                        }
                    },
                    State {
                        name: "conn"
                        PropertyChanges {
                            target: connBtn
                            width: 35
                        }
                        PropertyChanges {
                            target: roomidTip

                            visible: true
                            text: "正在连接"
                        }
                        PropertyChanges {
                            target: connBtnText
                            text: FA.Icons.faCircleNotch
                            font.family: FA.Fonts.solid
                        }
                        PropertyChanges {
                            target: connLoadingAni
                            running: true
                        }
                        PropertyChanges {
                            target: roomidInp
                            enabled: false
                            background.visible: false
                            height: 60
                            width: 300
                            font.pixelSize: 40
                            anchors.topMargin: 25
                        }
                    }
                ]
                TextField {
                    id: roomidInp
                    Component.onCompleted: {
                        roomidInp.forceActiveFocus()
                    }

                    horizontalAlignment: Text.AlignHCenter
                    selectByMouse: true
                    focus: true
                    anchors.top: parent.top
                    font.pixelSize: 15
                    anchors.topMargin: 40
                    height: 35
                    width: 200
                    placeholderText: "请输入直播间房间号"
                    font.family: "黑体"
                    anchors.horizontalCenter: parent.horizontalCenter
                    color: {
                        if (text != "" && /^\d+$/.test(text) == false) {
                            return "#ff0000"
                        }
                        return "#000000"
                    }
                    background: Rectangle {
                        radius: 25
                        border.color: roomidInp.enabled ? "#7a5fee" : "transparent"
                    }
                    Behavior on anchors.topMargin {
                        NumberAnimation {
                            alwaysRunToEnd: true
                            duration: 500
                        }
                    }
                    Behavior on font.pixelSize {
                        NumberAnimation {
                            alwaysRunToEnd: true
                            duration: 200
                            easing: Easing.InQuart
                        }
                    }
                }
                Text {
                    id: roomidTip
                    font.family: "黑体"
                    font.pointSize: 10
                    horizontalAlignment: Text.AlignHCenter
                    verticalAlignment: Text.AlignVCenter
                    visible: false
                    anchors.bottom: connBtn.top
                    anchors.bottomMargin: 5
                    anchors.horizontalCenter: parent.horizontalCenter
                }
                Button {
                    z: 100
                    id: connBtn

                    enabled: ((/^\d+$/.test(roomidInp.text))
                              && (roomidInp.text !== ""))
                    anchors.bottom: parent.bottom
                    anchors.bottomMargin: 40
                    anchors.horizontalCenter: parent.horizontalCenter
                    focus: true
                    height: 35
                    width: 200
                    onPressed: {
                        //多线程解决就不需要这个了
                        roomidArea.state = "conn"
                    }
                    Behavior on width {
                        NumberAnimation {
                            alwaysRunToEnd: true
                            duration: 500
                        }
                    }

                    Text {
                        id: connBtnText
                        anchors.centerIn: parent
                        color: "#ffffff"
                        text: "连接直播间"
                        font.family: "黑体"
                        font.pointSize: 12
                        horizontalAlignment: Text.AlignHCenter
                        verticalAlignment: Text.AlignVCenter
                        RotationAnimation on rotation {
                            id: connLoadingAni
                            running: false
                            loops: Animation.Infinite
                            from: 0
                            to: 360
                            duration: 1000
                            alwaysRunToEnd: false
                        }
                    }
                    background: Rectangle {
                        radius: 25
                        implicitWidth: 200
                        implicitHeight: 35
                        LinearGradient {
                            visible: false
                            id: gradCol1
                            anchors.fill: parent
                            start: Qt.point(0, 0)
                            end: Qt.point(200, 35)
                            gradient: Gradient {
                                GradientStop {
                                    position: 0.0
                                    color: "#4d32e1"
                                }
                                GradientStop {
                                    position: 1.0
                                    color: "#6914ca"
                                }
                            }
                        }
                        OpacityMask {
                            anchors.fill: gradCol1
                            source: gradCol1
                            maskSource: parent
                        }

                        border.color: roomidInp.enabled ? "#7a5fee" : "transparent"
                    }

                    onClicked: {
                        //roomidArea.state = "conn"
                        roomID = Number(roomidInp.text)
                        connectFeedBack.receiveRoomID(roomID)
                    }
                }
            }
        }
    }
    HandleMsg {
        id: handleMsg
        onSendDanMu: function (m) {
            var d = JSON.parse(m)
            dm.append({
                          "type": 1,
                          "avatar": d.avatar,
                          "utitle": d.utitle,
                          "title": "",
                          "uname": d.uname,
                          "msg": d.text
                      })
        }
        onSendGift: function (m) {
            var d = JSON.parse(m)
            gf.append({
                          "uname": d.uname,
                          "gift_action": d.action,
                          "nums": d.nums,
                          "gift_name": d.gname,
                          "price": d.price
                      })
        }
        onSendWelCome: function (m) {
            var d = JSON.parse(m)
            dm.append({
                          "type": 2,
                          "avatar": "",
                          "utitle": "",
                          "title": d.title,
                          "uname": d.uname,
                          "msg": ""
                      })
        }
        onSendWelComeGuard: function (m) {//                var d = JSON.parse(m)
            //                dm.append({
            //                              "type": 2,
            //                              "avatar": "",
            //                              "utitle": "",
            //                              "title": d.title,
            //                              "uname": d.uname,
            //                              "msg": ""
            //                          })
        }
        onSendGreatSailing: function (m) {
            var d = JSON.parse(m)
            dm.append({
                          "type": 2,
                          "avatar": "",
                          "utitle": "",
                          "title": d.title,
                          "uname": d.uname,
                          "msg": ""
                      })
        }
        onSendOnlineChanged: function (m) {
            online = m
        }
        onSendFansChanged: function (m) {
            fans = m
        }
    }
    Component {
        id: main
        Rectangle {
            visible: true
            width: 300
            height: 500
            //            color: "#40002010"
            color: "#00000000"
            radius: 5
            Connections {
                target: handleMsg
                onSendMusicURI: function (u, n, s) {
                    musicbox.addMusic(u, n, s)
                }
            }
            Column {
                id: column
                spacing: 1
                anchors.fill: parent
                width: parent.width
                height: parent.height
                anchors.margins: 5
                Musicbox {
                    id: musicbox
                    z: 100
                    musicenable: false
                    height: {
                        if (musicenable)
                            return 60
                        else
                            return 0
                    }
                    width: parent.width
                    onMusicenableChanged: {
                        if (!musicenable)
                            musicSeek = false
                    }

                    onMusicSeekChanged: {
                        handleMsg.musicControl(musicSeek, musicbox.key)
                        dm.append({
                                      "type": 0,
                                      "avatar": "",
                                      "utitle": "",
                                      "title": "",
                                      "uname": "",
                                      "msg": musicSeek ? "点歌已启用" : "点歌已关闭"
                                  })
                    }
                    onKeyChanged: {
                        handleMsg.musicControl(musicSeek, musicbox.key)
                    }
                }
                Rectangle {
                    id: line0
                    visible: musicbox.musicenable
                    width: parent.width
                    height: 1
                    color: "#aaeeeeee"
                }
                Item {
                    id: infobox
                    width: parent.width
                    height: 14
                    Row {
                        id: infoview
                        anchors.fill: parent
                        spacing: 10
                        Text {
                            color: "#FFFFFF"
                            text: '<font color="#FF93D3">' + FA.Icons.faCube
                                  + '</font>' + ' ' + roomID
                            anchors.verticalCenter: parent.verticalCenter
                            font.family: FA.Fonts.solid
                            verticalAlignment: Text.AlignVCenter
                            font.pixelSize: 11
                        }
                        Text {
                            color: "#FFFFFF"
                            text: '<font color="#23ade5">' + FA.Icons.faUsers
                                  + '</font>' + ' ' + online
                            anchors.verticalCenter: parent.verticalCenter
                            font.family: FA.Fonts.solid
                            verticalAlignment: Text.AlignVCenter
                            font.pixelSize: 11
                        }
                        Text {
                            color: "#FFFFFF"
                            text: '<font color="#fb7299">' + FA.Icons.faHeart
                                  + '</font>' + ' ' + fans
                            anchors.verticalCenter: parent.verticalCenter
                            font.family: FA.Fonts.solid
                            verticalAlignment: Text.AlignVCenter
                            font.pixelSize: 11
                        }
                        // Text {
                        //     id: testmsg
                        //     // text: dm.count + "/" + gf.count
                        //     color: "#000000"
                        //     anchors.verticalCenter: parent.verticalCenter
                        //     font.family: "微软雅黑"
                        //     verticalAlignment: Text.AlignVCenter
                        //     font.pixelSize: 11
                        // }
                    }
                    //                    Glow {
                    //                        source: infoview
                    //                        anchors.fill: infoview
                    //                        radius: 1
                    //                        samples: 20
                    //                        spread: 1
                    //                        color: "white"
                    //                    }
                }

                Rectangle {
                    id: line1
                    width: parent.width
                    height: 1
                    color: "#aaeeeeee"
                }
                ScrollView {
                    id: dmview
                    width: parent.width
                    height: {
                        if (musicbox.musicenable)
                            return 310
                        else
                            return 370
                    }
                    clip: true
                    ListView {
                        id: dmbox
                        width: parent.width
                        anchors.margins: 5
                        snapMode: ListView.NoSnap
                        boundsBehavior: Flickable.StopAtBounds
                        model: dm
                        spacing: 2
                        delegate: Danmu {}

                        add: Transition {
                            ParallelAnimation {
                                NumberAnimation {
                                    property: "opacity"
                                    from: 0
                                    to: 1.0
                                    duration: 500
                                }

                                NumberAnimation {
                                    property: "x"
                                    from: 20
                                    duration: 800
                                }
                            }
                        }
                        displaced: Transition {
                            SpringAnimation {
                                property: "y"
                                spring: 2
                                damping: 0.5
                                epsilon: 200
                            }
                        }
                        populate: Transition {
                            NumberAnimation {
                                property: "opacity"
                                from: 0
                                to: 1.0
                                duration: 500
                            }
                        }

                        property int dmcnt: 0

                        onContentHeightChanged: {
                            if (dmcnt == dm.count)
                                return
                            else
                                dmcnt = dm.count
                            dmbox.positionViewAtEnd()
                        }
                    }
                    Behavior on height {
                        NumberAnimation {
                            duration: 200
                            easing.type: Easing.InOutQuad
                        }
                    }
                }
                Rectangle {
                    id: line2
                    width: parent.width
                    height: 1
                    color: "#aaeeeeee"
                }
                ScrollView {
                    id: giftview
                    width: parent.width
                    height: 100
                    clip: true
                    ListView {
                        id: giftbox
                        spacing: 2
                        width: parent.width
                        model: gf
                        delegate: Gift {}

                        add: Transition {
                            ParallelAnimation {
                                NumberAnimation {
                                    property: "opacity"
                                    from: 0
                                    to: 1.0
                                    duration: 500
                                }
                                NumberAnimation {
                                    property: "scale"
                                    from: 1.5
                                    to: 1.0
                                    duration: 200
                                }
                                NumberAnimation {
                                    property: "rotation"
                                    from: 30
                                    to: 0
                                    duration: 300
                                }
                                NumberAnimation {
                                    property: "x"
                                    from: 20
                                    duration: 800
                                }
                            }
                        }
                    }
                    onContentHeightChanged: {
                        giftbox.positionViewAtEnd()
                    }
                }
            }

            Button {
                background: Rectangle {
                    anchors.fill: parent
                    radius: width
                    Text {
                        id: bbtn
                        text: musicbox.visible ? FA.Icons.faChevronUp : FA.Icons.faMusic
                        font.family: FA.Fonts.solid
                        //                        color: "#0772A1"
                        color: "#000000"
                        anchors.horizontalCenter: parent.horizontalCenter
                        anchors.verticalCenter: parent.verticalCenter
                        font.bold: true
                        font.pixelSize: 10
                        visible: false
                        anchors.centerIn: parent.Center
                        verticalAlignment: Text.AlignVCenter
                        horizontalAlignment: Text.AlignHCenter
                    }
                    color: "#35cd4b"
                    MouseArea {
                        anchors.fill: parent
                        hoverEnabled: true
                        onEntered: {
                            bbtn.visible = true
                            bbtnon.start()
                        }
                        onClicked: {
                            if (musicbox.musicenable == true) {
                                // handleMsg.musicControl(false,musicbox.key)
                                musicbox.musicenable = false
                            } else {
                                // handleMsg.musicControl(true,musicbox.key)
                                musicbox.musicenable = true
                            }
                        }
                        onExited: {
                            bbtn.visible = false
                        }

                        NumberAnimation {
                            id: bbtnon
                            target: bbtn
                            property: "rotation"
                            from: -90
                            to: 0
                            duration: 200
                            easing.type: Easing.InOutQuad
                        }
                    }
                }
                anchors.top: parent.top
                anchors.right: parent.right
                anchors.rightMargin: 40
                width: 15
                height: 15
                z: 100
            }
            MouseArea {
                id: dragRegion

                anchors.fill: parent
                anchors.top: parent.top
                anchors.topMargin: {
                    if (musicbox.musicenable)
                        return 60
                    return 0
                }
                property point prePoint: Qt.point(0, 0)
                property int datX: 0
                property int datY: 0

                onPressed: {
                    prePoint = Qt.point(mouseX, mouseY)
                }

                onPositionChanged: {
                    if (pressed && !(datX == -1 * (mouseX - prePoint.x)
                                     && datY == -1 * (mouseY - prePoint.y))) {
                        datX = mouseX - prePoint.x
                        datY = mouseY - prePoint.y
                        mainwindow.x += datX
                        mainwindow.y += datY
                    }
                }
                // onDoubleClicked: {
                //     testmsg.text = "clk"
                //     handleMsg.musicControl(true, "点歌")
                //     testmsg.text = musicbox.key
                // }
            }
        }
    }
}
