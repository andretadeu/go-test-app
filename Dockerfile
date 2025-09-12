FROM golang:1.25.1-alpine AS builder

# Move to working directory /build
WORKDIR /build

COPY main main
COPY go.mod .
COPY go.sum .

RUN go build -o test-app ./main

FROM alpine:3.22 AS prod_image

RUN mkdir -p /opt/test-app
COPY --from=builder /build/test-app /opt/test-app/test-app

ENTRYPOINT [ "/opt/test-app/test-app" ]