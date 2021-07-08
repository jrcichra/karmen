FROM golang:alpine3.14 as firststage
WORKDIR /karmen
ADD . .
RUN go build -o karmen .
FROM alpine:3.14
WORKDIR /karmen
COPY --from=firststage /karmen/karmen .
CMD ["./karmen"]
