FROM golang:alpine
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
# Create appuser.
ENV USER=appuser
ENV UID=10001
ENV TOEMAIL1=ljz0508@gmail.com
ENV TOEMAIL2=ljx477@gmail.com
ENV USERNAME=walmart-bot
ENV PASSWORD=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
ENV PERIOD=3600
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
