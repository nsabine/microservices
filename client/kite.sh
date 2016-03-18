#!/bin/sh

export KONTROL_KONTROLURL="http://${MICROSERVICES_SERVICE_HOST}:6000/kite"
mkdir -p $KITE_HOME
mkdir -p $KONTROL_CERTS

env

exec /go/bin/client $@

