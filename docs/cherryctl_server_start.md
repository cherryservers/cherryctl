## cherryctl server start

Starts a server.

### Synopsis

Starts or powers on a server that is currently stopped or powered off.

```
cherryctl server start ID [flags]
```

### Examples

```
  # Starts the specified server:
  cherryctl server start 12345
```

### Options

```
  -h, --help   help for start
```

### Options inherited from parent commands

```
      --config string    Path to JSON or YAML configuration file
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl server](cherryctl_server.md)	 - Server operations. For more information on server types check Product Docs: https://docs.cherryservers.com/knowledge/product-docs#compute

