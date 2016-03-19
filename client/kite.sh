#!/bin/sh

export KITE_KONTROL_URL="http://${KONTROL_SERVICE_HOST}:6000/kite"
export KITE_USERNAME="openshift"
export KITE_ENVIRONMENT="openshift"

mkdir -p $KITE_HOME

env

#echo "Registering"
#exec /go/bin/kitectl register -to=$KITE_KONTROL_URL -username=$KITE_USERNAME

echo "Starting Microservice"
exec /go/bin/microservice $@

