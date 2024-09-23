package inmemory

import (
	"context"
	"errors"
	"homework/internal/domain"
	"sync"
	"time"
)

var ErrSensorNotFound = errors.New("sensor not found")

type SensorRepository struct {
	sensors []domain.Sensor
	sMu     sync.Mutex
}

func NewSensorRepository() *SensorRepository {
	return &SensorRepository{}
}

func (r *SensorRepository) SaveSensor(ctx context.Context, sensor *domain.Sensor) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if sensor == nil {
		return errors.New("sensor is nil")
	}
	r.sMu.Lock()
	sensor.RegisteredAt = time.Now()
	sensor.ID = int64(len(r.sensors)) + 1
	r.sensors = append(r.sensors, *sensor)
	r.sMu.Unlock()
	return nil
}

func (r *SensorRepository) GetSensors(ctx context.Context) ([]domain.Sensor, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return r.sensors, nil
}

func (r *SensorRepository) GetSensorByID(ctx context.Context, id int64) (*domain.Sensor, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	for _, s := range r.sensors {
		if s.ID == id {
			return &s, nil
		}
	}
	return nil, ErrSensorNotFound
}

func (r *SensorRepository) GetSensorBySerialNumber(ctx context.Context, sn string) (*domain.Sensor, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	for _, s := range r.sensors {
		if s.SerialNumber == sn {
			return &s, nil
		}
	}
	return nil, ErrSensorNotFound
}
