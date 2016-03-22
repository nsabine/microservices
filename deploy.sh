#oc new-app registry.access.redhat.com/rhel7/etcd
#oc volume dc/etcd --add --overwrite --name=etcd-storage --mount-path /var/lib/etcd -t emptyDir

# math example
oc new-app https://github.com/nsabine/microservices.git --context-dir=square --name=square
oc new-app https://github.com/nsabine/microservices.git --context-dir=client --name=client

# game example
oc new-app https://github.com/nsabine/microservices.git --context-dir=messaging --name=messaging
oc new-app https://github.com/nsabine/microservices.git --context-dir=controller --name=controller
oc new-app https://github.com/nsabine/microservices.git --context-dir=robber --name=robber
oc new-app https://github.com/nsabine/microservices.git --context-dir=cop --name=cop

