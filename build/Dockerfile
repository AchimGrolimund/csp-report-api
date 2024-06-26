# Obtain certs for final stage
FROM alpine:3.19 as authority
RUN mkdir /user && \
    echo 'appuser:x:1000:1000:appuser:/:' > /user/passwd && \
    echo 'appgroup:x:1000:' > /user/group
RUN apk --no-cache add ca-certificates

# Build app binary for final stage
FROM --platform=$BUILDPLATFORM golang:1.22-alpine3.19 AS builder
WORKDIR /app
ARG TARGETOS TARGETARCH
COPY .. .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -a -o /csp-report-api ./cmd/csp-report-api/main.go

# Final stage
FROM scratch
COPY --from=authority /user/group /user/passwd /etc/
COPY --from=authority /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /csp-report-api ./
USER appuser:appgroup
EXPOSE 8080
ENTRYPOINT ["./csp-report-api"]