# Build Go App
FROM golang:1.16.0-alpine3.13 AS build-env
# All these steps will be cached
RUN mkdir /app
WORKDIR /app

COPY go.mod . 
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

# COPY the source code as the last step
COPY . .
WORKDIR /internal/cmd/server

# Build the binary for the site
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -o /go/bin/app

# <- Second step to build minimal image
FROM scratch 
COPY --from=build-env /go/bin/app /go/bin/app

# Setup a mount point for volumes
# VOLUME [ "/var/assets" ]

# Set environment variables
ENV BRANCH=main
# ENV ACTIVITY_FILE_PATH=/var/assets/activity
ENV PORT=8080

# Expose port 8080
EXPOSE 8080
ENTRYPOINT ["/go/bin/app"]