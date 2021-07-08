FROM golang:alpine3.13 as firststage
WORKDIR /karmen
ADD . .
RUN go build -o karmen .
FROM alpine:3.13
WORKDIR /karmen
COPY --from=firststage /karmen/karmen .
CMD ["./karmen"]
