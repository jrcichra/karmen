FROM golang:1.18-alpine3.15 as firststage
WORKDIR /karmen
ADD . .
RUN go build -o karmen .
FROM alpine:3.15
WORKDIR /karmen
COPY --from=firststage /karmen/karmen .
CMD ["./karmen"]
