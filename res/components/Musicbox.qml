import QtQuick 2.12
import QtQuick.Controls 2.5
import QtMultimedia 5.12
import QtGraphicalEffects 1.0
import QtQuick.Controls.Styles 1.4
import "../fonts/FontAwesome" as FA

Item {
    id: element
    //    id: musicbox
    width: 300
    height: 60
    visible: musicenable
    property string key: "M"
    property bool musicSeek: false
    property bool musicenable: false
    property bool settingmode: settingPop.opened
    onVisibleChanged: {
        if (visible)
            settingPop.open()
    }

    function addMusic(url, name, singer) {
        if (playlist.itemCount !== 0 && url === playlist.itemSource(
                    playlist.itemCount - 1))
            return
        player.infolst.push([url, name, singer])
        playlist.addItem(url)
        if (player.playbackState != Audio.PausedState)
            player.play()
        textAni1.start()
    }
    function findMusicInfo(url) {
        var i
        for (i = 0; i < player.infolst.length; i++) {
            if (player.infolst[i][0] == url) {
                return i
            }
        }
        return -1
    }
    Audio {
        id: player
        volume: volSlider.value
        onStatusChanged: {
            if (playbackState != Audio.PlayingState
                    && playlist.itemCount != 0) {
                play()
            }

            if (status == Audio.EndOfMedia) {
                playlist.removeItem(0)
            }
        }
        onMediaObjectChanged: {
            if (player.mediaObject) {
                player.mediaObject.notifyInterval = 20
            }
        }
        property var infolst: []

        playlist: Playlist {
            id: playlist
            onCurrentItemSourceChanged: {
                if (itemCount === 0)
                    return
                var i = findMusicInfo(currentItemSource)
                if (i !== -1) {
                    nowplaying.text = player.infolst[i][1]
                    singer.text = player.infolst[i][2]
                } else {
                    nowplaying.text = "NULL"
                    singer.text = "NULL"
                }
            }
            onItemCountChanged: {
                if (itemCount === 0) {
                    if (musicSeek) {
                        nowplaying.text = tips0.text
                        singer.text = ""
                    } else {
                        nowplaying.text = ""
                        singer.text = ""
                    }
                }
            }
        }
    }
    Popup {
        id: settingPop
        width: parent.width
        height: parent.height
        anchors.centerIn: parent
        TextField {
            id: keyInp
            width: parent.width
            height: 20
            anchors.top: parent.Top
            hoverEnabled: true
            text: key
            horizontalAlignment: Text.AlignHCenter
            bottomPadding: 1
            topPadding: 1
            font.pixelSize: 10
            font.family: "微软雅黑"
            placeholderText: "输入点歌关键字"
            onTextChanged: {
                if (text.length != 0 && text.length < 6) {
                    keyEnter.text = '使用"' + keyInp.text + '+歌名"来点歌'
                    keyEnter.enabled = true
                } else if (text.length == 0) {
                    keyEnter.text = "关键字不能为空"
                    keyEnter.enabled = false
                } else {
                    keyEnter.text = "关键字过长"
                    keyEnter.enabled = false
                }
            }
        }
        Button {
            id: keyEnter
            width: parent.width
            height: 15
            font.family: "微软雅黑"
            anchors.top: keyInp.bottom
            onClicked: {
                key = keyInp.text
                musicSeek = true
                settingPop.close()
            }
        }
    }
    Text {
        id: singer
        anchors.bottom: nowplaying.top
        anchors.bottomMargin: 1
        verticalAlignment: Text.AlignVCenter
        horizontalAlignment: Text.AlignHCenter
        anchors.horizontalCenter: parent.horizontalCenter
        font.pixelSize: 10
        font.family: "思源黑体 Light"
        layer.enabled: true
        layer.effect: Glow {
            color: "white"
            samples: 20
            spread: 0.9
            radius: 4
        }
    }
    Text {
        id: nowplaying
        width: parent.width
        text: "Welcome"
        height: 23
        font.family: "思源黑体 Light"
        anchors.bottom: playProgress.top
        anchors.bottomMargin: 1
        anchors.horizontalCenter: parent.horizontalCenter
        verticalAlignment: Text.AlignVCenter
        horizontalAlignment: Text.AlignHCenter
        font.pixelSize: 20
        layer.enabled: true
        layer.effect: Glow {
            color: "white"
            samples: 25
            radius: 6
            spread: 0.9
        }
    }
    ProgressBar {
        id: playProgress
        width: 300
        height: 1.5
        background: Rectangle {
            z: 1
            implicitHeight: 1
            implicitWidth: 300
            anchors.verticalCenter: playProgress.verticalCenter
            anchors.left: playProgress.left
            color: "#e6e6e6"
        }
        contentItem: Rectangle {
            anchors.left: control.left
            anchors.verticalCenter: playProgress.verticalCenter
            width: playProgress.visualPosition * playProgress.width
            height: playProgress.height
            color: "#1573e6"
            z: 2
        }
        anchors.bottom: infoRow.top
        anchors.bottomMargin: 3
        opacity: 0.8
        indeterminate: (player.status == Audio.Loading || player.duration == 0)
                       && (player.playbackState == Audio.PlayingState)
        anchors.horizontalCenterOffset: 0
        anchors.horizontalCenter: parent.horizontalCenter
        value: player.position / player.duration
    }
    Row {
        id: infoRow
        width: parent.width
        height: 15
        anchors.horizontalCenter: parent.horizontalCenter
        anchors.bottom: parent.bottom
        anchors.bottomMargin: 1
        spacing: 5
        Text {
            id: tips0
            font.family: "微软雅黑"
            verticalAlignment: Text.AlignVCenter
            visible: key !== "" && musicSeek
            text: '发弹幕"' + key + '+歌名"来点歌'
            font.pixelSize: 12
        }

        Text {
            id: info1
            text: qsTr("当前列表：" + playlist.itemCount + "首")
            font.pixelSize: 12
            verticalAlignment: Text.AlignVCenter
            font.family: "微软雅黑"
            ColorAnimation on color {
                id: textAni1
                from: "red"
                to: "black"
                duration: 500
            }
            color: "black"
        }

        layer.enabled: true
        layer.effect: Glow {
            color: "white"
            samples: 20
            radius: 2
            spread: 1
        }
    }
    Text {
        id: playSym
        anchors.right: parent.right
        font.family: FA.Fonts.solid
        font.pixelSize: 13
        text: {
            switch (player.playbackState) {
            case Audio.PlayingState:
                return FA.Icons.faPlayCircle
            case Audio.PausedState:
                return FA.Icons.faPauseCircle
            case Audio.StoppedState:
                return FA.Icons.faStopCircle
            }
        }
        color: "#0772A1"
        anchors.verticalCenter: infoRow.verticalCenter
    }
    Row {
        id: row
        width: 300
        height: 15
        spacing: 1
        anchors.top: parent.top
        anchors.topMargin: 0
        anchors.left: parent.left
        anchors.leftMargin: 0

        Button {
            id: playBtn
            width: 18
            height: 15
            text: {
                if (player.playbackState == Audio.PlayingState)
                    return FA.Icons.faPause
                else
                    return FA.Icons.faPlay
            }
            font.family: FA.Fonts.solid
            font.pixelSize: 10
            onClicked: {
                if (player.playbackState == Audio.PlayingState)
                    player.pause()
                else
                    player.play()
            }
        }

        Button {
            id: nextBtn
            width: 18
            height: 15
            text: FA.Icons.faForward
            font.pixelSize: 10
            font.family: FA.Fonts.solid
            onClicked: {
                player.stop()
                playlist.removeItem(0)
                player.play()
            }
        }

        Button {
            id: settingBtn
            width: 18
            height: 15
            text: FA.Icons.faCog
            font.family: FA.Fonts.solid
            font.pixelSize: 10
            onClicked: {
                settingPop.open()
            }
        }

        Button {
            id: muteBtn
            width: 18
            height: 15
            text: player.volume == 0 ? FA.Icons.faVolumeOff : FA.Icons.faVolumeUp
            font.family: FA.Fonts.solid
            font.pixelSize: 10
            property var lastVol: 0
            onClicked: {
                if (volSlider.value != 0) {
                    lastVol = volSlider.value
                    volSlider.value = 0
                } else {
                    volSlider.value = lastVol
                }
            }
            highlighted: player.volume == 0
        }

        Slider {
            id: volSlider
            visible: muteBtn.hovered || volSlider.hovered
            width: 90
            height: 15
            rightPadding: 0
            leftPadding: 0
            value: 0.5
            handle: Rectangle {
                x: volSlider.leftPadding + volSlider.visualPosition
                   * (volSlider.availableWidth - width)
                y: volSlider.topPadding + volSlider.availableHeight / 2 - height / 2
                implicitWidth: 3
                implicitHeight: 10
                radius: 0
                color: volSlider.pressed ? "#f0f0f0" : "#f6f6f6"
                border.color: "#bdbebf"
            }
        }
    }
}
