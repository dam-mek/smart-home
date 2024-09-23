package usecase

import (
	"context"
	"homework/internal/domain"
)

type User struct {
	sRepo  SensorRepository
	soRepo SensorOwnerRepository
	uRepo  UserRepository
}

func NewUser(ur UserRepository, sor SensorOwnerRepository, sr SensorRepository) *User {
	return &User{uRepo: ur, soRepo: sor, sRepo: sr}
}

func (u *User) RegisterUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if user.Name == "" {
		return nil, ErrInvalidUserName
	}
	err := u.uRepo.SaveUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) AttachSensorToUser(ctx context.Context, userID, sensorID int64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if _, err := u.uRepo.GetUserByID(ctx, userID); err != nil {
		return err
	}
	if _, err := u.sRepo.GetSensorByID(ctx, sensorID); err != nil {
		return err
	}

	err := u.soRepo.SaveSensorOwner(ctx, domain.SensorOwner{SensorID: sensorID, UserID: userID})
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetUserSensors(ctx context.Context, userID int64) ([]domain.Sensor, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if _, err := u.uRepo.GetUserByID(ctx, userID); err != nil {
		return nil, err
	}
	sensOwns, err := u.soRepo.GetSensorsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var sensors []domain.Sensor
	for _, sensOwn := range sensOwns {
		sens, err := u.sRepo.GetSensorByID(ctx, sensOwn.SensorID)
		if err != nil {
			return nil, err
		}
		sensors = append(sensors, *sens)
	}
	return sensors, nil
}
