FROM golang:1.8

ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update -qq \
&& apt-get install -y vim \
&& apt-get install -y locales \
&& apt-get install -y fonts-ipafont

#“ú–{ŒêÝ’è
RUN echo "ja_JP.UTF-8 UTF-8" > /etc/locale.gen \
&& locale-gen ja_JP.UTF-8 \
&& dpkg-reconfigure locales \
&& /usr/sbin/update-locale LANG=ja_JP.UTF-8
ENV TZ='Asia/Tokyo'
ENV LC_ALL ja_JP.UTF-8

WORKDIR /go/src/app
COPY ./app .

ENV GOPATH $GOPATH:$HOME/go
RUN go get github.com/lib/pq
RUN go get -u github.com/golang/dep/cmd/dep

COPY dbtest.sh /dbtest.sh
RUN chmod 744 /dbtest.sh
#CMD ["/dbtest.sh"]
