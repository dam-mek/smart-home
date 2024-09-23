package usecase

import (
	"context"
	"homework/internal/domain"
	"time"
)

type Event struct {
	eRepo EventRepository
	sRepo SensorRepository
}

func NewEvent(er EventRepository, sr SensorRepository) *Event {
	return &Event{eRepo: er, sRepo: sr}
}

func (e *Event) ReceiveEvent(ctx context.Context, event *domain.Event) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if event.Timestamp.IsZero() {
		return ErrInvalidEventTimestamp
	}
	s, err := e.sRepo.GetSensorBySerialNumber(ctx, event.SensorSerialNumber)
	if err != nil {
		return err
	}
	event.SensorID = s.ID
	err = e.eRepo.SaveEvent(ctx, event)
	if err != nil {
		return err
	}
	s.CurrentState = event.Payload
	s.LastActivity = time.Now()
	err = e.sRepo.SaveSensor(ctx, s)
	if err != nil {
		return err
	}
	return nil
}

func (e *Event) GetLastEventBySensorID(ctx context.Context, id int64) (*domain.Event, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return e.eRepo.GetLastEventBySensorID(ctx, id)
}
