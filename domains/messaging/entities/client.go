package entities

import (
	"context"
	"time"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/util"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type Client struct {
	ID           uint
	conn         *websocket.Conn
	writeChannel chan *Event
	ReadChannel  chan *Event
}

func NewClient(id uint, conn *websocket.Conn) *Client {
	return &Client{
		ID:           id,
		conn:         conn,
		writeChannel: make(chan *Event),
		ReadChannel:  make(chan *Event),
	}
}

func (c *Client) Start(ctx context.Context) {
	go c.pollRead(ctx)
	go c.pollWrite(ctx)
}

func (c *Client) pollRead(ctx context.Context) {
	defer func() {
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	logger := util.GetLoggerFromContext(ctx)

	for {
		var event Event
		err := c.conn.ReadJSON(&event)
		if err != nil {
			logger.Error("Read message from websocket client error", zap.Error(err))
			break
		}

		logger.Info("event received", zap.Any("event", event))

		c.ReadChannel <- &event
	}
}

func (c *Client) pollWrite(ctx context.Context) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.writeChannel:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.conn.WriteJSON(message)

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) SendEvent(event *Event) {
	c.writeChannel <- event
}
