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

var
(
	keyUrl      = "https://api.live.bilibili.com/room/v1/Danmu/getConf" // params: room_id=xxx&platform=pc&player=web
	json        = jsoniter.ConfigCompatibleWithStandardLibrary
	userInfoUrl = "https://api.bilibili.com/x/space/acc/info" //mid=382297465&jsonp=jsonp
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
	key = json.Get(rawdata, "data").Get("token").ToString()
	return
}

// 提取一条弹幕
func GetDanMu(src []byte) (*UserDanMu, error) {
	d := new(UserDanMu)
	u := json.Get(src, "info", 2, 0).ToUint32()
	a, err := GetUserAvatar(u)
	if err != nil {
		return nil, err
	}
	d.Avatar = a
	d.Uname = json.Get(src, "info", 2, 1).ToString()
	d.Text = json.Get(src, "info", 1).ToString()

	return d, nil
}

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
