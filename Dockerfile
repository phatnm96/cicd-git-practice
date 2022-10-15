# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /cutloss-trading

EXPOSE 8080

CMD [ "/cutloss-trading" ]

