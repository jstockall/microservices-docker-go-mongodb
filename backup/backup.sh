#! /bin/bash
mongodump -d users -c users -o users
mongodump -d movies -c movies -o movies
mongodump -d showtimes -c showtimes -o showtimes
mongodump -d bookings -c bookings -o bookings
