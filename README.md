# Cinema - Example of Microservices in Go with Docker and MongoDB
* Forked from mmorejon/microservices-docker-go-mongodb
* Changelog:
  * Switch to .dev domain name suffix for use with existing dnsmasq resolver as described in [Local docker development with virtual hosts](https://coderwall.com/p/qknu2g/local-docker-development-with-virtual-hosts)
  * Fix some movie IDs in the bookings and showtimes sample data (backup/*.bson)
  * Add URL handlers for /path/{id} for all objects

Overview
========

Cinema is an example project which demonstrates the use of microservices for a fictional movie theater.
The Cinema backend is powered by 4 microservices, all of witch happen to be written in Go, using MongoDB for manage the database and Docker to isolate and deploy the ecosystem.

 * Movie Service: Provides information like movie ratings, title, etc.
 * Show Times Service: Provides show times information.
 * Booking Service: Provides booking information.
 * Users Service: Provides movie suggestions for users by communicating with other services.

The Cinema use case is based on the project written in Python by [Umer Mansoor](https://github.com/umermansoor/microservices).

The project structure is based in the knowledge learned in the book: Web Development with Go by Shiju Varghese, ISBN 978-1-4842-1053-6

[Full multilanguage documentation with screenshots here.](http://mmorejon.github.io/en/blog/microservices-example-with-docker-go-and-mongodb/)

Requirements
===========

* Docker 1.12
* Docker Compose 1.8

We're using docker virtual domains to access each service behind the ngnix reverse proxy [from Jason Wilder (Automated Nginx Reverse Proxy for Docker)](https://github.com/jwilder/nginx-proxy).
We must setup our local DNS resolver in order to access each **api entry point**. **Virtual domains** have been defined in `docker-compose.yml` file as: **movies.dev**, **bookings.dev**, **users.dev** and **showtimes.dev**

To setup local DNS resolution add Add the following line in your `/etc/hosts` file:
```
127.0.0.1   movies.dev bookings.dev users.dev showtimes.dev
```
or setup dnsmasq with [Local docker development with virtual hosts](https://coderwall.com/p/qknu2g/local-docker-development-with-virtual-hosts)

Starting services
==============================

```
docker-compose up -d
```

Stoping services
==============================

```
docker-compose stop
```

Including new changes
==============================

If you need change some source code you can deploy it typing:

```
docker-compose build
```

Restore database information
======================

You can start using an empty database for all microservices, but if you want you can restore a preconfigured data following these steps:

**_Access to mongodb container typing:_**

```
docker exec -it cinema-db /bin/bash
```

**_Restore data typing:_**

```
/backup/restore.sh
```

**_Leave the container:_**

```
exit
```

Backup database information
======================

You can backup the data for all microservices following these steps:

**_Access to mongodb container typing:_**

```
docker exec -it cinema-db /bin/bash
```

**_Backup data typing:_**

```
/backup/backup.sh
```

**_Leave the container:_**

```
exit
```

Documentation
======================

## User Service

This service returns information about the users of Cinema.

**_Routes:_**

* GET - http://users.dev/users : Get all users
* GET - http://users.dev/users/{id} : Get user by id
* POST - http://users.dev/users : Create user
* DELETE - http://users.dev/users/{id} : Remove user by id

## Movie Service

This service is used to get information about a movie. It provides the movie title, rating on a 1-10 scale, director and other information.

**_Routes:_**

* GET - http://movies.dev/movies : Get all movies
* GET - http://movies.dev/movies/{id} : Get movie by id
* POST - http://movies.dev/movies : Create movie
* DELETE - http://movies.dev/movies/{id} : Remove movie by id

## Showtimes Service

This service is used get a list of movies playing on a certain date.

**_Routes:_**

* GET - http://showtimes.dev/showtimes : Get all showtimes
* GET - http://showtimes.dev/showtimes/{id} : Get showtime by id
* POST - http://showtimes.dev/showtimes : Create showtime
* DELETE - http://showtimes.dev/showtimes/{id} : Remove showtime by id

## Booking Service

Used to lookup booking information for users.

**_Routes:_**

* GET - http://bookings.dev/bookings : Get all bookings
* GET - http://bookings.dev/bookings/{id} : Get booking by id
* POST - http://bookings.dev/bookings : Create booking

### Significant Resources

* [Microservices - Martin Fowler](http://martinfowler.com/articles/microservices.html)
* [Web Development with Go](http://www.apress.com/9781484210536)
* [Umer Mansoor - Cinema](https://github.com/umermansoor/microservices)
* [Local docker development with virtual hosts](https://coderwall.com/p/qknu2g/local-docker-development-with-virtual-hosts)
* [Automated Nginx Reverse Proxy for Docker by Jason Wilder](https://github.com/jwilder/nginx-proxy).
