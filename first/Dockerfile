FROM centos
MAINTAINER Nick Sabine <nsabine@redhat.com>

RUN yum -y install golang-bin git openssl && yum clean all
ENV GOPATH /go
ENV KITE_HOME /srv/kite

ENV GO_SRC /go/src/first
RUN mkdir -p ${GO_SRC}
ADD first.go ${GO_SRC}
WORKDIR ${GO_SRC}

RUN go get github.com/koding/kite
RUN go install

RUN mkdir -p ${KITE_HOME}
ADD kite.sh ${KITE_HOME}/kite.sh
RUN chmod +x ${KITE_HOME}/kite.sh
RUN chmod -R 777 ${KITE_HOME}

EXPOSE 6001

CMD ${KITE_HOME}/kite.sh

