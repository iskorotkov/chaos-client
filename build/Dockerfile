#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 go install -v ./...

#final stage
FROM debian:stable
RUN apt update && apt install -y ca-certificates
COPY --from=builder /go/bin/chaos-client /app
ENTRYPOINT ["./app"]
