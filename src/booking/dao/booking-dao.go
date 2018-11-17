package bookingDao
import (
	
  "database/sql"
  "fmt"
  "time"
  _ "github.com/lib/pq"
    "driver/dao"

)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "postgres"
  dbname   = "finaccel"
)


type Booking struct {
  BookingId int
  Status  string
  BookingTime time.Time
  SourceX float64
  SourceY float64
  DestinationX float64
  DestinationY float64
  DriverId int
}

type BookingInProgress struct {
  BookingId int
  LocationX float64
  LocationY float64
  DriverId int
}


var db *sql.DB

func AddBooking(b *Booking, d *driverDao.Driver) {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  } 
  defer db.Close()
  	tx, err := db.Begin()

      // create a booking record with status = pending
	  sqlStatement:= "INSERT INTO booking (STATUS, BOOKING_TIME, SOURCE_X,SOURCE_Y,DESTINATION_X,DESTINATION_Y,DRIVER_ID) VALUES ($1, $2, $3, $4,$5,$6,$7) RETURNING BOOKING_ID;"
	  if _, err := tx.Exec(sqlStatement, &b.Status,&b.BookingTime,&b.SourceX,&b.SourceY,&b.DestinationX,&b.DestinationY,&b.DriverId) ; err != nil {
	  	 tx.Rollback() 
	  	 panic(err)
	  }

      // update driver status after booking made
	  if _, err := tx.Exec("Update Driver set status = $1 where driver_id = $2", &d.Status,&d.Id) ; err != nil {
		 tx.Rollback() 
	  	 panic(err)
	  }
	
	tx.Commit()
}


