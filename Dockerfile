# -----------------------------------------------------------------------------
# Stages
# -----------------------------------------------------------------------------

# ARG IMAGE_GO_BUILDER=golang:1.20.5@sha256:6b3fa4b908676231b50acbbc00e84d8cee9c6ce072b1175c0ff352c57d8a612f - Bad
ARG IMAGE_GO_BUILDER=golang:1.20.4@sha256:690e4135bf2a4571a572bfd5ddfa806b1cb9c3dea0446ebadaf32bc2ea09d4f9
ARG IMAGE_FINAL=senzing/senzingapi-runtime:3.6.0

# -----------------------------------------------------------------------------
# Stage: go_builder
# -----------------------------------------------------------------------------

# define where we need to copy senzing files from
FROM ${IMAGE_FINAL} as senzing-runtime
FROM ${IMAGE_GO_BUILDER} as go_builder
ENV REFRESHED_AT=2023-06-29
LABEL Name="senzing/senzing-tools-builder" \
      Maintainer="support@senzing.com" \
      Version="0.2.6"

# Build arguments.

ARG PROGRAM_NAME="unknown"
ARG BUILD_VERSION=0.0.0
ARG BUILD_ITERATION=0
ARG GO_PACKAGE_NAME="unknown"

# Copy remote files from DockerHub.

COPY --from=senzing/senzingapi-runtime:3.5.3  "/opt/senzing/g2/lib/"   "/opt/senzing/g2/lib/"
COPY --from=senzing/senzingapi-runtime:3.5.3  "/opt/senzing/g2/sdk/c/" "/opt/senzing/g2/sdk/c/"

# Copy local files from the Git repository.

COPY ./rootfs /
COPY . ${GOPATH}/src/${GO_PACKAGE_NAME}

# Build go program.

WORKDIR ${GOPATH}/src/${GO_PACKAGE_NAME}
RUN make build

# --- Test go program ---------------------------------------------------------

# Run unit tests.

# RUN go get github.com/jstemmer/go-junit-report \
#  && mkdir -p /output/go-junit-report \
#  && go test -v ${GO_PACKAGE_NAME}/... | go-junit-report > /output/go-junit-report/test-report.xml

# Copy binaries to /output.

RUN mkdir -p /output \
      && cp -R ${GOPATH}/src/${GO_PACKAGE_NAME}/target/*  /output/

# -----------------------------------------------------------------------------
# Stage: final
# -----------------------------------------------------------------------------

FROM ${IMAGE_FINAL} as final
ENV REFRESHED_AT=2023-06-29
LABEL Name="senzing/senzing-tools" \
      Maintainer="support@senzing.com" \
      Version="0.2.6"

# Copy local files from the Git repository.

COPY ./rootfs /

# Copy files from prior step.

COPY --from=go_builder "/output/linux/senzing-tools" "/app/senzing-tools"

# Runtime environment variables.

ENV LD_LIBRARY_PATH=/opt/senzing/g2/lib/

# Runtime execution.

WORKDIR /app
ENTRYPOINT ["/app/senzing-tools"]
