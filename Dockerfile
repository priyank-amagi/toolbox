FROM ubuntu:22.04

RUN apt-get update -y -qq && DEBIAN_FRONTEND=noninteractive apt-get -q -y install --no-install-recommends apt-utils \
    locales \
    lsb-release \
    openssl \
    ca-certificates && update-ca-certificates \
    wget

WORKDIR /app

COPY ./bin/toolbox /app/toolbox

ENTRYPOINT ["/app/toolbox"]

# CMD ["todos list"]