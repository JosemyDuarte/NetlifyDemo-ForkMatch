ARG GO_VERSION

FROM golang:${GO_VERSION} AS build-env

ARG APP

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN make build

# Use distroless as minimal base image to package the app binary
FROM gcr.io/distroless/base-debian11@sha256:73deaaf6a207c1a33850257ba74e0f196bc418636cada9943a03d7abea980d6d

ARG APP

# Have to rename the binary because build args don't work on ENTRYPOINT
COPY --from=build-env /src/dist/${APP} /app

USER nonroot:nonroot

EXPOSE 80

ENTRYPOINT ["/app"]
CMD []
