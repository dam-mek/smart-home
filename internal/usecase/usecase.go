package usecase

import (
	"context"
	"errors"
	"homework/internal/domain"
)

var (
	ErrWrongSensorSerialNumber = errors.New("wrong sensor serial number")
	ErrWrongSensorType         = errors.New("wrong sensor type")
	ErrInvalidEventTimestamp   = errors.New("invalid event timestamp")
	ErrInvalidUserName         = errors.New("invalid user name")
)

//go:generate mockgen -source usecase.go -package usecase -destination usecase_mock.go
type SensorRepository interface {
	SaveSensor(ctx context.Context, sensor *domain.Sensor) error
	GetSensors(ctx context.Context) ([]domain.Sensor, error)
	GetSensorByID(ctx context.Context, id int64) (*domain.Sensor, error)
	GetSensorBySerialNumber(ctx context.Context, sn string) (*domain.Sensor, error)
}

type EventRepository interface {
	SaveEvent(ctx context.Context, event *domain.Event) error
	GetLastEventBySensorID(ctx context.Context, id int64) (*domain.Event, error)
}

type UserRepository interface {
	SaveUser(ctx context.Context, user *domain.User) error
	GetUserByID(ctx context.Context, id int64) (*domain.User, error)
}

type SensorOwnerRepository interface {
	SaveSensorOwner(ctx context.Context, sensorOwner domain.SensorOwner) error
	GetSensorsByUserID(ctx context.Context, userID int64) ([]domain.SensorOwner, error)
}
