package bookingStatus

import (
    "errors"
)

var BookingStatuses = newBookingStatusRegistry()

type BookingStatus struct {
	StatusCode string
    StringRepresentation string
    
}

func (c *BookingStatus) String() string {
    return c.StringRepresentation
}

func newBookingStatusRegistry() *BookingStatusRegistry {
    booked := &BookingStatus{"0", "Booked"}
    pending := &BookingStatus{"1", "Pending"}
    completed := &BookingStatus{"2", "Completed"}
    rejected := &BookingStatus{"3", "Rejected"}

    return &BookingStatusRegistry{
        Pending:    pending,
        Booked:  booked,
        Completed: completed,
        Rejected: rejected,
        bookingStatuses: []*BookingStatus{pending, booked,completed,rejected},
    }
}

type BookingStatusRegistry struct {
    Pending   *BookingStatus
    Booked *BookingStatus
    Completed *BookingStatus
    Rejected *BookingStatus

    bookingStatuses []*BookingStatus
}

func (c *BookingStatusRegistry) List() []*BookingStatus {
    return c.bookingStatuses
}

func (c *BookingStatusRegistry) Parse(s string) (*BookingStatus, error) {
    for _, status := range c.List() {
        if status.String() == s {
            return status, nil
        }
    }
    return nil, errors.New("Booking Status Not Found")
}
