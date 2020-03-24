package backend

import (
	"LiveAssistant/bilibili"
	"github.com/go-qamel/qamel"
)

type Connect struct {
	qamel.QmlObject

	_ func(int32) `slot:"RecRoomID"`

	_ func(string) `signal:"SendFeedbackMsg"`

	_ func(string) `signal:"SendGift"`
	_ func(string) `signal:"SendWelCome"`
	_ func(int32)  `signal:"SendOnlineChanged"`
}

func init() {
	RegisterQmlConnect("Connect", 1, 0, "Connect")
	RegisterQmlDanMu("DanMu", 1, 0, "DanMu")
}

var p *bilibili.Pool

type DanMu struct {
	qamel.QmlObject
	_ func()       `constructor:"init"`
	_ func(string) `signal:"SendDanMu"`
}

func (d *DanMu) init() {
	go func() {
		for {
			b := <-p.DanMu
			if e, err := GetDanMu(b); err != nil {
				s, err := json.Marshal(e)
				if err != nil {
					continue
				}
				d.SendDanMu(string(s))
			}
		}
	}()
}

func (m *Connect) RecRoomID(roomid int32) {
	p = bilibili.NewPool()

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
	err = c.Start()
	if err != nil {
		m.SendFeedbackMsg("启动客户端失败")
	}
}
