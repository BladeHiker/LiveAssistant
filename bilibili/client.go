package bilibili

import (
	"bytes"
	"encoding/binary"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	"net/url"
	"time"
)

// 客户端实例
type Client struct {
	RoomID      int32           // 房间 ID
	Online      int32           // 用来判断人气是否变动
	Conn        *websocket.Conn // 连接后的对象
	IsConnected bool
}

// HandShakeMsg 定义了握手包的信息格式
type HandShakeMsg struct {
	Uid       int32  `json:"uid"`
	RoomID    int32  `json:"roomid"`
	Protover  int32  `json:"protover"`
	Platform  string `json:"platform"`
	Clientver string `json:"clientver"`
	Type      int32  `json:"type"`
	Key       string `json:"key"`
}

// 返回一个初始化了的握手包信息实例
func NewHandShakeMsg(roomid int32) *HandShakeMsg {
	return &HandShakeMsg{
		Uid:       0,
		RoomID:    roomid,
		Protover:  2,
		Platform:  "web",
		Clientver: "1.10.3",
		Type:      2,
		Key:       "",
	}
}

type CMD string

var (
	RealID      = "http://api.live.bilibili.com/room/v1/Room/room_init" // params: id=xxx
	DanMuServer = "ks-live-dmcmt-bj6-pm-02.chat.bilibili.com:443"
	json        = jsoniter.ConfigCompatibleWithStandardLibrary
	P           *Pool

	CMDDanmuMsg                  CMD = "DANMU_MSG"                     // 普通弹幕信息
	CMDSendGift                  CMD = "SEND_GIFT"                     // 普通的礼物，不包含礼物连击
	CMDWELCOME                   CMD = "WELCOME"                       // 欢迎VIP
	CMDWelcomeGuard              CMD = "WELCOME_GUARD"                 // 欢迎房管
	CMDEntry                     CMD = "ENTRY_EFFECT"                  // 欢迎舰长等头衔
	CMDRoomRealTimeMessageUpdate CMD = "ROOM_REAL_TIME_MESSAGE_UPDATE" // 房间关注数变动
)

// 获取一个连接好的客户端实例
func NewClient(roomid int32) (c *Client, err error) {
	c = new(Client)

	realid, err := GetRealRoomID(roomid)
	if err != nil {
		return nil, err
	}

	// 连接弹幕服务器并发送握手包
	u := url.URL{Scheme: "wss", Host: DanMuServer, Path: "sub"}
	c.Conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}

	c.IsConnected = true
	c.RoomID = int32(realid)
	return
}

// 发送握手包并开始监听消息
func (c *Client) Start(key string) (err error) {
	m := NewHandShakeMsg(c.RoomID)
	m.Key = key
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = c.SendPackage(0, 16, 1, 7, 1, b)

	go c.ReceiveMsg()
	go c.HeartBeat()
	return
}

func (c *Client) SendPackage(packetlen uint32, magic uint16, ver uint16, typeID uint32, param uint32, data []byte) (err error) {
	packetHead := new(bytes.Buffer)

	if packetlen == 0 {
		packetlen = uint32(len(data) + 16)
	}
	var pdata = []interface{}{
		packetlen,
		magic,
		ver,
		typeID,
		param,
	}

	// 将包的头部信息以大端序方式写入字节数组
	for _, v := range pdata {
		if err = binary.Write(packetHead, binary.BigEndian, v); err != nil {
			return
		}
	}

	// 将包内数据部分追加到数据包内
	sendData := append(packetHead.Bytes(), data...)

	if err = c.Conn.WriteMessage(websocket.BinaryMessage, sendData); err != nil {
		return
	}

	return
}

func (c *Client) ReceiveMsg() {
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
				// 代表数据需要压缩，如DANMU_MSG，SEND_GIFT等信息量较少的数据包
				for len(inflated) > 0 {
					l := ByteArrToDecimal(inflated[:4])
					c := json.Get(inflated[16:l], "cmd").ToString()
					switch CMD(c) {
					case CMDDanmuMsg:
						P.DanMu <- inflated[16:l]
					case CMDSendGift:
						P.Gift <- inflated[16:l]
					case CMDWELCOME:
						P.WelCome <- inflated[16:l]
					case CMDWelcomeGuard:
						P.WelComeGuard <- inflated[16:l]
					case CMDEntry:
						P.GreatSailing <- inflated[16:l]
					case CMDRoomRealTimeMessageUpdate:
						P.Fans <- inflated[16:l]
					}
					inflated = inflated[l:]
				}
			}
		}
	}
}

func (c *Client) HeartBeat() {
	for {
		if c.IsConnected {
			obj := []byte("5b6f626a656374204f626a6563745d")
			if err := c.SendPackage(0, 16, 1, 2, 1, obj); err != nil {
				continue
			}
			time.Sleep(30 * time.Second)
		}
	}
}
