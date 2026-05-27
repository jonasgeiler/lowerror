# syntax=docker/dockerfile:1.22.0@sha256:4a43a54dd1fedceb30ba47e76cfcf2b47304f4161c0caeac2db1c61804ea3c91
# check=error=true

# This Dockerfile is used by Docker Compose during development to build the
# image from the local source code.
# The actual production image is built by GoReleaser using
# `./.goreleaser.Dockerfile`, which relies on the pre-built binaries.

FROM golang:1.26.3-alpine@sha256:91eda9776261207ea25fd06b5b7fed8d397dd2c0a283e77f2ab6e91bfa71079d AS build-stage

# Make default shell more strict and easier to debug.
SHELL ["/bin/ash", "-euxo", "pipefail", "-c"]
WORKDIR /app

# Download and cache Go dependencies.
# We use `./go.*` instead of `./go.mod ./go.sum` so a "go.sum" file is optional.
COPY ./go.* ./
RUN --mount=type=cache,target=/go/pkg/mod \
	go mod download

# Add Go source code and build.
COPY ./*.go ./
ENV CGO_ENABLED=0
ARG TARGETPLATFORM
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o "/${TARGETPLATFORM}/lowerror" -trimpath ./*.go

################################################################################

# Same as `.goreleaser.Dockerfile`, but with binaries from the build-stage:
FROM dhi.io/static:20251003-alpine3.23@sha256:a08d9a53a4758b4006d56341aa88b1edf583ddebd93e620a32acd5135535573c

WORKDIR /

# Copy the built binary from the build context using the target platform folder.
ARG TARGETPLATFORM
COPY --from=build-stage "${TARGETPLATFORM}/lowerror" /lowerror

EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/lowerror"]
