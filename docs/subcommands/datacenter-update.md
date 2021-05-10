---
description: Update a Data Center
---

# DatacenterUpdate

## Usage

```text
ionosctl datacenter update [flags]
```

## Description

Use this command to change a Virtual Data Center's name, description.

You can wait for the Request to be executed using `--wait-for-request` option.

Required values to run command:

* Data Center Id

## Options

```text
  -u, --api-url string                  Override default API endpoint (default "https://api.ionos.com/cloudapi/v5")
      --cols strings                    Columns to be printed in the standard output (default [DatacenterId,Name,Location,State])
  -c, --config string                   Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
      --datacenter-description string   Description of the Data Center
      --datacenter-id string            The unique Data Center Id (required)
      --datacenter-name string          Name of the Data Center
      --force                           Force command to execute without user input
  -h, --help                            help for update
  -o, --output string                   Desired output format [text|json] (default "text")
  -q, --quiet                           Quiet output
      --timeout int                     Timeout option for Request for Data Center update [seconds] (default 60)
      --wait-for-request                Wait for the Request for Data Center update to be executed
```

## Examples

```text
ionosctl datacenter update --datacenter-id 8e543958-04f5-4872-bbf3-b28d46393ac7 --datacenter-description demoDescription --cols "DatacenterId,Description"
DatacenterId                           Description
8e543958-04f5-4872-bbf3-b28d46393ac7   demoDescription
RequestId: 46af6915-9003-4f11-a1fe-bab1eac9bccc
Status: Command datacenter update has been successfully executed
```
