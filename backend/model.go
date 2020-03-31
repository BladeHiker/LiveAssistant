/*
	获取各种参数以及一些工具类型的函数
*/

package backend

import (
	"fmt"
	"github.com/json-iterator/go"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var
(
	json        = jsoniter.ConfigCompatibleWithStandardLibrary
	keyUrl      = "https://api.live.bilibili.com/room/v1/Danmu/getConf" // params: room_id=xxx&platform=pc&player=web
	userInfoUrl = "https://api.bilibili.com/x/space/acc/info"           //mid=382297465&jsonp=jsonp
	server      = "shiluo.design:3000"
	RoomInfoURI = "https://api.live.bilibili.com/xlive/web-room/v1/index/getInfoByRoom" // params:?room_id=923833
	MusicInfo   chan string
)

const (
	// TODO 舰长身份的识别
	CommonUser = 1      	// 普通用户
	Vip        = 1 << 1 	// 老爷
	Guard      = 1 << 2 	// 房管
	Sailing    = 1 << 3 	// 大航海
)

type UserDanMu struct {
	Avatar string `json:"avatar"`
	// 用户头衔
	Utitle int `json:"utitle"`
	// 用户等级
	UserLevel int `json:"user_level"`
	// 用户牌子
	MedalName string `json:"medal_name"`
	// 牌子等级
	MedalLevel int    `json:"medal_level"`
	Uname      string `json:"uname"`
	Text       string `json:"text"`
}

type UserGift struct {
	Uname  string `json:"uname"`
	Avatar string `json:"avatar"`
	Action string `json:"action"`
	Gname  string `json:"gname"`
	Nums   int32  `json:"nums"`
	Price  int    `json:"price"`
}

type WelCome struct {
	Uname string `json:"uname"`
	Title string `json:"title"`
}

type LocalInfo struct {
	MemUsedPercent float64   `json:"mem"`        // 内存使用率
	CpuUsedPercent float64   `json:"cpu"`        // CPU使用率
	SendBytes      int64     `json:"send"`       // 单位时间发送字节数
	RecvBytes      int64     `json:"recv"`       // 单位时间接收字节数
	DiskUsed       []float64 `json:"disk_used"`  // 磁盘使用率
	DiskRead       []int64   `json:"disk_read"`  // 磁盘读取字节数
	DiskWrite      []int64   `json:"disk_write"` // 磁盘写入字节数
}

// 获取发送握手包必须的 key
func GetAccessKey(roomid int32) (key string, err error) {
	u := fmt.Sprintf("%s?room_id=%d&platform=pc&player=web", keyUrl, roomid)

	resp, err := http.Get(u)
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
	u := fmt.Sprintf("%s?mid=%d&jsonp=jsonp", userInfoUrl, userid)

	resp, err := http.Get(u)
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
	d.MedalName = json.Get(src, "info", 3, 1).ToString()
	d.MedalLevel = json.Get(src, "info", 3, 0).ToInt()
	d.UserLevel = json.Get(src, "info", 4, 0).ToInt()

	// 判定用户称呼，比如 房管 | 老爷 | 舰长等等，用二进制位按位与表示

	guard := json.Get(src, "info", 2, 2).ToInt()
	vip := json.Get(src, "info", 2, 3).ToInt()
	d.Utitle = guard | vip | CommonUser

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
	g.Price = json.Get(src, "data", "price").ToInt()

	if g.Price == 0 {
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
		return nil, s
	}
	if w.Uname == "" || w.Title == "" {
		return nil, ""
	} else {
		return w, ""
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

// 根据官方的api获取关注数
func GetFansByAPI(roomid int) int {
	u := fmt.Sprintf("%s?room_id=%d", RoomInfoURI, roomid)

	resp, err := http.Get(u)
	if err != nil {
		return 0
	}

	rawdata, err := ioutil.ReadAll(resp.Body)

	_ = resp.Body.Close()
	if err != nil {
		return 0
	}

	fans := json.Get(rawdata, "data", "anchor_info", "relation_info", "attention").ToInt()
	return fans
}

func GetCompInfo() (l *LocalInfo) {
	l = new(LocalInfo)
	vm, _ := mem.VirtualMemory()
	f, _ := cpu.Percent(time.Second, false)
	io, _ := net.IOCounters(true)
	for _, v := range io {
		// qamel 不支持uint64类型，转换一下
		l.SendBytes += int64(v.BytesSent)
		l.RecvBytes += int64(v.BytesRecv)
	}

	// 不判错，若获取失败返回零值
	l.MemUsedPercent = vm.UsedPercent
	l.CpuUsedPercent = f[0]

	//TODO 磁盘使用率，读写量暂定

	return
}
