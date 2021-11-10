package repository

import (
	"github.com/edenofjinx/go-bookings/internal/models"
	"time"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDates(start, end time.Time, roomId int) (bool, error)
}
