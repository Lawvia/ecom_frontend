FROM golang:1.13.5-alpine

RUN mkdir -p /go/src/ecom_frontend
WORKDIR /go/src/ecom_frontend

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o index .

EXPOSE 3000

CMD [ "./index" ]
