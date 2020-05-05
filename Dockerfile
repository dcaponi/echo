# Dockerfile for app that gets built in the container
# Likely won't work for private repos

# Using golang (yes its larger than golang:alpine) but it has more out of box stuff
FROM golang:latest

# Standard copy from host to container stuff
WORKDIR /go/src/echo
COPY ./ /go/src/echo

# Download the modules we import
RUN go mod download

# Using this to trigger rebuild on code change since doing this ourselves takes effort
RUN go get github.com/githubnemo/CompileDaemon

# Fire it up!
ENTRYPOINT CompileDaemon --build="go build -o ./echo main.go" --command=./echo
