package backend

import (
	"LiveAssistant/bilibili"
	_ "LiveAssistant/bilibili"
	"github.com/go-qamel/qamel"
	"strings"
	"time"
)

func init() {
	RegisterQmlConnectFeedBack("ConnectFeedBack", 1, 0, "ConnectFeedBack")
	RegisterQmlHandleMsg("HandleMsg", 1, 0, "HandleMsg")
	RegisterQmlHandleMsg("Music", 1, 0, "Music")
}

// 连接直播间模块定义
type ConnectFeedBack struct {
	qamel.QmlObject

	_ func()        `constructor:"init"`
	_ func(int) `slot:"receiveRoomID"`
	_ func(int)     `signal:"sendFansNums"`
	_ func(string)  `signal:"sendCompInfo"`
	_ func(int) `signal:"sendErr"`
}

func (m *ConnectFeedBack) init() {
	go func() {
		for {
			i := GetCompInfo()
			b, err := json.Marshal(i)
			if err != nil {
				continue
			}
			m.sendCompInfo(string(b))
			time.Sleep(5 * time.Second)
		}
	}()
}

func (m *ConnectFeedBack) receiveRoomID(roomid int) {
	ConnectAndServe(roomid)
	if bilibili.UserClient.IsConnected == false {
		m.sendErr(-1)
	}
	m.sendFansNums(GetFansByAPI(roomid))

	// 监听客户端连接状态，如果连接中断，则重新建立客户端连接
	go func() {
		for {
			if bilibili.UserClient.IsConnected == false {
				ConnectAndServe(roomid)
				time.Sleep(time.Second * 3)
			}
		}
	}()
}

// 消息处理模块定义
type HandleMsg struct {
	qamel.QmlObject
	_ func() `constructor:"init"`

	_ func(string) `signal:"sendDanMu"`
	_ func(string) `signal:"sendGift"`
	_ func(string) `signal:"sendWelCome"`
	_ func(string) `signal:"sendWelComeGuard"`
	_ func(string) `signal:"sendGreatSailing"`
	_ func(int)    `signal:"sendOnlineChanged"`
	_ func(int)    `signal:"sendFansChanged"`

	_      func(bool, string) `slot:"musicControl"`
	Button bool               // 点歌功能的开关
	Key    string             // 点歌关键字
}

// 处理各种需要发送到 QML 的消息
func (h *HandleMsg) init() {
	go func() {
		for {
			select {
			// 处理用户弹幕
			case a := <-bilibili.P.DanMu:
				if e := GetDanMu(a); e != nil && h.Button == false {
					s, err := json.Marshal(e)
					if err != nil {
						continue
					}
					h.sendDanMu(string(s))
				} else if e != nil && h.Button == true {
					if strings.HasPrefix(e.Text, h.Key) {
						MusicInfo <- e.Text
					}
				}
			// 处理用户礼物
			case b := <-bilibili.P.Gift:
				if e := GetGift(b); e != nil {
					s, err := json.Marshal(e)
					if err != nil {
						continue
					}
					h.sendGift(string(s))
				}
			// 处理贵宾进场，如老爷
			case c := <-bilibili.P.WelCome:
				if w := GetWelCome(c, 1); w != nil {
					s, err := json.Marshal(w)
					if err != nil {
						continue
					}
					h.sendWelCome(string(s))
				}
			// 处理房管进场
			case d := <-bilibili.P.WelComeGuard:
				if w := GetWelCome(d, 2); w != nil {
					s, err := json.Marshal(w)
					if err != nil {
						continue
					}
					h.sendWelComeGuard(string(s))
				}
			// 处理舰长等贵宾进场
			case e := <-bilibili.P.GreatSailing:
				if w := GetWelCome(e, 3); w != nil {
					s, err := json.Marshal(w)
					if err != nil {
						continue
					}
					h.sendGreatSailing(string(s))
				}
			// 处理关注数变动消息
			case f := <-bilibili.P.Fans:
				i := json.Get(f, "data", "fans").ToInt()
				h.sendFansChanged(i)
			// 处理在线人气变动处理
			case g := <-bilibili.P.Online:
				h.sendOnlineChanged(g)
			}
		}
	}()
}

func (h *HandleMsg) musicControl(b bool, key string) {
	go func() {
		for {
			// 代表打开点歌功能
			if b == true && key != "" {
				h.Button = true
				h.Key = key
			} else if b == false {
				h.Button = false
				h.Key = ""
			} else {
				continue
			}
		}
	}()
}

// 音乐模块定义
type Music struct {
	qamel.QmlObject

	_ func()                       `constructor:"init"`
	_ func(string, string, string) `signal:"sendMusicURI"`
}

func (m *Music) init() {
	MusicInfo := make(chan string, 10)
	go func() {
		for {
			select {
			case text := <-MusicInfo:
				s := strings.Split(text, " ")
				uri, err := GetMusicURI(s[1], s[2])
				if err != nil || uri == "" {
					continue
				}
				m.sendMusicURI(uri, s[1], s[2])
			}
		}
	}()
}
