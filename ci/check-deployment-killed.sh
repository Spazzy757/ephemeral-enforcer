#!/bin/bash

deployment_name=$1

deployment_killed=false
while [[ "$deployment_killed" != "true" ]]
do
    status=$(kubectl get deployment $deployment_name)
    check="NotFound"
    if grep -q "$status" <<< "$check"; then
        deployment_killed=true
    fi
    sleep 1
done
echo "${deployment_name} has been killed"
