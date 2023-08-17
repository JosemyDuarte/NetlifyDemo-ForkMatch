ARG GO_VERSION

FROM golang:${GO_VERSION} AS build-env

ENV APP "ForkMatch"

WORKDIR /src
COPY . .

# Compilation options:
# - CGO_ENABLED=0: Disable cgo
# - GOOS=linux: explicitly target Linux
# - GOARCH: explicitly target 64bit CPU
# - -trimpath: improve reproducibility by trimming the pwd from the binary
# - -ldflags: extra linker flags
#   - -s: omit the symbol table and debug information making the binary smaller
#   - -w: omit the DWARF symbol table making the binary smaller
#   - -extldflags 'static': extra linker flags: produce a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
      -trimpath \
      -ldflags "-s -w -extldflags '-static'" \
      -o dist/${APP} \
        ./cmd/${APP}

# Use distroless as minimal base image to package the app binary
FROM gcr.io/distroless/base-debian11@sha256:73deaaf6a207c1a33850257ba74e0f196bc418636cada9943a03d7abea980d6d

ENV APP "ForkMatch"

# Have to rename the binary because build args don't work on ENTRYPOINT
COPY --from=build-env /src/dist/${APP} /app

USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT ["/app"]
CMD []
