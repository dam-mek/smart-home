package usecase

import (
	"context"
	"errors"
	"homework/internal/domain"
	"homework/internal/repository/sensor/inmemory"
	"regexp"
)

type Sensor struct {
	sRepo SensorRepository
}

var validTypes = map[domain.SensorType]struct{}{
	domain.SensorTypeContactClosure: {},
	domain.SensorTypeADC:            {},
}

func checkSensorType(sType domain.SensorType) bool {
	_, in := validTypes[sType]
	return in
}

var serialNumberCheck = regexp.MustCompile(`^[0-9]{10}$`)

func checkSerialNumber(sn string) bool {
	return serialNumberCheck.MatchString(sn)
}

func NewSensor(sr SensorRepository) *Sensor {
	return &Sensor{sRepo: sr}
}

func (s *Sensor) RegisterSensor(ctx context.Context, sensor *domain.Sensor) (*domain.Sensor, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	isValidType := checkSensorType(sensor.Type)
	if !isValidType {
		return nil, ErrWrongSensorType
	}
	isValidSN := checkSerialNumber(sensor.SerialNumber)
	if !isValidSN {
		return nil, ErrWrongSensorSerialNumber
	}

	existedSensor, err := s.sRepo.GetSensorBySerialNumber(ctx, sensor.SerialNumber)
	if err != nil {
		if errors.Is(err, inmemory.ErrSensorNotFound) {
			err = s.sRepo.SaveSensor(ctx, sensor)
			if err != nil {
				return nil, err
			}
			return sensor, nil
		}
		return nil, err
	}
	return existedSensor, nil
}

func (s *Sensor) GetSensors(ctx context.Context) ([]domain.Sensor, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return s.sRepo.GetSensors(ctx)
}

func (s *Sensor) GetSensorByID(ctx context.Context, id int64) (*domain.Sensor, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return s.sRepo.GetSensorByID(ctx, id)
}
