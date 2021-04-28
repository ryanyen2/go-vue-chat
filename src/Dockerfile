FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.* .
RUN go mod download

COPY . .

RUN go build -o server .

###############
FROM alpine:3
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server /
COPY ./front-dist /front-dist
COPY ./credentials /credentials

EXPOSE 8080
CMD ["/server"]

