package backend

import (
	"fmt"
	"testing"
)

var roomid int32 = 21773215

func TestGetAccessKey(t *testing.T) {
	key, err := GetAccessKey(roomid)
	if err != nil || key == "" {
		t.Error("`GetAccessKey(757349)` key error")
	}
	fmt.Println(key)
}

var userid int32 = 265975162

func TestGetUserAvatar(t *testing.T) {
	key, err := GetUserAvatar(userid)
	if err != nil || key == "" {
		t.Error("`GetUserAvatar(265975162)` key error")
	}
}

var DanMu = `
{
	"cmd": "DANMU_MSG",
	"info": [
		[0, 1, 25, 16777215, 1585411648979, 1585440447, 0, "c6cb1863", 0, 0, 0], "翻车了？", [73395517, "山南玖玖玖", 0, 0, 0, 10000, 1, ""],
		[],
		[11, 0, 6406234, "\u003e50000"],
		["title-111-1", "title-111-1"], 0, 0, null, {
			"ts": 1585411648,
			"ct": "D432158E"
		},
		0, 0, null, null, 0
	]
}`

func TestGetDanMu(t *testing.T) {
	d := GetDanMu([]byte(DanMu))
	if d == nil {
		t.Error("GetDanMu([]byte(DanMu)) err")
	}
}

var Gift = `
{
	"cmd": "SEND_GIFT",
	"data": {
		"giftName": "\u8fa3\u6761",
		"num": 3,
		"uname": "\u56de\u5fc6_____Memento",
		"face": "http:\/\/i0.hdslb.com\/bfs\/face\/ba36802882c41ebe19aef77523cb36f26ee1934c.jpg",
		"guard_level": 0,
		"rcost": 23340190,
		"uid": 115591474,
		"top_list": [],
		"timestamp": 1585412026,
		"giftId": 1,
		"giftType": 0,
		"action": "\u5582\u98df",
		"super": 0,
		"super_gift_num": 0,
		"super_batch_gift_num": 0,
		"batch_combo_id": "",
		"price": 100,
		"rnd": "1388506129",
		"newMedal": 0,
		"newTitle": 0,
		"medal": [],
		"title": "",
		"beatId": "",
		"biz_source": "live",
		"metadata": "",
		"remain": 0,
		"gold": 0,
		"silver": 0,
		"eventScore": 0,
		"eventNum": 0,
		"smalltv_msg": [],
		"specialGift": null,
		"notice_msg": [],
		"smallTVCountFlag": true,
		"capsule": null,
		"addFollow": 0,
		"effect_block": 1,
		"coin_type": "silver",
		"total_coin": 300,
		"effect": 0,
		"broadcast_id": 0,
		"draw": 0,
		"crit_prob": 0,
		"tag_image": "",
		"send_master": null,
		"is_first": true,
		"demarcation": 1,
		"combo_stay_time": 2,
		"combo_total_coin": 0
	}
}
`

func TestGetGift(t *testing.T) {
	g := GetGift([]byte(Gift))
	if g == nil {
		t.Error("GetGift([]byte(Gift)) err")
	}
}

func TestGetMusicURI(t *testing.T) {
	singer, mname := "beyond", "海阔天空"
	_, err := GetMusicURI(singer, mname)
	if err != nil {
		t.Error("GetMusicURI(singer,mname) err:",err)
	}
}
