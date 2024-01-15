package websocat

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jaronnie/ragingcd/core/server/service/terminal/message"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 3000 * time.Second
	maxMessageSize = 8192
	pongWait       = 6000 * time.Second
)

type Connection interface {
	ReadMsg() (*message.TerminalMessage, error)
	WriteMsg(msg *message.TerminalMessage) error
	Close() error
	CloseGracingWithinTime(gracingDuration time.Duration) error
}

var upgrader = func() websocket.Upgrader {
	upgrader := websocket.Upgrader{}
	upgrader.HandshakeTimeout = time.Second * 2
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	return upgrader
}()

type WsConn struct {
	conn *websocket.Conn
}

func NewConn(w http.ResponseWriter, r *http.Request) (Connection, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	return &WsConn{conn: conn}, nil
}

func (c *WsConn) ReadMsg() (*message.TerminalMessage, error) {
	err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		return nil, err
	}
	c.conn.SetReadLimit(maxMessageSize)
	frameType, frameData, err := c.conn.ReadMessage()
	if err != nil {
		if !isExpectedError(err) {
			fmt.Printf("Fail to read message. Err: [%v]", err)
		}
		return nil, err
	}
	if frameType == websocket.TextMessage {
		return message.Stdin(frameData)
	}
	return nil, fmt.Errorf("fail to support frame type %d", frameType)
}

func (c *WsConn) WriteMsg(msg *message.TerminalMessage) error {
	err := c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	if err != nil {
		return err
	}
	if msg.Data == "\u0000" || msg.Data == "?" {
		return nil
	}
	err = c.conn.WriteMessage(websocket.TextMessage, msg.Bytes())
	if err != nil {
		if !isExpectedError(err) && err != websocket.ErrCloseSent {
			fmt.Printf("Fail to write msg to websocket connection. Err: [%v]", err)
		}
	}
	return err
}

func isExpectedError(err error) bool {
	expectedError := []int{
		websocket.CloseGoingAway,
		websocket.CloseNormalClosure,
		websocket.CloseNoStatusReceived,
		websocket.CloseAbnormalClosure,
	}
	return websocket.IsCloseError(err, expectedError...)
}

func (c *WsConn) Close() error {
	return c.CloseGracingWithinTime(time.Second * 10)
}

func (c *WsConn) CloseGracingWithinTime(gracingDuration time.Duration) error {
	if c.conn != nil {
		defer func() {
		}()

		// 等待客户端主动关闭连接，超时服务端主动关闭
		err := c.conn.SetWriteDeadline(time.Now().Add(gracingDuration))
		if err != nil {
			return err
		}
		err = c.conn.WriteMessage(websocket.TextMessage, message.Close("").Bytes())
		if err != nil {
			return err
		}
		time.Sleep(gracingDuration)
		c.conn.Close()
	}
	return nil
}
