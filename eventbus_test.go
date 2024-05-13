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
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewEventBus(t *testing.T) {
	require.NotNil(t, NewEventBus())
}

func TestEventBus_Subscribe(t *testing.T) {
	eventBus := NewEventBus()
	subscribe := eventBus.Subscribe("test")
	require.NotNil(t, subscribe)
	require.Equal(t, 1, len(eventBus.subscribers["test"]))
	require.Equal(t, subscribe, eventBus.subscribers["test"][0])
}

func TestEventBus_Unsubscribe(t *testing.T) {
	eventBus := NewEventBus()
	subscribe := eventBus.Subscribe("test")
	require.Equal(t, 1, len(eventBus.subscribers["test"]))
	eventBus.Unsubscribe("test", subscribe)
	require.Equal(t, 0, len(eventBus.subscribers["test"]))
}

func TestEventBus_Publish(t *testing.T) {
	eventBus := NewEventBus()
	subscribe := eventBus.Subscribe("test")
	go func() {
		eventBus.Publish("test", Event{Payload: []byte("test")})
	}()
	event := <-subscribe
	require.Equal(t, "test", string(event.Payload))
}
