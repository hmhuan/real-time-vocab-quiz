package messages

type EventType string

type Base[M any] struct {
	EventType EventType `json:"type"`
	Data      M         `json:"data"`
}
