---
description: List LANs
---

# List

## Usage

```text
ionosctl lan list [flags]
```

## Description

Use this command to get a list of LANs on your account.

Required values to run command:
- Data Center Id

## Options

```text
  -u, --api-url string         Override default API endpoint (default "https://api.ionos.com/cloudapi/v5")
      --cols strings           Columns to be printed in the standard output. Example: --cols "ResourceId,Name" (default [LanId,Name,Public])
  -c, --config string          Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
      --datacenter-id string   The unique Data Center Id
  -h, --help                   help for list
      --ignore-stdin           Force command to execute without user input
  -o, --output string          Desired output format [text|json] (default "text")
  -q, --quiet                  Quiet output
  -v, --verbose                Enable verbose output
```

## Examples

```text
ionosctl lan list --datacenter-id f28c0edd-d5ef-48f2-b8a3-aa8f6b55da3d 
LanId   Name                                                Public
4       demoLan                                             false
3       demoLAN                                             true
2       Switch of LB f16dfcc1-9181-400b-a08d-7fe15ca0e9af   false
1       Switch of LB 3f9f14a9-5fa8-4786-ba86-a91f9daded2c   false
```
