package entities

import (
	"encoding/json"
	"time"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/messaging/models"
)

type EventType int

const (
	EventType_SendMessage EventType = iota
	EventType_ReceiveMessage
)

type Event struct {
	Timestamp time.Time       `json:"timestamp"`
	EventType EventType       `json:"event_type"`
	EventData json.RawMessage `json:"event_data"`
}

type EventSendMessage struct {
	Content      string    `json:"content"`
	TargetUserID *uint     `json:"target_user_id"`
	GroupID      *uint     `json:"group_id"`
	Timestamp    time.Time `json:"timestamp"`
}

type EventReceiveMessage struct {
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id"`
	Username  string    `json:"username"`
	GroupID   *uint     `json:"group_id"`
	Timestamp time.Time `json:"timestamp"`
}

func NewEvent(eventType EventType, eventData any) *Event {
	jsonData, _ := json.Marshal(eventData)
	return &Event{
		Timestamp: time.Now(),
		EventType: eventType,
		EventData: jsonData,
	}
}

func NewEventSendMessageUser(content string, toID uint) *Event {
	return NewEvent(EventType_SendMessage, &EventSendMessage{
		Content:      content,
		TargetUserID: &toID,
		Timestamp:    time.Now(),
	})
}

func NewEventReceiveMessageUser(content string, user *models.User) *Event {
	return NewEvent(EventType_ReceiveMessage, &EventReceiveMessage{
		Content:   content,
		UserID:    user.ID,
		Username:  user.Username,
		Timestamp: time.Now(),
	})
}
