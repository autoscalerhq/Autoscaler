# Name the first stage "builder" and use Rust's official Docker image.
FROM rust:1.79-alpine3.20 as builder

WORKDIR /usr/src/app

RUN apk update && apk add build-base musl-dev gcc

# This step will cache our dependencies to make sure that `cargo build` only runs when our dependencies change
COPY .  ./

# Get the correct archictecture, build it, then move it to the folder /app/target/release as app.
# This way the runtime container does not need to know what architecture it is.
# Trade-off is that this must run everytime with the build however its rebuilding every time anyway
# Potinetial CD speedup is to leverage buildx build to cache the dependencies and targets then have another run step to build
RUN PACKAGE_NAME=$(grep name Cargo.toml | sed 's/name = "\(.*\)"/\1/' | tr -d '[[:space:]]') && \
    TARGET=$(rustc --print cfg | grep target_arch | grep '"x86_64"' > /dev/null && echo 'x86_64-unknown-linux-musl' || echo 'aarch64-unknown-linux-musl') && \
    rustup target add $TARGET && \
    cargo build --release --target $TARGET && \
    if [ ! -d "/app/target/release" ]; then mkdir -p "/app/target/release"; fi && \
    cp /usr/src/app/target/$TARGET/release/$PACKAGE_NAME /app/target/release/app


# Runtime Stage
FROM alpine:3.20.1

# Install necessary runtime dependencies
RUN apk add --no-cache libgcc libstdc++

# Copy the app binary from the build stage
COPY --from=builder /app/target/release/app /usr/local/bin

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["app"]