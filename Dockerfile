# syntax=docker/dockerfile:1

############################
# STEP 1 build executable binary
############################
# golang alpine 1.13.5
FROM golang@sha256:0991060a1447cf648bab7f6bb60335d1243930e38420bee8fec3db1267b84cfa AS builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata dumb-init && update-ca-certificates

# Create appuser
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

# Fetch dependencies.
RUN go get -d -v

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' -a \
    -o /go/bin/todo-app-api .

############################
# STEP 2 build a small image
############################
FROM alpine@sha256:e7d88de73db3d3fd9b2d63aa7f447a10fd0220b7cbf39803c803f2af9ba256b3

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /usr/bin/dumb-init /usr/bin/dumb-init

# Copy our static executable
COPY --from=builder /go/bin/todo-app-api /go/bin/todo-app-api

# Use an unprivileged user.
USER appuser:appuser

# Run the todo-app-api binary.
# See https://github.com/gofiber/fiber/issues/1036#issuecomment-841763449
ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["sh", "-c", "/go/bin/todo-app-api"]