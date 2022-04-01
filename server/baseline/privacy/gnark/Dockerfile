FROM golang:1.17-buster as app-builder

WORKDIR /usr/tmp

COPY ./service ./service

WORKDIR ./service
RUN go install && go build

FROM debian:buster as contract-builder

ARG DEBIAN_FRONTEND=noninteractive

SHELL ["/bin/bash", "-l", "-euxo", "pipefail", "-c"]

RUN apt-get update; \
    apt-get full-upgrade -y; \
    apt-get install -y --no-install-recommends \
        ca-certificates \
        curl \
        apt-transport-https \
        build-essential \
        git \
        libssl-dev \
        wget \
    ; \
    apt-get install -y python3; \
    apt-get clean; \
    rm -rf /var/lib/apt/lists/*

ENV NVM_DIR /usr/local/nvm

RUN mkdir -p "$NVM_DIR"; \
    curl -o- \
        "https://raw.githubusercontent.com/nvm-sh/nvm/master/install.sh" | \
        bash \
    ; \
    source $NVM_DIR/nvm.sh; \
    nvm install --lts --latest-npm

RUN command -v nvm; \
    command -v node

RUN npm install -g truffle

WORKDIR usr/tmp

RUN truffle init

COPY ./service/contracts ./contracts

RUN truffle compile

FROM ubuntu:20.04

WORKDIR usr/app/dist

COPY --from=app-builder /usr/tmp/service/service .
COPY --from=app-builder /usr/tmp/service/keys ./keys

COPY --from=contract-builder /usr/tmp/build/contracts ./contracts

CMD ["./service"]


