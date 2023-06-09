MAINTAINER Arthur Davis <arthur8davis@gmail.com>
FROM golang:1.19.0 AS builder

ENV GOPRIVATE=dev.azure.com

RUN mkdir /home/storage

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 G00S=linux go build -o app main.go

FROM alpine AS dockerize

COPY --from=builder /src/app .

CMD ["./app"]