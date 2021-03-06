#!/bin/sh

export KONTROL_CERTS=$KITE_HOME/certs
export KONTROL_PORT=6000
export KITE_USERNAME="openshift"
export KITE_ENVIRONMENT="openshift"
export KONTROL_STORAGE="etcd"
export KONTROL_MACHINES="${ETCD_SERVICE_HOST}:4001"
export KITE_KONTROL_URL="http://${KONTROL_SERVICE_HOST}:6000/kite"
export KONTROL_PUBLICKEYFILE="$KONTROL_CERTS/key_pub.pem"
export KONTROL_PRIVATEKEYFILE="$KONTROL_CERTS/key.pem"

mkdir -p $KITE_HOME
mkdir -p $KONTROL_CERTS


[ -f $KONTROL_PRIVATEKEYFILE ] || openssl genrsa -out $KONTROL_PRIVATEKEYFILE 2048
[ -f $KONTROL_PUBLICKEYFILE ] || openssl rsa -in $KONTROL_PRIVATEKEYFILE -pubout > $KONTROL_PUBLICKEYFILE

env

echo "Starting Controller Microservice"
exec /go/bin/microservice $@

