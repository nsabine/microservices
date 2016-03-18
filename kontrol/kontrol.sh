#!/bin/sh

export KONTROL_CERTS=$KITE_HOME/certs
export KONTROL_PORT=6000
export KONTROL_USERNAME="openshift"
export KONTROL_ENVIRONMENT="openshift"
export KONTROL_STORAGE="etcd"
export KONTROL_MACHINES="${ETCD_SERVICE_HOST}:4001"
export KONTROL_URL="http://${MICROSERVICES_SERVICE_HOST}:6000/kite"
export KONTROL_PUBLICKEYFILE="$KONTROL_CERTS/key_pub.pem"
export KONTROL_PRIVATEKEYFILE="$KONTROL_CERTS/key.pem"

mkdir -p $KITE_HOME
mkdir -p $KONTROL_CERTS


[ -f $KONTROL_PRIVATEKEYFILE ] || openssl genrsa -out $KONTROL_PRIVATEKEYFILE 2048
[ -f $KONTROL_PUBLICKEYFILE ] || openssl rsa -in $KONTROL_PRIVATEKEYFILE -pubout > $KONTROL_PUBLICKEYFILE

env

/go/bin/kontrol -initial || exit 1
exec /go/bin/kontrol $@

