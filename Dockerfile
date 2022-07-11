## NOTE: This image uses goreleaser to build image
# if building manually please run go build ./cmd/harvey first and then build

# Choose alpine as a base image to make this useful for CI, as many
# CI tools expect an interactive shell inside the container
FROM alpine:latest as production

#COPY --from=builder /build/harvey /usr/bin/harvey
COPY harvey /usr/bin/harvey
RUN chmod +x /usr/bin/harvey

WORKDIR /workdir

ENTRYPOINT ["/usr/bin/harvey"]
CMD ["--help"]
