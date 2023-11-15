# -----------------------------------------------------------------------------
# Stages
# -----------------------------------------------------------------------------

# FIXME: ARG IMAGE_SENZINGAPI_RUNTIME=senzing/senzingapi-runtime:3.7.1
ARG IMAGE_SENZINGAPI_RUNTIME=senzing/senzingapi-runtime:staging

ARG IMAGE_GO_BUILDER=golang:1.21.4-bullseye@sha256:31848c4f02b08469e159ea1ee664a3f29602418b13e7d67dfd4560d169e14d55

# FIXME: ARG IMAGE_FINAL=senzing/senzingapi-runtime:3.7.1
ARG IMAGE_FINAL=senzing/senzingapi-runtime:3.8.0

# -----------------------------------------------------------------------------
# Stage: senzingapi_runtime
# -----------------------------------------------------------------------------

FROM ${IMAGE_SENZINGAPI_RUNTIME} as senzingapi_runtime

# -----------------------------------------------------------------------------
# Stage: go_builder
# -----------------------------------------------------------------------------

FROM ${IMAGE_GO_BUILDER} as go_builder
ENV REFRESHED_AT=2023-11-14
LABEL Name="senzing/senzing-tools-builder" \
      Maintainer="support@senzing.com" \
      Version="0.5.0"

# Copy local files from the Git repository.

COPY ./rootfs /
COPY . ${GOPATH}/src/senzing-tools

# Copy files from prior stage.

COPY --from=senzingapi_runtime  "/opt/senzing/g2/lib/"   "/opt/senzing/g2/lib/"
COPY --from=senzingapi_runtime  "/opt/senzing/g2/sdk/c/" "/opt/senzing/g2/sdk/c/"

# Set path to Senzing libs.

ENV LD_LIBRARY_PATH=/opt/senzing/g2/lib/

# Build go program.

WORKDIR ${GOPATH}/src/senzing-tools
RUN make build

# Copy binaries to /output.

RUN mkdir -p /output \
 && cp -R ${GOPATH}/src/senzing-tools/target/*  /output/

# -----------------------------------------------------------------------------
# Stage: final
# -----------------------------------------------------------------------------

FROM ${IMAGE_FINAL} as final
ENV REFRESHED_AT=2023-11-14
LABEL Name="senzing/senzing-tools" \
      Maintainer="support@senzing.com" \
      Version="0.5.0"

# Copy local files from the Git repository.

COPY ./rootfs /

# Copy files from prior stage.

COPY --from=go_builder "/output/linux-amd64/senzing-tools" "/app/senzing-tools"

# Runtime environment variables.

ENV LD_LIBRARY_PATH=/opt/senzing/g2/lib/

# Runtime execution.

WORKDIR /app
ENTRYPOINT ["/app/senzing-tools"]
