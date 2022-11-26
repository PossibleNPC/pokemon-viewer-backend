FROM golang:1.19.3-alpine as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build  -o pokemon-viewer-backend cmd/main.go

from alpine:latest as production

WORKDIR /app

COPY --from=builder /app/pokemon-viewer-backend /app

CMD [ "/app/pokemon-viewer-backend" ]