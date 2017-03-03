#!/bin/bash

set -e
set -x

# start dummy server
kubectl create -f 10-dummy-deployment.yml

# create service
kubectl create -f 20-service.yml

# show the address
bash 30-show-service.sh
