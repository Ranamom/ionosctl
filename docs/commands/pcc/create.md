---
description: Create a Private Cross-Connect
---

# Create

## Usage

```text
ionosctl pcc create [flags]
```

## Description

Use this command to create a Private Cross-Connect. You can specify the name and the description for the Private Cross-Connect.

## Options

```text
  -u, --api-url string           Override default API endpoint (default "https://api.ionos.com/cloudapi/v5")
      --cols strings             Columns to be printed in the standard output (default [UserId,Firstname,Lastname,Email,S3CanonicalUserId,Administrator,ForceSecAuth,SecAuthActive,Active])
  -c, --config string            Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -h, --help                     help for create
      --ignore-stdin             Force command to execute without user input
  -o, --output string            Desired output format [text|json] (default "text")
      --pcc-description string   The description for the Private Cross-Connect
      --pcc-name string          The name for the Private Cross-Connect
  -q, --quiet                    Quiet output
      --timeout int              Timeout option for Private Cross-Connect to be created [seconds] (default 60)
      --wait                     Wait for Private Cross-Connect to be created
```

## Examples

```text
ionosctl pcc create --pcc-name test --pcc-description "test test" --wait 
PccId                                  Name   Description
e2337b40-52d9-48d2-bcbc-41c5abc29d11   test   test test
RequestId: 64720266-c6e8-4e78-8e31-6754f006dcb1
Status: Command pcc create and request have been successfully executed
```
