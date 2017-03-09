#! /bin/sh
mongodump -h db -d users -c users -o /backup/users -vvvvv
mongodump -h db -d movies -c movies -o /backup/movies -vvvvv
mongodump -h db -d showtimes -c showtimes -o /backup/showtimes -vvvvv
mongodump -h db -d bookings -c bookings -o /backup/bookings -vvvvv
