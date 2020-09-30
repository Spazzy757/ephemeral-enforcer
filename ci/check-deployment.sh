#!/bin/bash

deployment_name=$1

deployment_running=false
while [[ "$deployment_running" != "true" ]]
do
    status=$(kubectl get deployment $deployment_name -o jsonpath="{.status.availableReplicas}")
    if [[ "$status" = "1" ]]; then
        deployment_running=true
    fi
done
echo "${deployment_name} is available"
