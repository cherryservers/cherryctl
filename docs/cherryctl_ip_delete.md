## cherryctl ip delete

Delete an IP address.

### Synopsis

Deletes the specified IP address with a confirmation prompt. To skip the confirmation use --force.

```
cherryctl ip delete ID [flags]
```

### Examples

```
  # Deletes the specified IP:
  cherryctl ip delete 30c15082-a06e-4c43-bfc3-252616b46eba
  >
  ✔ Are you sure you want to delete IP address 30c15082-a06e-4c43-bfc3-252616b46eba: y
  		
  # Deletes a server, skipping confirmation:
  cherryctl ip delete 30c15082-a06e-4c43-bfc3-252616b46eba -f
```

### Options

```
  -f, --force   Skips confirmation for the server deletion.
  -h, --help    help for delete
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

* [cherryctl ip](cherryctl_ip.md)	 - IP address operations. For more information on IP addresses check Product Docs: https://docs.cherryservers.com/knowledge/product-docs#ip-addressing

