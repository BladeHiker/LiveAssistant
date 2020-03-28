package bilibili

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

var roomid int32 = 813364
var key string

func TestNewClient(t *testing.T) {
	c, err := NewClient(roomid)
	if err != nil || c == nil {
		t.Error("NewClient(roomid) err")
		return
	}

	key, err = GetAccessKey(roomid)
	if err != nil {
		t.Error("TestNewClient() -> GetAccessKey(roomid) err")
		return
	}

	err = c.StartForTest(key)
	if err != nil {
		t.Error("c.Start(key) err")
		return
	}
	// 测试 30 秒
	time.Sleep(time.Second*30)
}

// 重写是重新调用新写的c.ReceiveMsgForTest()方法
func (c *Client) StartForTest(key string) (err error) {
	m := NewHandShakeMsg(c.RoomID)
	m.Key = key

	//fmt.Println(key)

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
	// P = NewPool()

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
				// P.Online <- h
				fmt.Println(int32(h))
			}
		case 5:
			inflated, err := ZlibInflate(msg[16:])
			if err == nil {
				// 代表数据需要压缩，如DANMU_MSG，SEND_GIFT等信息量较少的数据包
				for len(inflated) > 0 {
					l := ByteArrToDecimal(inflated[:4])
					c := json.Get(inflated[16:l], "cmd").ToString()
					switch CMD(c) {
					case CMDDanmuMsg:
						// P.DanMu <- inflated[16:l]
						fmt.Println(string(inflated[16:l]))
					case CMDSendGift:
						// P.Gift <- inflated[16:l]
						fmt.Println(string(inflated[16:l]))
					case CMDWELCOME:
						// P.WelCome <- inflated[16:l]
						fmt.Println(string(inflated[16:l]))
					case CMDWelcomeGuard:
						// P.WelComeGuard <- inflated[16:l]
						fmt.Println(string(inflated[16:l]))
					case CMDEntry:
						// P.GreatSailing <- inflated[16:l]
						fmt.Println(string(inflated[16:l]))
					case CMDRoomRealTimeMessageUpdate:
						// P.Fans <- inflated[16:l]
						fmt.Println(string(inflated[16:l]))
					}
					inflated = inflated[l:]
				}
			}
		}
	}
}
