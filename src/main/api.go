package main


import (
	"encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
    "driver/service"
    "booking/service"
    "errors"
)


func main() {
    router := mux.NewRouter()
    router.HandleFunc("/driver/find-driver-within-distance", FindDriverWithinDistanceInKM).Methods("GET")
    router.HandleFunc("/booking/create-booking", CreateBooking).Methods("POST")
    log.Fatal(http.ListenAndServe(":8181", router))
}



func FindDriverWithinDistanceInKM(w http.ResponseWriter, r *http.Request) {
	var validInput bool =true;
	
	var invalidMsg string;
	
	distance, ok := r.URL.Query()["distance"]
	
	 if !ok || len(distance[0]) < 1 {
        invalidMsg = "Url Param 'distance' is missing"
        validInput = false
    }else{
		_, err := strconv.ParseFloat(distance[0], 64)
		if err != nil  {
	        invalidMsg += "\nUrl Invalid Param Value for Distance"
	        validInput = false
	    }
    }
    decimalDegreeParamArray := []string{"locationX", "locationY"}

	if _,err :=validateDecimalDegreeParam(r,decimalDegreeParamArray); err !=nil{
		validInput=false
		invalidMsg+=err.Error()
	}

	if !validInput{
	    http.Error(w, invalidMsg, http.StatusBadRequest)
		return
	}
	
	// confirm safe as aleady check in validation
	distanceInFloat, _ := strconv.ParseFloat(distance[0], 64)
	locationXInFloat, _:= strconv.ParseFloat(r.URL.Query()["locationX"][0], 64)
	locationYInFloat, _:= strconv.ParseFloat(r.URL.Query()["locationY"][0], 64)

	
	//optional driverStatus, if no status provided, driver with status = available (1) will be return
	driverStatus, ok := r.URL.Query()["status"]
	
	 if !ok || len(driverStatus[0]) < 1 {
	    json.NewEncoder(w).Encode(driverService.FindAvailableDriverWithinDistanceInKM(distanceInFloat,locationXInFloat,locationYInFloat))
	 }else{
	 	json.NewEncoder(w).Encode(driverService.FindDriverWithinDistanceInKM(distanceInFloat,locationXInFloat,locationYInFloat,driverStatus[0]))
	 }
}

/**
Booking Flow
1. Create booking with particular driver id, booking_status = pendingDriverPickUp
2. Driver to decide whether to accept booking
3. If Driver decided to pick up, Booking status will be updated to Booked, booking_status = booked.
4. Else If Driver Reject the booking, booking_status = rejected.
5. If Driver is not available for booking, system will reject the booking.
6. Once the booking is confirm booked, a record will be created in booking_in_progress table.
7. Once the booking is completed, the booking record in booking table will be deleted and archive to the booking_history table.

**/
func CreateBooking(w http.ResponseWriter, r *http.Request) {
	decimalDegreeParamArray := []string{"sourceX", "sourceY", "destinationX","destinationY"}

	// confirmed safe as aleady check in validation
	sourceXInFloat, _:= strconv.ParseFloat(r.URL.Query()["sourceX"][0], 64);
	sourceYInFloat, _:= strconv.ParseFloat(r.URL.Query()["sourceY"][0], 64);
	destinationXInFloat, _:= strconv.ParseFloat(r.URL.Query()["destinationX"][0], 64);
	destinationYInFloat, _:= strconv.ParseFloat(r.URL.Query()["destinationY"][0], 64);
	
	if _,err :=validateDecimalDegreeParam(r,decimalDegreeParamArray); err !=nil{
		    http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b, err := bookingService.CreateBooking(2,sourceXInFloat,sourceYInFloat,destinationXInFloat,destinationYInFloat)
	
	if err !=nil{
	    http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
    json.NewEncoder(w).Encode(b)
}


func validateDecimalDegreeParam(r *http.Request, paramToValidate[] string ) (bool,error){
	  validInput:=true
	  var invalidMsg string

	 for _,element := range paramToValidate{
          param, ok := r.URL.Query()[element]
          validElementInput := true
		 if !ok || len(param[0]) < 1 {
	        invalidMsg +="\nUrl Param '"+element+"' is missing"
	        validElementInput = false;
	    }
		 
		 if(validElementInput){
				 _, err := strconv.ParseFloat(param[0], 64)
				if err != nil  {
			        invalidMsg += "\nUrl Invalid Param Value for "+element+"InFloat"
			        validElementInput = false;
		    } 
		 }
		 if !validElementInput{
			 validInput = validElementInput
		 }
    }   
	 if !validInput {
		 return false, errors.New(invalidMsg)
	 }
	
	return true,nil;
}
