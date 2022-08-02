FROM golang:1.19.0-alpine3.15 as firststage
WORKDIR /karmen
ADD . .
RUN apk add git gcc
RUN go build -o karmen .
FROM alpine:3.16.1
WORKDIR /karmen
COPY --from=firststage /karmen/karmen .
CMD ["./karmen"]
