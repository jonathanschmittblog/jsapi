FROM golang:1.16
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /opt/
RUN git clone https://github.com/jonathanschmittblog/jsapi.git
WORKDIR /opt/jsapi
RUN go mod init
RUN go get
RUN go build jsapi

CMD ["air"]