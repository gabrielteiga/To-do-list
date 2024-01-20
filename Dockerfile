FROM golang:1.21 AS builder

WORKDIR /usr/src/app

COPY . .
RUN go build -v -o /usr/local/bin/app ./

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /usr/local/bin/app /usr/local/bin/app

ENTRYPOINT ["app"]