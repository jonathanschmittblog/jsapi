FROM golang:latest
WORKDIR /go/src
# RUN rm -r jsapi
RUN git clone https://github.com/jonathanschmittblog/jsapi.git
WORKDIR /go/src/jsapi
RUN go mod init
RUN go get -u
RUN go install
EXPOSE 3000