# =============================================================================
# Multi-stage Dockerfile — built by Harness CI "Build and Push to ECR" (Kaniko).
# =============================================================================
# Kaniko (used by the BuildAndPushECR step) builds this WITHOUT a Docker daemon
# and WITHOUT privileged mode — a security win over the docker-in-docker pattern
# common in GitHub Actions runners.
#
# Stage 1 compiles a static binary; stage 2 is distroless (no shell, no package
# manager, runs as nonroot) — a tiny, hardened final image. Good Santander story.
# =============================================================================

# ---- Stage 1: build ----
FROM golang:1.22-alpine AS build
WORKDIR /src
# Copy go.mod first for layer caching (no deps here, but keeps the pattern).
COPY go.mod ./
RUN go mod download
COPY . .
# Static build so it runs on distroless. ARG lets CI inject the version/tag.
ARG APP_VERSION=dev
RUN CGO_ENABLED=0 GOOS=linux go build \
      -ldflags "-X main.version=${APP_VERSION}" \
      -o /out/app .

# ---- Stage 2: minimal runtime ----
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/app /app
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app"]
