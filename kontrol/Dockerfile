FROM centos
MAINTAINER Nick Sabine <nsabine@redhat.com>

RUN yum -y install golang-bin git openssl && yum clean all
ENV GOPATH=/go

RUN go get github.com/koding/kite/kontrol/kontrol

ADD kontrol.sh /srv/kite/kontrol.sh
RUN chmod +x /srv/kite/kontrol.sh

EXPOSE 6000

ENTRYPOINT ["/srv/kite/kontrol.sh"]

