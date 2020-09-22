# builder image
FROM golang:1.13-alpine3.11 as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -o ephemeral-enforcer ./cmd/ephemeral-enforcer

# generate clean, final image for end users
FROM golang:1.13-alpine3.11
RUN mkdir /app
COPY --from=builder /build/ephemeral-enforcer /app
WORKDIR /app

# executable
ENTRYPOINT [ "/app/ephemeral-enforcer" ]
