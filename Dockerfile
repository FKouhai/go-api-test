FROM golang:1.13-alpine3.11 as builder
RUN mkdir /build
RUN mkdir /build/config
COPY config/* /build/config/
ADD *go* /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -o api main
FROM alpine:3.16
COPY --from=builder /build/api .
ENTRYPOINT [ "./api" ]
