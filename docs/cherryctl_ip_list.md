## cherryctl ip list

Retrieves list of IP addresses.

### Synopsis

Retrieves the details of IP addresses in the project.

```
cherryctl ip list -p <project_id> [flags]
```

### Examples

```
  # Gets a list of IP addresses in the specified project:
  cherryctl ip list -p 12345
```

### Options

```
  -h, --help             help for list
  -p, --project-id int   The project's ID.
      --type strings     Comma separated list of available IP addresses types (subnet,primary-ip,floating-ip,private-ip)
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

