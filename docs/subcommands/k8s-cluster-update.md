---
description: Update a Kubernetes Cluster
---

# K8sClusterUpdate

## Usage

```text
ionosctl k8s cluster update [flags]
```

## Description

Use this command to update the name, Kubernetes version, maintenance day and maintenance time of an existing Kubernetes Cluster.

You can wait for the Cluster to be in "ACTIVE" state using `--wait-for-state` flag together with `--timeout` option.

Required values to run command:

* K8s Cluster Id

## Options

```text
  -u, --api-url string            Override default API endpoint (default "https://api.ionos.com/cloudapi/v5")
      --cluster-id string         The unique K8s Cluster Id (required)
      --cluster-name string       The name for the K8s Cluster
      --cols strings              Columns to be printed in the standard output (default [ClusterId,Name,K8sVersion,State,MaintenanceWindow])
  -c, --config string             Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
      --force                     Force command to execute without user input
  -h, --help                      help for update
      --k8s-version string        The K8s version for the Cluster
      --maintenance-day string    The day of the week for Maintenance Window has the English day format as following: Monday or Saturday
      --maintenance-time string   The time for Maintenance Window has the HH:mm:ss format as following: 08:00:00
  -o, --output string             Desired output format [text|json] (default "text")
  -q, --quiet                     Quiet output
      --timeout int               Timeout option for waiting for Cluster to be in ACTIVE state after updating [seconds] (default 600)
      --wait-for-state            Wait for specified Cluster to be in ACTIVE state after updating
```

## Examples

```text
ionosctl k8s cluster update --cluster-id cb47b98f-b8dd-4108-8ac0-b636e36a161d --cluster-name testCluster
ClusterId                              Name          K8sVersion   State
cb47b98f-b8dd-4108-8ac0-b636e36a161d   testCluster   1.19.8       UPDATING
```
