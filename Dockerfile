FROM golang:1.21.5 AS build

WORKDIR /app

COPY . .

RUN go build -o hello-world .

FROM alpine:3.19.1

COPY --from=build /app/hello-world /

CMD ["/hello-world"]