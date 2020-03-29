/*
	获取各种参数以及一些工具类型的函数
*/

package backend

import (
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

type UserDanMu struct {
	Avatar string `json:"avatar"`
	Uname  string `json:"uname"`
	Text   string `json:"text"`
}

type UserGift struct {
	Uname  string `json:"uname"`
	Avatar string `json:"avatar"`
	Action string `json:"action"`
	Gname  string `json:"gname"`
	Nums   int32  `json:"nums"`
}

type WelCome struct {
	Uname string `json:"uname"`
	Title string `json:"title"`
}

var
(
	json        = jsoniter.ConfigCompatibleWithStandardLibrary
	keyUrl      = "https://api.live.bilibili.com/room/v1/Danmu/getConf" // params: room_id=xxx&platform=pc&player=web
	userInfoUrl = "https://api.bilibili.com/x/space/acc/info"           //mid=382297465&jsonp=jsonp
)

// 获取发送握手包必须的 key
func GetAccessKey(roomid int32) (key string, err error) {
	url := fmt.Sprintf("%s?room_id=%d&platform=pc&player=web", keyUrl, roomid)

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	rawdata, err := ioutil.ReadAll(resp.Body)

	_ = resp.Body.Close()
	if err != nil {
		return
	}
	key = json.Get(rawdata, "data", "token").ToString()
	return
}

// GetUserAvatar 获取用户头像
func GetUserAvatar(userid int32) (ava string, err error) {
	url := fmt.Sprintf("%s?mid=%d&jsonp=jsonp", userInfoUrl, userid)

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	rawdata, err := ioutil.ReadAll(resp.Body)

	_ = resp.Body.Close()
	if err != nil {
		return
	}
	ava = json.Get(rawdata, "data", "face").ToString()

	return
}

// GetDanMu 提取一条弹幕
func GetDanMu(src []byte) *UserDanMu {
	d := new(UserDanMu)
	u := json.Get(src, "info", 2, 0).ToInt32()
	a, err := GetUserAvatar(u)
	if err != nil {
		return nil
	}
	d.Avatar = a
	d.Uname = json.Get(src, "info", 2, 1).ToString()
	d.Text = json.Get(src, "info", 1).ToString()

	return d
}

// GetGift 获取一条礼物信息
func GetGift(src []byte) *UserGift {
	g := new(UserGift)
	g.Uname = json.Get(src, "data", "uname").ToString()
	g.Avatar = json.Get(src, "data", "face").ToString()
	g.Action = json.Get(src, "data", "action").ToString()
	g.Gname = json.Get(src, "data", "giftName").ToString()
	g.Nums = json.Get(src, "data", "num").ToInt32()

	if g.Nums == 0 || g.Gname == "" || g.Action == "" || g.Avatar == "" || g.Uname == "" {
		return nil
	}
	return g
}

// 1是老爷，2是房管，3是舰长/提督等
func GetWelCome(src []byte, typeID uint8) *WelCome {
	w := new(WelCome)
	switch typeID {
	case 1:
		w.Uname = json.Get(src, "data", "uname").ToString()
		level := json.Get(src, "data", "svip").ToInt()
		if level == 1 {
			w.Title = "年费老爷"
		} else {
			w.Title = "老爷"
		}
	case 2:
		w.Uname = json.Get(src, "data", "username").ToString()
		w.Title = "房管"
	case 3:
		s := json.Get(src, "data", "copy_writing").ToString()
		b := []byte(s)
		// 拿到的字符串形如“欢迎舰长 xxx 进入直播间”
		w.Uname = string(b[6:12])
		w.Title = string(b[14 : len(b)-16])
	}
	if w.Uname == "" || w.Title == "" {
		return nil
	}
	return w
}
