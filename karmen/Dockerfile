FROM golang:alpine3.11 as firststage
WORKDIR /karmen
ADD . .
RUN go build -o karmen .
FROM alpine:3.11
WORKDIR /karmen
COPY --from=firststage /karmen/karmen .
CMD ["./karmen"]