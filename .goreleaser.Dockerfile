# syntax=docker/dockerfile:1.22.0@sha256:4a43a54dd1fedceb30ba47e76cfcf2b47304f4161c0caeac2db1c61804ea3c91
# check=error=true

# This Dockerfile is used by GoReleaser to build the production Docker image.
# It is not intended to be used directly, as it relies on the build context
# provided by GoReleaser and the already built binaries.
# For development, check out Docker Compose and `./compose.Dockerfile`.

FROM dhi.io/static:20251003-alpine3.23@sha256:a08d9a53a4758b4006d56341aa88b1edf583ddebd93e620a32acd5135535573c

WORKDIR /

# Copy the built binary from the build context using the target platform folder.
ARG TARGETPLATFORM
COPY "${TARGETPLATFORM}/lowerror" /lowerror

EXPOSE 80
USER nonroot:nonroot

ENTRYPOINT ["/lowerror"]
