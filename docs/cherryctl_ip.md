## cherryctl ip

IP address operations. For more information on IP addresses, check out the Product Docs: https://docs.cherryservers.com/knowledge/product-docs#ip-addressing

### Synopsis

IP address operations: get, list, create, update, assign, unassign and delete.

### Options

```
  -h, --help   help for ip
```

### Options inherited from parent commands

```
      --api-url string   Override default API endpoint (default "https://api.cherryservers.com/v1/")
      --config string    Path to configuration file directory. The CHERRY_CONFIG environment variable can be used as well.
      --context string   Specify a custom context name (default "default")
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl](cherryctl.md)	 - Cherry Servers Command Line Interface (CLI)
* [cherryctl ip assign](cherryctl_ip_assign.md)	 - Assign an IP address to a specified server or other IP address.
* [cherryctl ip create](cherryctl_ip_create.md)	 - Create floating IP address.
* [cherryctl ip delete](cherryctl_ip_delete.md)	 - Delete an IP address.
* [cherryctl ip get](cherryctl_ip_get.md)	 - Get an IP address details.
* [cherryctl ip list](cherryctl_ip_list.md)	 - Retrieves list of IP addresses.
* [cherryctl ip unassign](cherryctl_ip_unassign.md)	 - Unassign an IP address.
* [cherryctl ip update](cherryctl_ip_update.md)	 - Update IP address.

