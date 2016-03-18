FROM centos
MAINTAINER Nick Sabine <nsabine@redhat.com>

RUN yum -y install golang-bin git openssl && yum clean all
ENV GOPATH=/go
ENV KITE_HOME=/srv/kite

RUN go get github.com/koding/kite/kontrol/kontrol

ADD kontrol.sh /srv/kite/kontrol.sh

RUN chmod +x /srv/kite/kontrol.sh
RUN chmod -R 777 /srv/kite

EXPOSE 6000 2379 4001

CMD ["/srv/kite/kontrol.sh"]

