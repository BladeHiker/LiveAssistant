/*
	获取各种参数以及一些工具类型的函数
*/

package backend

import (
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"net/url"
)

var
(
	json        = jsoniter.ConfigCompatibleWithStandardLibrary
	keyUrl      = "https://api.live.bilibili.com/room/v1/Danmu/getConf" // params: room_id=xxx&platform=pc&player=web
	userInfoUrl = "https://api.bilibili.com/x/space/acc/info"           //mid=382297465&jsonp=jsonp
	server      = "shiluo.design:3000"
	MusicInfo   chan string
)

type UserDanMu struct {
	Avatar string `json:"avatar"`
	// 老爷0，房管1，舰长2，提督3，总督4，普通5
	Utitle int`json:"utitle"`
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
func GetWelCome(src []byte, typeID uint8) (*WelCome, string) {
	w := new(WelCome)
	var s string
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
		s = json.Get(src, "data", "copy_writing").ToString()
		return nil,s
	}
	if w.Uname == "" || w.Title == "" {
		return nil, ""
	} else {
		return w,""
	}
}

// 根据歌手名和歌曲获取音乐URI地址
func GetMusicURI(singer, mname string) (URI string, err error) {
	// 根据歌手名，音乐名获取歌曲id
	q := url.Values{}
	q.Set("keywords", singer+" "+mname)
	q.Set("limit", "1")
	u := url.URL{
		Scheme:   "http",
		Host:     server,
		Path:     "search",
		RawQuery: q.Encode(),
	}
	resp, err := http.Get(u.String())
	if err != nil {
		return
	}

	rawdata, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return
	}
	id := json.Get(rawdata, "result", "songs", 0, "id").ToInt()

	fmt.Println("歌曲id是", id)

	// 根据id获取歌曲uri
	r := fmt.Sprintf("http://%s/song/url?id=%d", server, id)
	res, err := http.Get(r)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(res.Body)

	_ = resp.Body.Close()
	if err != nil {
		return
	}

	URI = json.Get(data, "data", 0, "url").ToString()

	fmt.Println("URI是", URI)

	return
}
