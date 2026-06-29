// TODO добить реализацию (gracefull shutdown возможность отписки настрока буферов улучшение метода публикации)

package bus

import (
	"sync"
)

type EventType string

const (
	RoundStarted EventType = "round_started"
	BetsClosed   EventType = "bets_closed"
	RollFinished EventType = "roll_finished"
)

type Event struct {
	Type EventType
	Data any
}

type EventBus struct {
	mu          sync.RWMutex
	subscribers map[EventType][]chan Event
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[EventType][]chan Event),
	}
}

// создаем канал под слушание ивента
func (b *EventBus) Subscribe(eventType EventType) <-chan Event {
	b.mu.Lock()
	defer b.mu.Unlock()

	ch := make(chan Event, 100) // поиграться с буфером канала найти оптимальный вариант

	b.subscribers[eventType] = append(b.subscribers[eventType], ch)

	return ch
}

func (b *EventBus) UnSubscribe() {
	//TODO функция отписки
}

func (b *EventBus) Publish(event Event) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, ch := range b.subscribers[event.Type] {
		select {
		case ch <- event:
		default:
			//Егор: канал переполнен - пропускаем
		}
	}
}
