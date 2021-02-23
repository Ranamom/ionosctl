---
description: Ionos Cloud CLI
---

# Ionosctl

## Usage

```text
ionosctl [command]
```

## Description

IonosCTL is a command-line interface (CLI) for the Ionos Cloud API.

## Options

```text
  -u, --api-url string   Override default API endpoint (default "https://api.ionos.com/cloudapi/v5")
  -c, --config string    Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -h, --help             help for ionosctl
      --ignore-stdin     Force command to execute without user input
  -o, --output string    Desired output format [text|json] (default "text")
  -q, --quiet            Quiet output
  -v, --verbose          Enable verbose output
```

## Related commands

| Command | Description |
| :------ | :---------- |
| [ionosctl completion](completion/) | Generate code to enable auto-completion with `TAB` key |
| [ionosctl datacenter](datacenter/) | Data Center Operations |
| [ionosctl lan](lan/) | LAN Operations |
| [ionosctl loadbalancer](loadbalancer/) | Load Balancer Operations |
| [ionosctl location](location/) | Location Operations |
| [ionosctl login](login.md) | Authentication command for SDK |
| [ionosctl nic](nic/) | Network Interfaces Operations |
| [ionosctl request](request/) | Request Operations |
| [ionosctl server](server/) | Server Operations |
| [ionosctl version](version.md) | Show the current version |
| [ionosctl volume](volume/) | Volume Operations |
