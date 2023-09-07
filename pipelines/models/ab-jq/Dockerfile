FROM alpine:3.18.3@sha256:7144f7bab3d4c2648d7e59409f15ec52a18006a128c733fcff20d3a4a54ba44a

# `jq` is required to parse and format JSON
# `apache2-utils` provides ab benchmarking tool
# `bash` is required to allow stnadardized scripting in bash
# `coreutils` is required for using functions like `date -d`
# hadolint ignore=DL3018
RUN apk update &&       \
    apk add --no-cache jq apache2-utils bash coreutils &&  \
    rm -rf /var/cache/apk/*

LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.title="ab-jq"

CMD [ "bash" ]
