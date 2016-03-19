#!/bin/sh

export KITE_KONTROL_URL="http://${KONTROL_SERVICE_HOST}:6000/kite"
export KONTROL_USERNAME="openshift"
export KONTROL_ENVIRONMENT="openshift"

mkdir -p $KITE_HOME

env

exec /go/bin/microservice $@

