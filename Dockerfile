# builder image
FROM golang:1.14-buster as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -o ephemeral-enforcer ./cmd/ephemeral-enforcer

# generate clean, final image for end users
FROM gcr.io/distroless/base-debian10
COPY --from=builder /build/ephemeral-enforcer /
CMD ["/ephemeral-enforcer"]