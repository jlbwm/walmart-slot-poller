FROM golang:alpine
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
# Create appuser.
ENV USER=appuser
ENV UID=10001
ENV TOEMAIL=jiangsid87@gmail.com
ENV USERNAME=apikey
ENV PASSWORD=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
ENV PERIOD=1800
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