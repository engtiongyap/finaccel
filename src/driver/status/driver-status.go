package driverStatus

import (
    "errors"
)

var DriverStatuses = newDriverStatusRegistry()

type DriverStatus struct {
	StatusCode string
    StringRepresentation string
    
}

func (c *DriverStatus) String() string {
    return c.StringRepresentation
}

func newDriverStatusRegistry() *DriverStatusRegistry {

    available := &DriverStatus{"1", "Available"}
    booked := &DriverStatus{"0", "Booked"}
    pendingBookingPickUp := &DriverStatus{"2", "Pending Booking Pick Up"}

    return &DriverStatusRegistry{
        Available:    available,
        Booked:  booked,
        PendingBookingPickUp:  pendingBookingPickUp,
        driverStatuses: []*DriverStatus{available, booked,pendingBookingPickUp},
    }
}

type DriverStatusRegistry struct {
    Available   *DriverStatus
    Booked *DriverStatus
    PendingBookingPickUp *DriverStatus
    driverStatuses []*DriverStatus
}

func (c *DriverStatusRegistry) List() []*DriverStatus {
    return c.driverStatuses
}

func (c *DriverStatusRegistry) Parse(s string) (*DriverStatus, error) {
    for _, status := range c.List() {
        if status.String() == s {
            return status, nil
        }
    }
    return nil, errors.New("Driver Status Not Found")
}
