# Binaries are cross-compiled natively in CI (see .github/workflows/karmen.yml)
# and staged under bin/ — this image only repackages them, so multi-arch
# builds need no QEMU/binfmt support.
ARG TARGETARCH

FROM gcr.io/distroless/static-debian11
WORKDIR /karmen
COPY bin/karmen-${TARGETARCH} /karmen/karmen
CMD ["./karmen"]
