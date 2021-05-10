---
description: List Labels from Resources
---

# LabelList

## Usage

```text
ionosctl label list [flags]
```

## Description

Use this command to list all Labels from all Resources under your account. If you want to list all Labels from a specific Resource, use `--resource-type` option together with the Resource Id: `--datacenter-id`, `--server-id`, `--volume-id`.

## Options

```text
  -u, --api-url string         Override default API endpoint (default "https://api.ionos.com/cloudapi/v5")
      --cols strings           Columns to be printed in the standard output (default [Key,Value])
  -c, --config string          Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
      --datacenter-id string   The unique Data Center Id
      --force                  Force command to execute without user input
  -h, --help                   help for list
      --ipblock-id string      The unique IpBlock Id
  -o, --output string          Desired output format [text|json] (default "text")
  -q, --quiet                  Quiet output
      --resource-type string   Resource Type
      --server-id string       The unique Server Id
      --snapshot-id string     The unique Snapshot Id
      --volume-id string       The unique Volume Id
```

## Examples

```text
ionosctl label list 
Key    Value     ResourceType   ResourceId
test   testing   datacenter     aa8e07a2-287a-4b45-b5e9-94761750a53c

ionosctl label list --resource-type datacenter --datacenter-id aa8e07a2-287a-4b45-b5e9-94761750a53c 
Key    Value
test   testing
```
