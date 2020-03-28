package backend

import (
	"LiveAssistant/bilibili"
	"github.com/go-qamel/qamel"
)

//func init() {
//	RegisterQmlConnectFeedBack("ConnectFeedBack", 1, 0, "ConnectFeedBack")
//	RegisterQmlHandleMsg("HandleMsg", 1, 0, "HandleMsg")
//}

type ConnectFeedBack struct {
	qamel.QmlObject

	_ func(int) int   `slot:"receiveRoomID"`
}

func (m *ConnectFeedBack) receiveRoomID(roomid int) int {
	key, err := GetAccessKey(int32(roomid))
	if err != nil {
		return -1
	}

	// 获取客户端实例
	c, err := bilibili.NewClient(int32(roomid))
	if err != nil || c == nil {
		return -1
	}

	// 启动客户端
	err = c.Start(key)
	if err != nil {
		return -1
	}

	return 0
}

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
}

// 处理各种需要发送到 QML 的消息
//func (h *HandleMsg) init() {
//	go func() {
//		for {
//			select {
//			// 处理用户弹幕
//			case a := <-bilibili.P.DanMu:
//				if e, err := GetDanMu(a); err != nil {
//					s, err := json.Marshal(e)
//					if err != nil {
//						continue
//					}
//					h.sendDanMu(string(s))
//				}
//			// 处理用户礼物
//			case b := <-bilibili.P.Gift:
//				if e := GetGift(b); e != nil {
//					s, err := json.Marshal(e)
//					if err != nil {
//						continue
//					}
//					h.sendGift(string(s))
//				}
//			// 处理贵宾进场，如老爷
//			case c := <-bilibili.P.WelCome:
//				w := GetWelCome(c, 1)
//				s, err := json.Marshal(w)
//				if err != nil {
//					continue
//				}
//				h.sendWelCome(string(s))
//			// 处理房管进场
//			case d := <-bilibili.P.WelComeGuard:
//				w := GetWelCome(d, 2)
//				s, err := json.Marshal(w)
//				if err != nil {
//					continue
//				}
//				h.sendWelComeGuard(string(s))
//			// 处理舰长等贵宾进场
//			case e := <-bilibili.P.GreatSailing:
//				w := GetWelCome(e, 3)
//				s, err := json.Marshal(w)
//				if err != nil {
//					continue
//				}
//				h.sendGreatSailing(string(s))
//			// 处理关注数变动消息
//			case f := <-bilibili.P.Fans:
//				i := json.Get(f, "data", "fans").ToInt()
//				h.sendFansChanged(i)
//			// 处理在线人气变动处理
//			case g := <-bilibili.P.Online:
//				h.sendOnlineChanged(g)
//			}
//		}
//	}()
//}
