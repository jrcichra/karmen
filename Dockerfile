FROM golang:1.25-bookworm as firststage
WORKDIR /karmen
ADD . .
RUN CGO_ENABLED=0 go build -o karmen .
FROM gcr.io/distroless/static-debian11
WORKDIR /karmen
COPY --from=firststage /karmen/karmen .
CMD ["./karmen"]
