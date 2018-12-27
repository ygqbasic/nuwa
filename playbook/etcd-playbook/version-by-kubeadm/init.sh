#! /bin/bash

set -e

path=`dirname $0`

image=k8s.gcr.io/etcd-amd64:${1}
echo "" >> ${path}/group_vars/etcd.yml
echo "image: ${image}" >> ${path}/group_vars/etcd.yml

docker pull ${image}
docker save ${image} > ${path}/file/etcd.tar
bzip2 -z --best ${path}/file/etcd.tar
