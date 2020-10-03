[![Go Report Card](https://goreportcard.com/badge/github.com/Spazzy757/ephemeral-enforcer)](https://goreportcard.com/report/github.com/Spazzy757/ephemeral-enforcer)
[![PkgGoDev](https://pkg.go.dev/badge/Spazzy757/ephemeral-enforcer)](https://pkg.go.dev/Spazzy757/ephemeral-enforcer)
[![codecov](https://codecov.io/gh/spazzy757/ephemeral-enforcer/branch/master/graph/badge.svg)](https://codecov.io/gh/spazzy757/ephemeral-enforcer)
# Kill All The Things

## How it works

Ephemeral Enforcer has one job, to destroy. Install it in a cluster and it will clean up resources after a predefined amount of time. This enforces your workloads to be ephemeral.

**Why would I do that ?**

* Demo environments 
* End to End testing environment
* Developer tooling and demos
* Workloads not defined in code

## Getting Started

To deploy to a cluster:

```bash
export NAMESPACE=${YOUR NAMESPACE}

kubectl -n $NAMESPACE kubectl apply -k manifests/
```
## Environment
```yaml
 # Comma seperated list of resources t skip deleteing
 - name: DISSALLOW_LIST
   value: "statefulsets,secrets"
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
