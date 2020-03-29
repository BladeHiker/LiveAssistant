package bilibili

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

var roomid int32 = 21773215
var key string

func TestNewClient(t *testing.T) {
	c, err := NewClient(roomid)
	if err != nil || c == nil {
		t.Error("NewClient(roomid) err")
		return
	}
	fmt.Println("NewClient后c是",*c)

	key, err = GetAccessKey(roomid)
	if err != nil {
		t.Error("TestNewClient() -> GetAccessKey(roomid) err")
		return
	}
	fmt.Println("key是",key)

	err = c.StartForTest(key)
	if err != nil {
		t.Error("c.Start(key) err")
		return
	}
	// 测试 30 秒
	time.Sleep(time.Second * 30)
}

// 重写是重新调用新写的c.ReceiveMsgForTest()方法
func (c *Client) StartForTest(key string) (err error) {
	m := NewHandShakeMsg(c.RoomID)
	m.Key = key
	fmt.Println("HandShake Msg是",*m)

	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = c.SendPackage(0, 16, 1, 7, 1, b)
	if err != nil {
		return
	}

	go c.ReceiveMsgForTest()
	go c.HeartBeat()
	go handle()

	return
}

// 防止循环引用，直接把另一个包的工具函数拷贝到这里使用
var keyUrl = "https://api.live.bilibili.com/room/v1/Danmu/getConf" // params: room_id=xxx&platform=pc&player=web
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

// 重写是不需要写入管道，直接输出就行
func (c *Client) ReceiveMsgForTest() {
	P = NewPool()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			continue
		}

		// 根据消息类型进行分类处理
		switch msg[11] {
		// 服务器发来的心跳包下行，实体部分仅直播间人气值
		case 3:
			h := ByteArrToDecimal(msg[16:])
			if int32(h) != c.Online {
				c.Online = int32(h)
				P.Online <- h
			}
		case 5:
			inflated, err := ZlibInflate(msg[16:])
			if err == nil {
				// 代表数据需要压缩，如DANMU_MSG，SEND_GIFT等信息量较大的数据包
				for len(inflated) > 0 {
					l := ByteArrToDecimal(inflated[:4])
					c := json.Get(inflated[16:l], "cmd").ToString()
					switch CMD(c) {
					case CMDDanmuMsg:
						P.DanMu <- inflated[16:l]
						//fmt.Println(string(inflated[16:l]))
					case CMDSendGift:
						P.Gift <- inflated[16:l]
						//fmt.Println(string(inflated[16:l]))
					case CMDWELCOME:
						P.WelCome <- inflated[16:l]
						//fmt.Println(string(inflated[16:l]))
					case CMDWelcomeGuard:
						P.WelComeGuard <- inflated[16:l]
						//fmt.Println(string(inflated[16:l]))
					case CMDEntry:
						P.GreatSailing <- inflated[16:l]
						//fmt.Println(string(inflated[16:l]))
					case CMDRoomRealTimeMessageUpdate:
						P.Fans <- inflated[16:l]
						//fmt.Println(string(inflated[16:l]))
					}
					inflated = inflated[l:]
				}
			}
		}
	}
}

// 处理各种需要发送到 QML 的消息
func handle() {
	go func() {
		for {
			select {
			// 处理用户弹幕
			case a := <-P.DanMu:
				fmt.Println("这是一条弹幕，正在提取出信息 准备处理。。。",a)
				if e := GetDanMu(a); e != nil {
					s, err := json.Marshal(e)
					if err != nil {
						continue
					}
					//h.sendDanMu(string(s))
					fmt.Println("DanMu: ", string(s))
				}
			// 处理用户礼物
			case b := <-P.Gift:
				fmt.Println("这是一条礼物，正在提取出信息 准备处理。。。")
				if e := GetGift(b); e != nil {
					s, err := json.Marshal(e)
					if err != nil {
						continue
					}
					//h.sendGift(string(s))
					fmt.Println("Gift: ", string(s))
				}
			// 处理贵宾进场，如老爷
			case c := <-P.WelCome:
				fmt.Println("这是一条老爷进场，正在提取出信息 准备处理。。。")
				w := GetWelCome(c, 1)
				s, err := json.Marshal(w)
				if err != nil {
					continue
				}
				//h.sendWelCome(string(s))
				fmt.Println("WelCome: ", string(s))
			// 处理房管进场
			case d := <-P.WelComeGuard:
				fmt.Println("这是一条房管进场，正在提取出信息 准备处理。。。")
				w := GetWelCome(d, 2)
				s, err := json.Marshal(w)
				if err != nil {
					continue
				}
				fmt.Println("WelComeGuard: ", string(s))
				//h.sendWelComeGuard(string(s))
			// 处理舰长等贵宾进场
			case e := <-P.GreatSailing:
				fmt.Println("这是一条大航海进场，正在提取出信息 准备处理。。。")
				w := GetWelCome(e, 3)
				s, err := json.Marshal(w)
				if err != nil {
					continue
				}
				fmt.Println("GreatSailing: ", string(s))
				//h.sendGreatSailing(string(s))
			// 处理关注数变动消息
			case f := <-P.Fans:
				fmt.Println("这是一条粉丝更新，正在提取出信息 准备处理。。。")
				i := json.Get(f, "data", "fans").ToInt()
				fmt.Println("Fans: ", i)
				//h.sendFansChanged(i)
			// 处理在线人气变动处理
			case g := <-P.Online:
				fmt.Println("这是一条人气更新，正在提取出信息 准备处理。。。")
				//h.sendOnlineChanged(g)
				fmt.Println("Online: ", g)
			}
		}
	}()
}

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
	userInfoUrl = "https://api.bilibili.com/x/space/acc/info" //mid=382297465&jsonp=jsonp
)

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
