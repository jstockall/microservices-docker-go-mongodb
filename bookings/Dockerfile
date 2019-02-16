# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.7

# Install dependencies
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux

# copy the local package files to the container workspace
ADD ./common /go/src/github.com/mmorejon/cinema/common
ADD ./bookings /go/src/github.com/mmorejon/cinema/bookings

# Setting up working directory
WORKDIR /go/src/github.com/mmorejon/cinema/bookings

# Build the bookings command inside the container.
RUN go install github.com/mmorejon/cinema/bookings

# Run the bookings microservice when the container starts.
ENTRYPOINT /go/bin/bookings

# Service listens on port 8080.
EXPOSE 8080
