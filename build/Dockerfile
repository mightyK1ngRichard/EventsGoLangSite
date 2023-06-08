FROM golang:1.18-buster

WORKDIR /opt/app

COPY . .

RUN go mod download

RUN go build -o app cmd/main/main.go

#EXPOSE 8010

CMD ["./app"]