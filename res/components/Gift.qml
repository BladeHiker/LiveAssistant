import QtQuick 2.0
import QtQuick.Controls 2.5
import QtGraphicalEffects 1.0

ItemDelegate {
    id: giftdel
    height: gift.height
    antialiasing: true
    smooth: true
    Rectangle {
        id: gift
        width: 240
        height: 15
        color: "#00000000"

        Text {
            id: txt
            text: '<font color="#0078d7">' + uname + "</font>" + gift_action
                  + '<font color="#0078d7">' + gift_name + "×" + nums + "</font>"
            textFormat: Text.RichText
            font.family: "思源黑体"
            font.bold: true
            verticalAlignment: Text.AlignVCenter
            font.pixelSize: 12
        }
    }
    Glow {
        source: gift
        anchors.fill: gift
        radius: 1
        samples: 20
        spread: 1
        color: "white"
    }
}
