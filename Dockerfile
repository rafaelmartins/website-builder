FROM golang:1.14-alpine AS golang_builder

ADD . /code

RUN set -x \
    && ( \
        cd /code \
        && CGO_ENABLED=0 go build -o website-builder . \
    )


FROM debian:buster AS c_builder

RUN set -x \
    && apt-get update \
    && apt-get install -y --no-install-recommends \
        autoconf \
        automake \
        libtool-bin \
        ronn \
        make \
        ca-certificates \
        git \
    && git clone https://github.com/blogc/blogc.git /code \
    && ( \
        cd /code \
        && ./autogen.sh \
        && ./configure \
            --enable-make \
        && make \
    )


FROM debian:buster
LABEL maintainer "Rafael Martins <rafael@rafaelmartins.eng.br>"

RUN set -x \
    && apt-get update \
    && apt-get install -y --no-install-recommends \
        bash \
        make \
        ca-certificates \
        libc-dev \
        libssl-dev \
        ruby-dev \
        rubygems \
        gcc \
        g++ \
    && gem install bundler -v "~>1.0" \
    && gem install bundler \
    && rm -rf \
        /root/.gem \
        /var/cache/apt \
        /var/lib/apt/lists/*

COPY --from=golang_builder /code/website-builder /usr/local/bin/website-builder
COPY --from=c_builder /code/blogc /usr/local/bin/blogc
COPY --from=c_builder /code/blogc-make /usr/local/bin/blogc-make

ENTRYPOINT ["/usr/local/bin/website-builder"]
