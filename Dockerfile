FROM golang:1.17.8-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY main.go ./
COPY assets ./assets

RUN go mod download

RUN go build -o /tallio

EXPOSE 8443

CMD [ "/tallio" ]