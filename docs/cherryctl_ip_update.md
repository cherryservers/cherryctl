## cherryctl ip update

Update IP address.

### Synopsis

Update tags, ptr record, a record or target server of a IP address.

```
cherryctl ip update ID [--ptr-record <ptr>] [--a-record <a>] [--tags <tags>] [flags]
```

### Examples

```
  # Updates a record and tags:
  cherryctl ip update 30c15082-a06e-4c43-bfc3-252616b46eba --a-record stage --tags="env=stage"
```

### Options

```
      --a-record string     Slug of the region from where IP address will requested.
  -h, --help                help for update
      --ptr-record string   Slug of the region from where IP address will requested.
      --tags strings        Tag or list of tags for the server: --tags="key=value,env=prod".
```

### Options inherited from parent commands

```
      --config string    Path to JSON or YAML configuration file
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl ip](cherryctl_ip.md)	 - IP address operations. For more information on IP addresses check Product Docs: https://docs.cherryservers.com/knowledge/product-docs#ip-addressing

