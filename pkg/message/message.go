// Package message provides the interface and implementation of messaging.
package message

import (
	"sync"
)

// Message represents an event payload.
type Message struct {
	Payload interface{}
}

// MessageChan represents a Message channel.
type MessageChan chan Message

// MessageChanSlice represents a slice of MessageChan.
type MessageChanSlice []MessageChan

// CallbackFunc represents a Message callback handler.
type CallbackFunc func(Message) error

// Bus represents a publish-subscribe event bus.
type Bus struct {
	subs map[string]MessageChanSlice
	mu   sync.RWMutex
}

func (b *Bus) Publish(topic string, data interface{}) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if ch, found := b.subs[topic]; found {
		for _, c := range ch {
			c <- Message{Payload: data}
		}
	}
}

func (b *Bus) Subscribe(topic string, c MessageChan) {
	b.mu.Lock()
	defer b.mu.Unlock()
	subs, found := b.subs[topic]
	switch found {
	case true:
		b.subs[topic] = append(subs, c)
	case false:
		b.subs[topic] = MessageChanSlice{c}
	}
}
