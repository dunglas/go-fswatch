# syntax=docker/dockerfile:1
FROM golang:1.23

RUN apt-get update && apt-get install -y \
    build-essential \
    jq \
    && rm -rf /var/lib/apt/lists/*

RUN mkdir fswatch/ && \
    curl -sL "$(curl -sL \
    -H 'Accept: application/vnd.github+json' \
    -H 'X-GitHub-Api-Version: 2022-11-28' \
    https://api.github.com/repos/emcrisostomo/fswatch/releases/latest | jq -r '.assets[0].browser_download_url')" | tar -xz -C fswatch --strip-components=1 && \
    cd fswatch/ && \
    ./configure && \
    make -j"$(getconf _NPROCESSORS_ONLN)" && \
    make install && \
    cd .. && \
    rm -Rf fswatch/

ENV LD_LIBRARY_PATH=/usr/local/lib
