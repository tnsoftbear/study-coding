FROM rust:1-alpine3.20 AS chef
WORKDIR /app
RUN apk add musl-dev
RUN cargo install cargo-chef

FROM chef AS planner
COPY . .
RUN cargo chef prepare --recipe-path recipe.json

FROM chef AS builder
COPY --from=planner /app/recipe.json recipe.json
# Build dependencies - this is the caching Docker layer!
RUN cargo chef cook --release --recipe-path recipe.json
# Build application
COPY . .
RUN cargo build --release --bin parcel_locker

# We do not need the Rust toolchain to run the binary!
FROM alpine:3.20 AS runtime
WORKDIR /app
RUN apk add --no-cache libgcc curl
COPY --from=builder /app/target/release/parcel_locker /app/parcel_locker_service
ENTRYPOINT ["/app/parcel_locker_service"]
