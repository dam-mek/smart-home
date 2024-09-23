package inmemory

import (
	"context"
	"errors"
	"homework/internal/domain"
	"sort"
	"sync"
)

var ErrEventNotFound = errors.New("event not found")

type EventRepository struct {
	events []*domain.Event
	eMu    sync.Mutex
}

func NewEventRepository() *EventRepository {
	return &EventRepository{}
}

func (r *EventRepository) SaveEvent(ctx context.Context, event *domain.Event) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if event == nil {
		return errors.New("event is nil")
	}
	r.eMu.Lock()
	r.events = append(r.events, event)
	r.eMu.Unlock()
	return nil
}

func (r *EventRepository) GetLastEventBySensorID(ctx context.Context, id int64) (*domain.Event, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	var events []*domain.Event
	for _, s := range r.events {
		if s.SensorID == id {
			events = append(events, s)
		}
	}
	if len(events) == 0 {
		return nil, ErrEventNotFound
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].Timestamp.Before(events[j].Timestamp)
	})
	return events[len(events)-1], nil
}
