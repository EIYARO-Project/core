# Build Eiyaro in a stock Go builder container
FROM golang:1.19.8-alpine as builder

RUN apk add --no-cache make git

ADD . /go/src/github.com/eiyaro/ey
WORKDIR /go/src/github.com/eiyaro/ey
RUN make eiyarod && make eiyarocli

# Pull Eiyaro into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/eiyaro/ey/cmd/eiyarod/eiyarod /usr/local/bin/eiyarod
COPY --from=builder /go/src/github.com/eiyaro/ey/cmd/eiyarocli/eiyarocli /usr/local/bin/eiyarocli

# Create a runnable that can check the data init
RUN mkdir /start
RUN echo '#!bin/sh' > /start/runme.sh && \
    echo 'eiyarod init --chain_id "$CHAIN_ID"' >> /start/runme.sh && \
    echo 'eiyarod node "$@"' >> /start/runme.sh

RUN chmod +x /start/runme.sh

VOLUME /root/.eiyaro
ENV CHAIN_ID=mainnet
EXPOSE 1999 46656 46657 9888

ENTRYPOINT ["/start/runme.sh"]

CMD [""]
