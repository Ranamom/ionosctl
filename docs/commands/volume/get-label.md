---
description: Get a Label from a Volume
---

# GetLabel

## Usage

```text
ionosctl volume get-label [flags]
```

## Description

Use this command to get information about a specified Label from a Volume.

Required values to run command:

* Data Center Id
* Volume Id
* Label Key

## Options

```text
  -u, --api-url string         Override default API endpoint (default "https://api.ionos.com/cloudapi/v5")
      --cols strings           Columns to be printed in the standard output (default [DatacenterId,Name,Location])
  -c, --config string          Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
      --datacenter-id string   The unique Data Center Id [Required flag]
  -h, --help                   help for get-label
      --ignore-stdin           Force command to execute without user input
      --label-key string       The unique Label Key [Required flag]
  -o, --output string          Desired output format [text|json] (default "text")
  -q, --quiet                  Quiet output
      --volume-id string       The unique Volume Id [Required flag]
```

## Examples

```text
ionosctl volume get-label --datacenter-id ed612a0a-9506-4b56-8d1b-ce2b04090f19 --volume-id 5d23eee2-45e5-44fe-96fe-e15aba2c48f5 --label-key test
Key    Value
test   testvolume
```
