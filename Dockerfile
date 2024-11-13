FROM golang:alpine AS builder

WORKDIR /opt
COPY . .

# Get go dependencies
RUN go get -v -d

# Build
RUN go build -a -gcflags=all="-l -B" -ldflags="-w -s" -o /opt/uao

# Compress executable with UPX
FROM gruebel/upx:latest as upx
COPY --from=builder /opt/uao /opt/uao
RUN upx --best --lzma -o /opt/uao-compressed /opt/uao

# Make a scratch image
FROM scratch

# Copy UAO from the builder to the scratch container
COPY --from=upx /opt/uao-compressed /opt/uao

# import curl from current repository image
COPY --from=ghcr.io/tarampampam/curl:8.0.1 /bin/curl /bin/curl

HEALTHCHECK --interval=5s --timeout=2s --retries=3 CMD [ \
    "curl", "--fail", "http://127.0.0.1:80/health" \
]

# Run UAO
ENTRYPOINT ["/opt/uao"]
