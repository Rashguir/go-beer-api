FROM golang:1.12.6-alpine3.10 as go-build

WORKDIR /app

RUN apk update && \
    apk --no-cache add \
        git

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=go-build /app .

CMD ["./app"]

