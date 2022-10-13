# Build the Go Binary.
FROM golang:1.19 as build_down_monitor
ENV CGO_ENABLED 0
ARG BUILD_REF
# Copy the source code into the container.
COPY . /service

# Build the app binary.
WORKDIR /service/cmd/app
RUN go build -ldflags "-X main.build=${BUILD_REF}"


# Run the Go Binary in Alpine.
FROM alpine:3.16
ARG BUILD_DATE
ARG BUILD_REF

RUN addgroup -g 1000 -S down_monitor && \
    adduser -u 1000 -h /service -G down_monitor -S down_monitor
COPY --from=build_down_monitor --chown=down_monitor:down_monitor /service/cmd/app /service
WORKDIR /service
USER down_monitor
CMD ["./app"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="down_monitor" \
      org.opencontainers.image.authors="Saeed Jalali" \
      org.opencontainers.image.source="" \
      org.opencontainers.image.revision="${BUILD_REF}" \
      org.opencontainers.image.vendor="Saeed Jalali"   