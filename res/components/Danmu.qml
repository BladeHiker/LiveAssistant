import QtQuick 2.12
import QtQuick.Controls 2.5
import QtGraphicalEffects 1.0
import "../fonts/FontAwesome" as FA
import "../theme" as T

ItemDelegate {
    id: danmu0
    property var colorlist: ["#67EB6D", "#FFB300", "#FF5F00", "#7C07A9", "#DC0055", "#530FAD", "#0772A1", "#FFE900", "#CCF600", "#CD0074"]
    antialiasing: true
    smooth: true
    height: txt.height + 2
    Row {
        id: row
        width: parent.width
        anchors.verticalCenter: parent.verticalCenter
        spacing: 3
        Rectangle {
            id: rectangle
            visible: type == 1
            width: txt.font.pixelSize + 3
            height: txt.font.pixelSize + 3
            anchors.top: parent.top
            anchors.topMargin: 0
            border.color: T.ColorDesign.avatarBorder
            clip: true
            color: T.ColorDesign.avatarBG
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
                // color: "#000000"
                color: T.ColorDesign.danmuTextMain
                width: 268

                text: {
                    switch (type) {
                    case 0:
                        return '<center>' + msg + '</center>'
                    case 1:
                        return '<font color="' + T.ColorDesign.danmuTextUname
                                + '">' + uname + ":  </font>" + msg
                    case 2:
                        return "欢迎" + '<font color="' + T.ColorDesign.danmuTextUtitle
                                + '">' + title + '</font>' + '<font color="'
                                + T.ColorDesign.danmuTextUname + '">' + uname + '</font>' + "进入直播间"
                    }
                }
                anchors.verticalCenter: parent.verticalCenter
                anchors.left: parent.left
                anchors.leftMargin: 0
                wrapMode: Text.WrapAnywhere
                textFormat: Text.StyledText
                //                font.family: "思源黑体"
                font.family: "黑体"
                lineHeightMode: Text.ProportionalHeight
                lineHeight: 1.2
                font.bold: true
                verticalAlignment: Text.AlignVCenter
                font.pixelSize: {
                    if (type == 0)
                        return 12

                    return 14
                }
            }
            // Glow {
            //     source: txt
            //     anchors.fill: txt
            //     radius: 3
            //     samples: 20
            //     spread: 1
            //     color: "white"
            //     antialiasing: true
            // }
        }
    }
}
