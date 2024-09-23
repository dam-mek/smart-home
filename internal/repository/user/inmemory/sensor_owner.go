package inmemory

import (
	"context"
	"homework/internal/domain"
	"sync"
)

type SensorOwnerRepository struct {
	sensorOwners []domain.SensorOwner
	soMu         sync.Mutex
}

func NewSensorOwnerRepository() *SensorOwnerRepository {
	return &SensorOwnerRepository{}
}

func (r *SensorOwnerRepository) SaveSensorOwner(ctx context.Context, sensorOwner domain.SensorOwner) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	r.soMu.Lock()
	r.sensorOwners = append(r.sensorOwners, sensorOwner)
	r.soMu.Unlock()
	return ctx.Err()
}

func (r *SensorOwnerRepository) GetSensorsByUserID(ctx context.Context, userID int64) ([]domain.SensorOwner, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	var sensors []domain.SensorOwner
	for _, s := range r.sensorOwners {
		if s.UserID == userID {
			sensors = append(sensors, s)
		}
	}
	return sensors, nil
}
