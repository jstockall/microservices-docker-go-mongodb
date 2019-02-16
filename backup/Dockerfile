# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.7

# Add MongoDB tools
RUN apt-get update && apt-get install -y mongodb-clients && rm -rf /var/lib/apt/lists/*

# Install dependencies
RUN go get gopkg.in/mgo.v2 github.com/gorilla/mux

# copy the local package files to the container workspace
ADD . /go/src/github.com/mmorejon/cinema/backup

# Setting up working directory
WORKDIR /go/src/github.com/mmorejon/cinema/backup

# Build the backup command inside the container.
RUN go install github.com/mmorejon/cinema/backup

# Run the users microservice when the container starts.
ENTRYPOINT /go/bin/backup

# Service listens on port 8080.
EXPOSE 8080
