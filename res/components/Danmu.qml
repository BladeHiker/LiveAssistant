import QtQuick 2.12
import QtQuick.Controls 2.5
import QtGraphicalEffects 1.0
import "../fonts/FontAwesome" as FA

ItemDelegate {
    id: danmu0
    property var colorlist: ["#00A08A", "#FFB300", "#FF5F00", "#7C07A9", "#DC0055", "#530FAD", "#0772A1", "#FFE900", "#CCF600", "#CD0074"]
    antialiasing: true
    smooth: true
    height: txt.height
    Row {
        id: row
        width: parent.width
        spacing: 3
        Rectangle {
            id: rectangle
            visible: type == 1
            width: txt.font.pixelSize + 3
            height: txt.font.pixelSize + 3
            anchors.top: parent.top
            anchors.topMargin: 3
            border.color: colorlist[0]
            clip: true
            color: "#eee"
            BorderImage {
                id: icon
                anchors.fill: parent
                layer.smooth: true
                antialiasing: true
                source: avatar
            }
        }
        Rectangle {
            id: txtbox
            height: txt.height
            width: 270
            color: "#00000000"
            Text {
                id: txt
                color: "#000000"
                width: 268
                text: {
                    switch (type) {
                    case 0:
                        return '<center>' + msg + '</center>'
                    case 1:
                        return '<font color="' + colorlist[0] + '">' + uname + ":  </font>" + msg
                    case 2:
                        return "欢迎 " + '<font color="' + colorlist[2] + '">  '
                                + title + '</font>' + '<font color="' + colorlist[0]
                                + '">' + " " + uname + '</font>' + "进入直播间"
                    }
                }

                anchors.verticalCenter: parent.verticalCenter
                anchors.left: parent.left
                anchors.leftMargin: 0
                wrapMode: Text.WrapAnywhere
                textFormat: Text.StyledText
                font.family: "思源黑体"
                font.bold: true
                font.pixelSize: {
                    if (type == 0)
                        return 11

                    return 13
                }
            }
            Glow {
                source: txt
                anchors.fill: txt
                radius: 3
                samples: 20
                spread: 1
                color: "white"
                antialiasing: true
            }
        }
    }
}
