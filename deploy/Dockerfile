FROM golang:1.19.3-alpine as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build  -o pokemon-viewer-backend cmd/main.go

CMD [ "/app/pokemon-viewer-backend" ]