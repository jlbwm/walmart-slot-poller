FROM golang:alpine
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
# Create appuser.
ENV USER=appuser
ENV UID=10001
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR /walmart-slot-poller

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod verify

COPY . .

# Build the binary.
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/walmart-slot-poller


# Use an unprivileged user.
USER appuser:appuser
# Run the hello binary.
ENTRYPOINT ["/go/bin/walmart-slot-poller"]