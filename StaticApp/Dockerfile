FROM golang:1.19-buster

WORKDIR /app

# Download necessary Go modules
COPY . ./

RUN go mod download
RUN go build -o /appProd

EXPOSE 8080

CMD [ "/appProd" ]