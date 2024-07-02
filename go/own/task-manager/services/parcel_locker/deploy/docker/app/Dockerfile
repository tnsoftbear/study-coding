# syntax=docker/dockerfile:1
FROM rust:1-alpine3.19 as builder
# This is important, see https://github.com/rust-lang/docker-rust/issues/85
ENV RUSTFLAGS="-C target-feature=-crt-static"
# if needed, add additional dependencies here
# RUN apk add --no-cache musl-dev
RUN apk add musl-dev
# set the workdir and copy the source into it
WORKDIR /app
COPY ./ /app
# do a release build
RUN cargo build --release
RUN strip target/release/parcel_locker

# use a plain alpine image, the alpine version needs to match the builder
FROM alpine:3.19
WORKDIR /app
# if needed, install additional dependencies here
# RUN apk add --no-cache libgcc curl
RUN apk add libgcc curl
# copy the binary into the final image
COPY --from=builder /app/target/release/parcel_locker /app/server
# set the binary as entrypoint
ENTRYPOINT ["/app/server"]
