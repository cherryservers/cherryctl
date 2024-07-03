## cherryctl server

Server operations. For more information on server types check Product Docs: https://docs.cherryservers.com/knowledge/product-docs#compute

### Synopsis

Server operations: create, get, list, delete, start, stop, reboot, reinstall and reset-bmc-password.

### Options

```
  -h, --help   help for server
```

### Options inherited from parent commands

```
      --api-url string   Override default API endpoint (default "https://api.cherryservers.com/v1/")
      --config string    Path to JSON or YAML configuration file
      --context string   Specify a custom context name (default "default")
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl](cherryctl.md)	 - Cherry Servers Command Line Interface (CLI)
* [cherryctl server create](cherryctl_server_create.md)	 - Create a server.
* [cherryctl server delete](cherryctl_server_delete.md)	 - Delete a server.
* [cherryctl server get](cherryctl_server_get.md)	 - Retrieves server details.
* [cherryctl server list](cherryctl_server_list.md)	 - Retrieves server list.
* [cherryctl server reboot](cherryctl_server_reboot.md)	 - Reboot a server.
* [cherryctl server reinstall](cherryctl_server_reinstall.md)	 - Reinstall a server.
* [cherryctl server rescue](cherryctl_server_rescue.md)	 - Rescue a server.
* [cherryctl server reset-bmc-password](cherryctl_server_reset-bmc-password.md)	 - Reset server BMC password.
* [cherryctl server start](cherryctl_server_start.md)	 - Starts a server.
* [cherryctl server stop](cherryctl_server_stop.md)	 - Stop a server.
* [cherryctl server update](cherryctl_server_update.md)	 - Update server.

