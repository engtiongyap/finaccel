package driverService
import (
  "driver/dao"
   "driver/status"
)

func FindAvailableDriverWithinDistanceInKM (distance float64,location_x float64,location_y float64) [] int {
  return FindDriverWithinDistanceInKM(distance,location_x,location_y,driverStatus.DriverStatuses.Available.StatusCode)
}


/**
able to find driver by status
status : optional
**/
func FindDriverWithinDistanceInKM (distance float64,location_x float64,location_y float64, driverStatus string) [] int {
  return driverDao.FindDriverWithinDistanceInKM(distance,location_x,location_y,driverStatus)
}




func FindDriver (driverId int, driverStatus string) []*driverDao.Driver {
  return driverDao.FindDriver(driverId,driverStatus)
}



