FROM golang:1.19-buster

WORKDIR /app

# Download necessary Go modules
COPY . ./

RUN go mod download
RUN go run /app/cmd/api/main.go

EXPOSE 8080
