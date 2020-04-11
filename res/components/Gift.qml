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
            text: '<font color="#FFE87A">' + uname + "</font>" + gift_action
                  + '<font color="#64E8FF">' + gift_name + "×" + nums + "</font>"
            textFormat: Text.RichText
            font.family: "黑体"
            font.bold: true
            verticalAlignment: Text.AlignVCenter
            font.pixelSize: 12
            anchors.left: gift.left
            color:"#ffffff"
        }
        //        Text {
        //            id: priceEq
        //            text: "=" + price
        //            font.family: "黑体"
        //            font.bold: true
        //            verticalAlignment: Text.AlignVCenter
        //            font.pixelSize: 12
        //            anchors.right: gift.right
        //        }
    }
    // Glow {
    //     source: gift
    //     anchors.fill: gift
    //     radius: 1
    //     samples: 20
    //     spread: 1
    //     color: "white"
    // }
}
