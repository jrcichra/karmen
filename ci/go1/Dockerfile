FROM golang:alpine3.11 as firststage
WORKDIR /go1
ADD . .
RUN go build -o go1 .
FROM alpine:3.11
WORKDIR /go1
COPY --from=firststage /go1/go1 .
CMD ["./go1"]