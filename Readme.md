# Kill All The Things

## How it works

Ephemeral Enforcer has one job, to destroy. Install it in a cluster and it will clean up resources after a predefined amount of time. This enforces your workloads to be ephemeral.

## Getting Started:

To deploy to a cluster:

```bash
export NAMESPACE=${YOUR NAMESPACE}

kubectl -n $NAMESPACE kubectl apply -k manifests/
```
## Environment
```yaml
 # Can Be Run IN or OUt Of the Cluster
 - name: IN_CLUSTER
   value: "true"
 # Which Namespace To Kill Workloads
 - name: NAMESPACE
   value: example
 # Name of the Enforcer (So it doesnt delete itself)
 - name: EPHEMERAL_ENFORCER_NAME
   value: "ephemeral-enforcer"
 # Comma Seperated List of Prefixes to skip
 - name: SKIPPED_PREFIXES
   value: "default,kube"
 # When Should the Enforcer Check
 - name: ENFORCER_SCHEDULE
   value: "*/5 * * * *"
 # How Long Workloads Should Be Allowed to Live For (in minutes)
 - name: WORKLOAD_TTL
   value: "60"
```
