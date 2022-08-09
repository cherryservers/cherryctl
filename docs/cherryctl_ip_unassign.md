## cherryctl ip unassign

Unassign an IP address.

### Synopsis

Unassign an IP address.

```
cherryctl ip unassign -i <ip_address_id> [flags]
```

### Examples

```
  # Unassign an IP address:
		cherryctl ip unassign -i 12345
```

### Options

```
  -h, --help                   help for unassign
  -i, --ip-address-id string   The ID of an IP address.
```

### Options inherited from parent commands

```
      --config string    Path to JSON or YAML configuration file
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl ip](cherryctl_ip.md)	 - IP address operations.

