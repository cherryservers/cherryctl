## cherryctl ip unassign

Unassign an IP address.

### Synopsis

Unassign an IP address.

```
cherryctl ip unassign ID [flags]
```

### Examples

```
  # Unassign an IP address:
		cherryctl ip unassign 30c15082-a06e-4c43-bfc3-252616b46eba
```

### Options

```
  -h, --help   help for unassign
```

### Options inherited from parent commands

```
      --api-url string   Override default API endpoint (default "https://api.cherryservers.com/v1/")
      --config string    Path to JSON or YAML configuration file
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl ip](cherryctl_ip.md)	 - IP address operations.

