FROM alpine:latest

COPY . /app

RUN apk add --update --no-cache \
        alpine-sdk \
        imagemagick \
        jpeg-dev \
        libffi-dev \
        libpng-dev \
        libusb-dev \
        libusb \
        openssl-dev \
        protobuf \
        protobuf-dev \
        python3 \
        python3-dev \
        py3-pip \
        tiff \
        tiff-dev
RUN pip3 install --upgrade \
        pip \
        setuptools \
        wheel
RUN pip3 install -r /app/requirements.txt

WORKDIR /app/server
ENTRYPOINT python3 main.py
