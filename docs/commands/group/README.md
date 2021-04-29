---
description: Group Operations
---

# Group

## Usage

```text
ionosctl group [command]
```

## Description

The sub-command of `ionosctl group` allows you to list, get, create, update, delete Groups, but also operations: add/remove/list/update User/Share from the Group.

## Options

```text
  -u, --api-url string   Override default API endpoint (default "https://api.ionos.com/cloudapi/v5")
      --cols strings     Columns to be printed in the standard output (default [GroupId,Name,CreateDataCenter,CreateSnapshot,ReserveIp,AccessActivityLog,CreatePcc,S3Privilege,CreateBackupUnit,CreateInternetAccess,CreateK8s])
  -c, --config string    Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -h, --help             help for group
      --ignore-stdin     Force command to execute without user input
  -o, --output string    Desired output format [text|json] (default "text")
  -q, --quiet            Quiet output
```

## Related commands

| Command | Description |
| :--- | :--- |
| [ionosctl group add-share](add-share.md) | Create/Add a Resource Share for a Group |
| [ionosctl group add-user](add-user.md) | Add User to a Group |
| [ionosctl group create](create.md) | Create a Group |
| [ionosctl group delete](delete.md) | Delete a Group |
| [ionosctl group get](get.md) | Get a Group |
| [ionosctl group get-share](get-share.md) | Get a Resource Share from a Group |
| [ionosctl group list](list.md) | List Groups |
| [ionosctl group list-resources](list-resources.md) | List Resources from a Group |
| [ionosctl group list-shares](list-shares.md) | List Resources Shares through a Group |
| [ionosctl group list-users](list-users.md) | List Users from a Group |
| [ionosctl group remove-share](remove-share.md) | Remove a Resource Share from a Group |
| [ionosctl group remove-user](remove-user.md) | Remove User from a Group |
| [ionosctl group update](update.md) | Update a Group |
| [ionosctl group update-share](update-share.md) | Update a Resource Share from a Group |
