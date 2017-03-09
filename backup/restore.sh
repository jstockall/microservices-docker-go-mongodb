#! /bin/sh
mongorestore -h $DATABASE_HOST -d users -c users --drop users/users/users.bson
mongorestore -h $DATABASE_HOST -d movies -c movies --drop movies/movies/movies.bson
mongorestore -h $DATABASE_HOST -d showtimes -c showtimes --drop showtimes/showtimes/showtimes.bson
mongorestore -h $DATABASE_HOST -d bookings -c bookings --drop bookings/bookings/bookings.bson
