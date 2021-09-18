FROM golang:1.16

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1

EXPOSE 8000

RUN apt-get update && \
    apt-get install build-essential protobuf-compiler librdkafka-dev -y

CMD ["tail", "-f", "/dev/null"]

