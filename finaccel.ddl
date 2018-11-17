

--driver table
create table driver (
   DRIVER_ID INT PRIMARY KEY     NOT NULL,
   NAME TEXT NOT NULL,
   STATUS CHAR(1) NOT NULL
);
CREATE SEQUENCE driver_id_seq;
ALTER TABLE driver ALTER driver_id SET DEFAULT NEXTVAL('driver_id_seq');
insert into driver  (name,status) values ('John','1');
insert into driver  (name,status) values ('May','1');
insert into driver  (name,status) values ('Jason','1');
insert into driver  (name,status) values ('Michael','1');
insert into driver  (name,status) values ('Jenny','1');

--the table which stored the all the available driver.
create table driver_location (
   LOCATION_X REAL NOT NULL,
   LOCATION_Y REAL NOT NULL,
   DRIVER_ID INT PRIMARY KEY REFERENCES driver(DRIVER_ID),
);

insert into driver_location  (location_x,location_y,driver_id) values (20.005,20.005,1);
insert into driver_location  (location_x,location_y,driver_id) values (22.001,17.005,2);
insert into driver_location  (location_x,location_y,driver_id) values (21.00,16.00,3);
insert into driver_location  (location_x,location_y,driver_id) values (22.78,18.00,4);
insert into driver_location  (location_x,location_y,driver_id) values (23.22,21.22,5);

create table booking (
   BOOKING_ID INT PRIMARY KEY     NOT NULL,
   STATUS CHAR(1),
   BOOKING_TIME TIMESTAMP,
   SOURCE_X REAL NOT NULL,
   SOURCE_Y REAL NOT NULL,
   DESTINATION_X REAL NOT NULL,
   DESTINATION_Y REAL NOT NULL,
 
);
CREATE SEQUENCE booking_id_seq;
ALTER TABLE booking ALTER booking_id SET DEFAULT NEXTVAL('booking_id_seq');

--table which hold all the booking in progress, the record will only appear here when the booking is currently in progress
create table booking_in_progress (
   BOOKING_ID INT NOT NULL,
   DRIVER_ID INT PRIMARY KEY REFERENCES driver(DRIVER_ID)
);

-- table which holds the booking_history, completed booking record will be moved here.
create table booking_history (
   BOOKING_ID INT PRIMARY KEY     NOT NULL,
   BOOKING_TIME TIMESTAMP,
   COMPLETION_TIME TIMESTAMP,
   FINAL_STATUS CHAR(1),
   SOURCE_X REAL NOT NULL,
   SOURCE_Y REAL NOT NULL,
   DESTINATION_X REAL NOT NULL,
   DESTINATION_Y REAL NOT NULL,
   DRIVER_ID INT REFERENCES driver(DRIVER_ID)

)


create or replace function findDriverWithinDistance(distance int,source_location_x real,source_location_y real, driverStatus char, kmInDecimalDegree real)
		returns setof int AS $$
		select driver.driver_id  from driver , driver_location
		where driver.driver_id = driver_location.driver_id	
		and driver.status = driverStatus
		and location_x >= (source_location_x - (distance / kmInDecimalDegree)) and location_x <= (source_location_x + (distance * kmInDecimalDegree))
		and location_y >= (source_location_y - (distance / kmInDecimalDegree)) and location_y <= (source_location_y + (distance * kmInDecimalDegree))																									
		and (sqrt(power(abs(abs(location_x) - source_location_x),2) + power(abs(abs(location_y) - source_location_y),2)) * kmInDecimalDegree)< distance
		$$ LANGUAGE SQL;	
