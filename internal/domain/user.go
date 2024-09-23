package domain

type User struct {
	ID   int64
	Name string
}

type SensorOwner struct {
	UserID   int64
	SensorID int64
}
