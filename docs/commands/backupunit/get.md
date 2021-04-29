---
description: Get a BackupUnit
---

# Get

## Usage

```text
ionosctl backupunit get [flags]
```

## Description

Use this command to retrieve details about a specific BackupUnit.

Required values to run command:

* BackupUnit Id

## Options

```text
  -u, --api-url string         Override default API endpoint (default "https://api.ionos.com/cloudapi/v5")
      --backupunit-id string   The unique BackupUnit Id [Required flag]
      --cols strings           Columns to be printed in the standard output (default [BackupUnitId,Name,Email])
  -c, --config string          Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -h, --help                   help for get
      --ignore-stdin           Force command to execute without user input
  -o, --output string          Desired output format [text|json] (default "text")
  -q, --quiet                  Quiet output
```

## Examples

```text
ionosctl backupunit get --backupunit-id 9fa48167-6375-4d93-b33c-e1ba3f461c17 
BackupUnitId                           Name          Email
9fa48167-6375-4d93-b33c-e1ba3f461c17   test1234567   testrandom20@ionos.com
```
