FROM golang:1.18.5-bullseye as firststage
WORKDIR /karmen
ADD . .
RUN go build -o karmen .
FROM gcr.io/distroless/base-debian11
WORKDIR /karmen
COPY --from=firststage /karmen/karmen .
CMD ["./karmen"]
