FROM golang:alpine AS builder

LABEL stage=gobuilder
ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata

ADD go.mod .
ADD go.sum .

RUN go mod download
COPY ../.. .
RUN go build -o /app/server cmd/server/main.go

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates
COPY --from=builder /usr/share/zoneinfo/Europe/Moscow /usr/share/zoneinfo/Europe/Moscow
ENV TZ Europe/Moscow

COPY --from=builder /app/server /app/server

CMD ["./app/server"]
