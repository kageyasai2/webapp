FROM golang:1.8

RUN apt-get update -qq

WORKDIR /go/src/app
COPY ./app .

ENV GOPATH $GOPATH:$HOME/go
RUN go get github.com/lib/pq

COPY dbtest.sh /dbtest.sh
RUN chmod 744 /dbtest.sh
CMD ["/dbtest.sh"]
#main.goを実行
#CMD ["go", "run", "main.go"]