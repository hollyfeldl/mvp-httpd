FROM golang:latest

ADD . /go/src/github.com/hollyfeldl/mvp-httpd
WORKDIR /go/src/github.com/hollyfeldl/mvp-httpd

RUN go get github.com/kylelemons/go-gypsy/yaml
RUN go install github.com/hollyfeldl/mvp-httpd

COPY mvp-httpd-docker.yaml /go/bin/mvp-httpd.yaml

ADD ./html /var/www/html/
COPY ./ssl/fullchain.pem /etc/ssl/certs
COPY ./ssl/privkey.pem /etc/ssl/private

WORKDIR /go/bin
ENTRYPOINT /go/bin/mvp-httpd

EXPOSE 8080 8443