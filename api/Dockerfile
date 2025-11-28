FROM golang:1.25-alpine

# Install ffmpeg, bash, and git
RUN apk add --no-cache ffmpeg bash git

WORKDIR /app

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o app cmd/app/main.go

EXPOSE 8080

CMD ["./app"]
