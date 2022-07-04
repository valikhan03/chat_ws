FROM golang:1.18-alpine

WORKDIR /chatapp


COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY * ./

RUN go build -o /app-build

EXPOSE 8080

CMD [ "/app-build" ]

