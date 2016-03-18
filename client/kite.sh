#!/bin/sh

export KONTROL_KONTROL_URL="http://${MICROSERVICES_SERVICE_HOST}:6000/kite"
export KITE_USERNAME="openshift"
export KITE_ENVIRONMENT="openshift"

mkdir -p $KITE_HOME

env

exec /go/bin/microservice $@

