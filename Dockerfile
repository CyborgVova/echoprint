FROM golang:1.24.10-alpine AS builder

WORKDIR /app

COPY go.mod go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o echo . &&\
 chmod +x ./echo


FROM alpine:latest

RUN addgroup -g 1000 -S echogroup &&\
 adduser -u 1000 -G echogroup -D echouser --home /home/echouser

USER echouser

WORKDIR /home/echouser/app

COPY --from=builder --chown=echouser:echogroup /app/echo .

EXPOSE 8888

ENTRYPOINT ["./echo"]
