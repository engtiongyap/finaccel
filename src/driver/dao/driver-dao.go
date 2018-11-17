package driverDao
import (
	
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
    "strings"
    "strconv"

)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "postgres"
  dbname   = "finaccel"
  //1 decimal degree = 111.32 km
  kmInDecimalDegree = 111.32
  
)


type Driver struct {
  Id int
  Name  string
  Status string
}


var db *sql.DB

/**

Function to call database function findDriverWithinDistance to find driver which is within target distance
**/
func FindDriverWithinDistanceInKM (distance float64,location_x float64,location_y float64, status string) []int {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  } 
  defer db.Close()
  
  rows, err := db.Query("select findDriverWithinDistance($1,$2,$3,$4,$5)", distance,location_x,location_y,status,kmInDecimalDegree)
  var driverIds []int
  for rows.Next() {
  	var id int;
    err = rows.Scan(&id)
    
    if err != nil {
      // handle this error
      panic(err)
    }
    driverIds = append(driverIds,id)

  }

  return driverIds;
}

func FindDriver(driverId int, driverStatus string) []*Driver {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  } 
  defer db.Close()
  
  	var argsValue [2]string

  sqlStatement := "select driver_id, name, status from driver"
  
  var argsCount int
  if driverId > 0 || len(strings.TrimSpace(driverStatus)) > 0{
  	sqlStatement += " where"
  	if driverId > 0{ 
  		argsCount +=1;
	  	sqlStatement += " driver_Id =$" +  strconv.Itoa(argsCount)
	  	argsValue[0] = strconv.Itoa(driverId)
  	}
  	if len(strings.TrimSpace(driverStatus)) > 0{ 
  		if argsCount >= 1{
  			sqlStatement += " and "
  		}
  		argsCount +=1;
	  	sqlStatement += " status =$" + strconv.Itoa(argsCount)
	  	argsValue[1] = driverStatus
  	}
  }

  rows, err := db.Query(sqlStatement,argsValue[0],argsValue[1])
  
  var drivers []*Driver
  for rows.Next() {
  	d := new(Driver)
    err = rows.Scan(&d.Id,&d.Name,&d.Status)
    
    if err != nil {
      // handle this error
      panic(err)
    }
    drivers = append(drivers,d)

  }
  return drivers;
}
