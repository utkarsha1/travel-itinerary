# travel-itinerary

This application persists activities across destinations world over. These events can then ne retrieved to build dynamic itineraries based on duration, rating, activity level, etc.


#### Technology Stack
- MongoDB : persistence storage
- Go : microservice application to perform CRUD operations

##### Dump and Restore DB
- ```utkarsha$ mongodump
     2018-08-05T12:39:25.576-0700	writing admin.system.version to 
     2018-08-05T12:39:25.577-0700	done dumping admin.system.version (1 document)
     2018-08-05T12:39:25.577-0700	writing travelitinerary.activity to 
     2018-08-05T12:39:25.578-0700	done dumping travelitinerary.activity (9 documents)
     
-     mongorestore --collection activity --db travelitinerary  dump/travelitinerary/activity.bson```
