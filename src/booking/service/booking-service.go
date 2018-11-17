package bookingService
import (
  "booking/dao"
   "driver/dao"
  "booking/status"
   "driver/status"
   "errors"
  "time"

)

func CreateBooking(driverId int,sourceX float64,sourceY float64, destinationX float64, destinationY float64) (*bookingDao.Booking,error) {
	
	// find if driver is available for book.
	
	 d := driverDao.FindDriver(driverId,driverStatus.DriverStatuses.Available.StatusCode)
	
	if d == nil {
		return nil,errors.New("Driver Not Available For Booking")
	}
	
  	 b := new(bookingDao.Booking)
	 b.BookingTime = time.Now()
	 b.DriverId = driverId
	 b.SourceX = sourceX
	 b.SourceY = sourceY
	 b.DestinationX = destinationX
	 b.DestinationY = destinationY
	 b.Status = bookingStatus.BookingStatuses.Booked.StatusCode
	 
	 d[0].Status = driverStatus.DriverStatuses.PendingBookingPickUp.StatusCode
	 
	 bookingDao.AddBooking(b,d[0]);
	 
     return b,nil;
}

