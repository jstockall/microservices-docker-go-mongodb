#! /bin/bash
mongorestore -d users -c users users/users/users.bson
mongorestore -d movies -c movies movies/movies/movies.bson
mongorestore -d showtimes -c showtimes showtimes/showtimes/showtimes.bson
mongorestore -d bookings -c bookings bookings/bookings/bookings.bson
