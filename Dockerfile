FROM alpine:latest

COPY . /app

RUN apk update && \
    apk add \
        alpine-sdk \
        python3 \
        libffi-dev \
        openssl-dev \
        python3-dev \
        protobuf \
        protobuf-dev && \
    pip3 install --upgrade \
        pip \
        setuptools && \
    pip3 install -r /app/requirements.txt

WORKDIR /app/server
ENTRYPOINT python3 main.py
