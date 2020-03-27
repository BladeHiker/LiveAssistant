package backend

import (
	"LiveAssistant/bilibili"
	"github.com/go-qamel/qamel"
)

type ConnectFeedBack struct {
	qamel.QmlObject

	_ func(int32)  `slot:"receiveRoomID"`
	_ func(string) `signal:"SendFeedbackMsg"`

}

type HandleMsg struct {
	qamel.QmlObject
	_ func() `constructor:"init"`

	_ func(string) `signal:"SendDanMu"`
	_ func(string) `signal:"SendGift"`
	_ func(string) `signal:"SendWelCome"`
	_ func(string) `signal:"SendWelComeGuard"`
	_ func(string) `signal:"SendGreatSailing"`
	_ func(string) `signal:"SendOnlineChanged"`
	_ func(string) `signal:"SendFansChanged"`
}

func init() {
	RegisterQmlConnectFeedBack("ConnectFeedBack", 1, 0, "ConnectFeedBack")
	RegisterQmlHandleMsg("HandleMsg", 1, 0, "HandleMsg")
}

// 处理各种需要发送到 QML 的消息
func (h *HandleMsg) HandleMsg() {
	go func() {
		for {
			select {
			// 处理用户弹幕
			case a := <-bilibili.P.DanMu:
				if e, err := GetDanMu(a); err != nil {
					s, err := json.Marshal(e)
					if err != nil {
						continue
					}
					h.SendDanMu(string(s))
				}
			// 处理用户礼物
			case b := <-bilibili.P.Gift:
				if e := GetGift(b); e != nil {
					s, err := json.Marshal(e)
					if err != nil {
						continue
					}
					h.SendGift(string(s))
				}
			// 处理贵宾进场，如老爷
			case c := <-bilibili.P.WelCome:
				w := GetWelCome(c, 1)
				s, err := json.Marshal(w)
				if err != nil {
					continue
				}
				h.SendWelCome(string(s))
			// 处理房管进场
			case d := <-bilibili.P.WelComeGuard:
				w := GetWelCome(d, 2)
				s, err := json.Marshal(w)
				if err != nil {
					continue
				}
				h.SendWelComeGuard(string(s))
			// 处理舰长等贵宾进场
			case e := <-bilibili.P.GreatSailing:
				w := GetWelCome(e, 3)
				s, err := json.Marshal(w)
				if err != nil {
					continue
				}
				h.SendGreatSailing(string(s))
			// 处理关注数变动消息
			case f := <-bilibili.P.Fans:
				i := json.Get(f,"data","fans").ToInt()
				h.SendFansChanged(i)
			// 处理在线人气变动处理
			case g := <-bilibili.P.Online:
				h.SendOnlineChanged(g)
			}
		}
	}()
}

func (m *ConnectFeedBack) receiveRoomID(roomid int32) {
	key, err := GetAccessKey(roomid)
	if err != nil {
		m.SendFeedbackMsg("房间号输入有误")
		return
	}

	// 获取客户端实例
	c, err := bilibili.NewClient(roomid)
	if err != nil || c == nil {
		m.SendFeedbackMsg("获取客户端实例失败")
		return
	}

	// 启动客户端
	err = c.Start(key)
	if err != nil {
		m.SendFeedbackMsg("启动客户端失败")
	}
}
