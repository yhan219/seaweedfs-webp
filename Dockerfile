FROM golang:1.12.4

LABEL version="1.0"
LABEL maintainer="yhan219@sina.com"

RUN apt-get -y update \
 && apt-get install -y git curl wget libjpeg-dev libpng-dev libtool autoconf automake make gcc g++

WORKDIR /usr/local/webp
RUN wget http://downloads.webmproject.org/releases/webp/libwebp-1.0.2.tar.gz \
      && tar -xvzf libwebp-1.0.2.tar.gz \
      && mv libwebp-1.0.2 libwebp && \
      rm libwebp-1.0.2.tar.gz && \
      cd libwebp && \
      ./configure --enable-everything && \
      make && \
      make install && \
      cd .. && \
      rm -rf libwebp

ENV PATH $PATH:/usr/local/webp/bin

RUN ldconfig

RUN mkdir -p /app/src
COPY src/* /app/src
WORKDIR /app

COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

EXPOSE 80

ENTRYPOINT ["./entrypoint.sh"]
