// Copyright 2024 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eventbus

import (
	"sync"
)

type Event struct {
	Payload []byte
}

type (
	EventChan chan Event
)

type EventBus struct {
	mu          sync.RWMutex
	subscribers map[string][]EventChan
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]EventChan),
	}
}

func (eb *EventBus) Subscribe(topic string) EventChan {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	ch := make(EventChan)
	eb.subscribers[topic] = append(eb.subscribers[topic], ch)
	return ch
}

func (eb *EventBus) Unsubscribe(topic string, ch EventChan) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	if subscribers, ok := eb.subscribers[topic]; ok {
		for i, subscriber := range subscribers {
			if ch == subscriber {
				eb.subscribers[topic] = append(subscribers[:i], subscribers[i+1:]...)
				close(ch)
				// 清空通道
				for range ch {
				}
				return
			}
		}
	}
}

func (eb *EventBus) Publish(topic string, event Event) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()
	// 复制一个新的订阅者列表，避免在发布事件时修改订阅者列表
	subscribers := append([]EventChan{}, eb.subscribers[topic]...)
	go func() {
		for _, subscriber := range subscribers {
			subscriber <- event
		}
	}()
}
