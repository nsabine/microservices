#!/bin/sh

export KONTROL_HOME=/srv/kite
export KONTROL_CERTS=$KONTROL_HOME/certs
export KONTROL_PORT=6000
export KONTROL_USERNAME="kontrol"
export KONTROL_STORAGE="etcd"
export KONTROL_MACHINES="${ETCD_SERVICE}"
export KONTROL_KONTROLURL="http://127.0.0.1:6000/kite"
export KONTROL_PUBLICKEYFILE="$KONTROL_CERTS/key_pub.pem"
export KONTROL_PRIVATEKEYFILE="$KONTROL_CERTS/key.pem"

mkdir -p $KONTROL_HOME
mkdir -p $KONTROL_CERTS


[ -f $KONTROL_PRIVATEKEYFILE ] || openssl genrsa -out $KONTROL_PRIVATEKEYFILE 2048
[ -f $KONTROL_PUBLICKEYFILE ] || openssl rsa -in $KONTROL_PRIVATEKEYFILE -pubout > $KONTROL_PUBLICKEYFILE

/go/bin/kontrol -initial || exit 1
exec /go/bin/kontrol $@

