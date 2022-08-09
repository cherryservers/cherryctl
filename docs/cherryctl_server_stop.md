## cherryctl server stop

Stop a server.

### Synopsis

Stops or powers off a server that is currently powered on.

```
cherryctl server stop -i <server_id> [flags]
```

### Examples

```
  # Stops the specified server:
  cherryctl server stop -i 12345
```

### Options

```
  -h, --help            help for stop
  -i, --server-id int   The ID of a server.
```

### Options inherited from parent commands

```
      --config string    Path to JSON or YAML configuration file
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl server](cherryctl_server.md)	 - Server operations. For more information on provisioning on Cherry Servers, visit https://docs.cherryservers.com/knowledge/product-docs.

