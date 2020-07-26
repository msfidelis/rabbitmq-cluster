FROM golang:1.13 AS builder

RUN pwd; ls -lha

# Install Air
RUN curl -fLo air https://git.io/linux_air
RUN mv air /usr/bin/air && chmod +x /usr/bin/air

WORKDIR $GOPATH/src/consumer

COPY . ./

RUN pwd; ls -lha

RUN go get -u

RUN mkdir -p tmp/main ; chmod 777 . -R

CMD ["/usr/bin/air"]