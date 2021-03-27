FROM golang:alpine3.13 as firststage
WORKDIR /go1
ADD . .
RUN go build -o go1 .
FROM alpine:3.13
WORKDIR /go1
COPY --from=firststage /go1/go1 .
EXPOSE 6060
CMD ["./go1"]